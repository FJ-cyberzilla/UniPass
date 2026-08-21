package geo

import (
	"strings"
)

// EvaluateTrust checks if an IP location payload indicates proxy/VPN distortion
func EvaluateTrust(locationPayload string) bool {
	lower := strings.ToLower(locationPayload)
	// Simple heuristic flags for known datacenter/VPN markers
	if strings.Contains(lower, "datacenter") || strings.Contains(lower, "vpn") || strings.Contains(lower, "proxy") {
		return false
	}
	return true
}
