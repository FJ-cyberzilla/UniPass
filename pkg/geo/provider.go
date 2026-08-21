package geo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

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
