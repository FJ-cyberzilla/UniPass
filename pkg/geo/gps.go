package geo

import (
	"context"
	"encoding/json"
	"os/exec"
	"time"
)

// CommandRunner allows overriding the command execution path during tests
var execCommandContext = exec.CommandContext

// GPSResponse represents the expected JSON output from termux-location
type GPSResponse struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// GPSProvider implements the GeolocationProvider interface
type GPSProvider struct{}

// Resolve attempts to fetch GPS location
func (p *GPSProvider) Resolve() (*GeolocationResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	cmd := execCommandContext(ctx, "termux-location")
	output, err := cmd.Output()
	if err != nil {
		return &GeolocationResult{Error: err}, err
	}

	var gps GPSResponse
	if err := json.Unmarshal(output, &gps); err != nil {
		return &GeolocationResult{Error: err}, err
	}

	return &GeolocationResult{
		Latitude:  gps.Latitude,
		Longitude: gps.Longitude,
		Source:    "gps",
		IsAuto:    true,
		Error:     nil,
	}, nil
}
