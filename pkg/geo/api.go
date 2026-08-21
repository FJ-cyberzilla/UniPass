package geo

import (
	"errors"
	"net"
)

// NetworkProvider implements the GeolocationProvider interface
type NetworkProvider struct{}

// Resolve attempts to fetch network fingerprint
func (p *NetworkProvider) Resolve() (*GeolocationResult, error) {
	interfaces, err := net.Interfaces()
	if err != nil || len(interfaces) == 0 {
		return &GeolocationResult{Error: errors.New("no network interfaces available")}, err
	}
	fingerprint := interfaces[0].Name + "-" + interfaces[0].HardwareAddr.String()
	return &GeolocationResult{
		Source:   "net:" + fingerprint,
		IsAuto:   false,
		IsManual: false,
		Error:    nil,
	}, nil
}
