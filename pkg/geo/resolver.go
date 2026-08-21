package geo

import (
	"errors"
	"fmt"
)

type Resolver struct {
	EnableGPS      bool
	EnableIP       bool
	EnableMachineID bool
}

func NewResolver(enableGPS, enableIP, enableMachineID bool) *Resolver {
	return &Resolver{
		EnableGPS:       enableGPS,
		EnableIP:        enableIP,
		EnableMachineID: enableMachineID,
	}
}

// ResolveBestVector executes the cascade based on configured priorities
func (gm *Resolver) ResolveBestVector() (string, string, error) {
	// Priority 1: Hardware GPS (if enabled and available)
	if gm.EnableGPS {
		if gpsBytes, err := FetchDeviceGPS(); err == nil && len(gpsBytes) > 0 {
			return "gps", string(gpsBytes), nil
		}
	}

	// Priority 2: IP Geolocation (with threat/VPN validation)
	if gm.EnableIP {
		if ipGeo, err := FetchIPCoordinates(); err == nil && ipGeo != nil {
			ipGeoStr := fmt.Sprintf("%f,%f", ipGeo.Lat, ipGeo.Lon)
			if EvaluateTrust(ipGeoStr) {
				return "ip-trusted", ipGeoStr, nil
			}
		}
	}

	// Priority 3: Isolated Machine/System Entropy Fallback
	if gm.EnableMachineID {
		machineEntropy := FetchSystemIdentity()
		if machineEntropy != "" {
			return "machine-fallback", machineEntropy, nil
		}
	}

	return "", "", errors.New("resolver: all priority cascade tiers exhausted or disabled")
}
