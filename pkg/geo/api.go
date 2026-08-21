package geo

import (
	"errors"
	"fmt"
	"net"
)

// GeolocationResult holds the resolved vector data and its source tier
type GeolocationResult struct {
	Latitude  string
	Longitude string
	Source    string
	Error     error
}

// ResolveVectorCascade executes the multi-tier fallback strategy out-of-the-box
func ResolveVectorCascade() (string, error) {
	// Tier 1: Try Hardware GPS
	if gpsBytes, err := FetchDeviceGPS(); err == nil && len(gpsBytes) > 0 {
		return "gps:" + string(gpsBytes), nil
	}

	// Tier 2: Try IP Geolocation API with threat/VPN validation
	if ipGeo, err := FetchIPCoordinates(); err == nil && ipGeo != nil {
		ipGeoStr := fmt.Sprintf("%f,%f", ipGeo.Lat, ipGeo.Lon)
		if EvaluateTrust(ipGeoStr) {
			return "ip:" + ipGeoStr, nil
		}
	}

	// Tier 3: Try Local Network / ISP Interface Fingerprinting
	if netInfo, err := fetchLocalNetworkFingerprint(); err == nil && netInfo != "" {
		return "net:" + netInfo, nil
	}

	// Tier 4: Ultimate offline deterministic machine fallback
	machineID := FetchSystemIdentity()
	if machineID != "" {
		return "fallback-machine:" + machineID, nil
	}

	return "", errors.New("critical: all geolocation cascade tiers failed")
}

func fetchLocalNetworkFingerprint() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil || len(interfaces) == 0 {
		return "", errors.New("no network interfaces available")
	}
	return interfaces[0].Name + "-" + interfaces[0].HardwareAddr.String(), nil
}
