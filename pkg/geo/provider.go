package geo

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"unipass/pkg/config"
)

type LocationDetails struct {
	Lat           float64
	Lon           float64
	EarthDiameter float64
}

type nominatimResponse struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

// FetchCoordinates resolves latitude and longitude safely via Nominatim
func FetchCoordinates(city, country string) (*LocationDetails, error) {
	cityClean := strings.TrimSpace(city)
	countryClean := strings.TrimSpace(country)

	if cityClean == "" || countryClean == "" {
		return nil, fmt.Errorf("city and country inputs cannot be empty")
	}

	// Safe URL construction using url.URL to prevent SSRF (CWE-918 / G704)
	baseURL := &url.URL{
		Scheme: "https",
		Host:   "nominatim.openstreetmap.org",
		Path:   "/search",
	}

	queryParams := url.Values{}
	queryParams.Set("city", cityClean)
	queryParams.Set("country", countryClean)
	queryParams.Set("format", "json")
	queryParams.Set("limit", "1")

	baseURL.RawQuery = queryParams.Encode()
	safeEndpoint := baseURL.String()

	req, err := http.NewRequest("GET", safeEndpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed creating HTTP request: %w", err)
	}

	req.Header.Set("User-Agent", "UniPass-App/1.0.0 (security-audit-compliant)")

	client := &http.Client{Timeout: config.HTTPTimeoutDuration}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("geocoding network request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("geocoding service responded with HTTP %d", resp.StatusCode)
	}

	var results []nominatimResponse
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return nil, fmt.Errorf("failed parsing geocoding response: %w", err)
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("no geographic coordinates found for %s, %s", cityClean, countryClean)
	}

	lat, err := strconv.ParseFloat(results[0].Lat, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid latitude value received: %w", err)
	}

	lon, err := strconv.ParseFloat(results[0].Lon, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid longitude value received: %w", err)
	}

	return &LocationDetails{
		Lat:           lat,
		Lon:           lon,
		EarthDiameter: config.EarthVolumetricMeanDiameterMeters,
	}, nil
}

// ===== NEW ASTRONOMICAL METHODS =====

// EquationOfTime returns the difference between apparent and mean solar time in seconds
// Range: -864 to +984 seconds, changes daily
func (loc *LocationDetails) EquationOfTime(timestamp int64) float64 {
	t := time.Unix(timestamp, 0)
	day := float64(t.YearDay())

	// B = orbital angle in degrees (day 81 = spring equinox)
	B := (360.0 / 365.25) * (day - 81.0)
	BRad := B * math.Pi / 180.0

	// EoT in minutes, then convert to seconds
	eotMin := 9.87*math.Sin(2*BRad) - 7.53*math.Cos(BRad) - 1.5*math.Sin(BRad)
	return eotMin * 60.0
}

// SolarDeclination returns the Sun's declination angle in degrees
// Range: -23.44° to +23.44°, changes ~0.4° per day
func (loc *LocationDetails) SolarDeclination(timestamp int64) float64 {
	t := time.Unix(timestamp, 0)
	day := float64(t.YearDay())

	// 23.44° axial tilt, day 80 = spring equinox (March 21)
	return 23.44 * math.Sin((day-80.0)*(360.0/365.25)*math.Pi/180.0)
}

// SunElevation returns the sun's elevation angle in degrees at the current location and time
// Range: -90° to +90°, changes every second
func (loc *LocationDetails) SunElevation(timestamp int64) float64 {
	hour := float64(time.Unix(timestamp, 0).Hour()) + float64(time.Unix(timestamp, 0).Minute())/60.0 + float64(time.Unix(timestamp, 0).Second())/3600.0

	// Solar declination
	decRad := loc.SolarDeclination(timestamp) * math.Pi / 180.0

	// Local solar time calculation
	latRad := loc.Lat * math.Pi / 180.0
	lonOffset := loc.Lon / 15.0 // Each 15° longitude = 1 hour

	// Solar noon at this longitude
	solarNoon := 12.0 - lonOffset

	// Hour angle (degrees from solar noon)
	hourAngleDeg := (hour - solarNoon) * 15.0
	hourAngleRad := hourAngleDeg * math.Pi / 180.0

	// Sin of elevation angle
	sinElevation := math.Sin(latRad)*math.Sin(decRad) +
		math.Cos(latRad)*math.Cos(decRad)*math.Cos(hourAngleRad)

	// Clamp to avoid domain errors (-1 to 1)
	if sinElevation > 1.0 {
		sinElevation = 1.0
	}
	if sinElevation < -1.0 {
		sinElevation = -1.0
	}

	return math.Asin(sinElevation) * 180.0 / math.Pi
}

