package geo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// IPGeoResponse represents the JSON response structure from ip-api.com
type IPGeoResponse struct {
	Status  string  `json:"status"`
	Country string  `json:"country"`
	City    string  `json:"city"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
	Message string  `json:"message"`
}

// FetchIPCoordinates queries public IP geolocation APIs as a fallback mechanism
func FetchIPCoordinates() (*LocationDetails, error) {
	client := &http.Client{
		Timeout: 3 * time.Second,
	}

	req, err := http.NewRequest("GET", "http://ip-api.com/json/?fields=status,message,country,city,lat,lon", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create IP geolocation request: %w", err)
	}

	req.Header.Set("User-Agent", "UniPass-CLI/1.0")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("IP geolocation service unreachable: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("IP geolocation service returned status: %s", resp.Status)
	}

	var payload IPGeoResponse
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return nil, fmt.Errorf("failed to decode IP location response: %w", err)
	}

	if payload.Status != "success" {
		return nil, fmt.Errorf("IP geolocation failed: %s", payload.Message)
	}

	return &LocationDetails{
		Lat:           payload.Lat,
		Lon:           payload.Lon,
		EarthDiameter: 12742000.0,
	}, nil
}
