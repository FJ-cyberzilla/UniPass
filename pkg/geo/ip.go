package geo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	IPGeoAPIURL    = "http://ip-api.com/json/?fields=status,message,country,city,lat,lon"
	IPGeoUserAgent = "UniPass-CLI/1.0"
	IPGeoTimeout   = 3 * time.Second
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

// IPProvider implements the GeolocationProvider interface
type IPProvider struct{}

// Resolve queries public IP geolocation APIs
func (p *IPProvider) Resolve() (*GeolocationResult, error) {
	client := &http.Client{
		Timeout: IPGeoTimeout,
	}

	req, err := http.NewRequest("GET", IPGeoAPIURL, nil)
	if err != nil {
		return &GeolocationResult{Error: err}, err
	}

	req.Header.Set("User-Agent", IPGeoUserAgent)

	resp, err := client.Do(req)
	if err != nil {
		return &GeolocationResult{Error: err}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return &GeolocationResult{Error: fmt.Errorf("status: %s", resp.Status)}, fmt.Errorf("status: %s", resp.Status)
	}

	var payload IPGeoResponse
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return &GeolocationResult{Error: err}, err
	}

	if payload.Status != "success" {
		return &GeolocationResult{Error: fmt.Errorf("failed: %s", payload.Message)}, fmt.Errorf("failed: %s", payload.Message)
	}

	return &GeolocationResult{
		Latitude:  payload.Lat,
		Longitude: payload.Lon,
		Source:    "ip",
		Error:     nil,
	}, nil
}