// SunAzimuth returns the sun's compass bearing in degrees (0° = North, 90° = East)
// Changes every second, location-specific
func (loc *LocationDetails) SunAzimuth(timestamp int64) float64 {
	hour := float64(time.Unix(timestamp, 0).Hour()) + float64(time.Unix(timestamp, 0).Minute())/60.0 + float64(time.Unix(timestamp, 0).Second())/3600.0

	decRad := loc.SolarDeclination(timestamp) * math.Pi / 180.0
	latRad := loc.Lat * math.Pi / 180.0
	lonOffset := loc.Lon / 15.0
	solarNoon := 12.0 - lonOffset
	hourAngleDeg := (hour - solarNoon) * 15.0
	hourAngleRad := hourAngleDeg * math.Pi / 180.0

	// Elevation first
	sinElevation := math.Sin(latRad)*math.Sin(decRad) +
		math.Cos(latRad)*math.Cos(decRad)*math.Cos(hourAngleRad)
	if sinElevation > 1.0 {
		sinElevation = 1.0
	}
	if sinElevation < -1.0 {
		sinElevation = -1.0
	}
	elevationRad := math.Asin(sinElevation)

	// Azimuth formula (0° = North, measured clockwise)
	// From: sin(azimuth) = -cos(declination) * sin(hour_angle) / cos(elevation)
	cosElevation := math.Cos(elevationRad)
	if math.Abs(cosElevation) < 1e-10 {
		return 0.0 // Sun at zenith or nadir
	}

	sinAzimuth := -math.Cos(decRad) * math.Sin(hourAngleRad) / cosElevation
	cosAzimuth := (math.Sin(decRad) - math.Sin(latRad)*math.Sin(elevationRad)) /
		(math.Cos(latRad) * cosElevation)

	// Clamp
	if sinAzimuth > 1.0 {
		sinAzimuth = 1.0
	}
	if sinAzimuth < -1.0 {
		sinAzimuth = -1.0
	}
	if cosAzimuth > 1.0 {
		cosAzimuth = 1.0
	}
	if cosAzimuth < -1.0 {
		cosAzimuth = -1.0
	}

	// Convert to degrees (0° = North)
	azimuthRad := math.Atan2(sinAzimuth, cosAzimuth)
	azimuthDeg := azimuthRad * 180.0 / math.Pi
	return math.Mod(azimuthDeg+360.0, 360.0)
}

// AstronomicalSeed returns a combined astronomical entropy string
// Includes: EoT, Declination, Elevation, Azimuth
func (loc *LocationDetails) AstronomicalSeed(timestamp int64) string {
	return fmt.Sprintf("%.10f|%.10f|%.10f|%.10f",
		loc.EquationOfTime(timestamp),
		loc.SolarDeclination(timestamp),
		loc.SunElevation(timestamp),
		loc.SunAzimuth(timestamp))
}

// EffectiveDiameter returns the local Earth diameter adjusted for latitude
// (Equatorial: 12,756 km, Polar: 12,714 km)
func (loc *LocationDetails) EffectiveDiameter() float64 {
	equatorial := 12756360.0 // meters
	polar := 12713780.0      // meters
	latRad := loc.Lat * math.Pi / 180.0

	// Ellipsoid radius at given latitude
	cosLat := math.Cos(latRad)
	sinLat := math.Sin(latRad)

	// WGS84 ellipsoid formula (semi-major axis a, semi-minor axis b)
	a := equatorial / 2.0 // semi-major axis
	b := polar / 2.0      // semi-minor axis

	// Radius at latitude
	radius := math.Sqrt(
		(math.Pow(a*a*cosLat, 2) + math.Pow(b*b*sinLat, 2)) /
			(math.Pow(a*cosLat, 2) + math.Pow(b*sinLat, 2)),
	)

	return radius * 2.0 // diameter
}
