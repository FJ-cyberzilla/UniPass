package geo

import (
	"os"
	"runtime"
)

// SystemProvider implements the GeolocationProvider interface
type SystemProvider struct{}

// Resolve returns system identification as a fallback location
func (p *SystemProvider) Resolve() (*GeolocationResult, error) {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown-host"
	}
	identity := runtime.GOOS + "-" + runtime.GOARCH + "-" + hostname
	return &GeolocationResult{
		Source:   "system:" + identity,
		IsAuto:   true,
		IsManual: false,
		Error:    nil,
	}, nil
}
