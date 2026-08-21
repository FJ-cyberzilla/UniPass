package geo

import (
	"strings"
)

// TrustValidatorImpl implements the TrustValidator interface
type TrustValidatorImpl struct{}

// Evaluate checks if an IP location payload indicates proxy/VPN distortion
func (v *TrustValidatorImpl) Evaluate(locationPayload string) bool {
	lower := strings.ToLower(locationPayload)
	// Simple heuristic flags for known datacenter/VPN markers
	if strings.Contains(lower, "datacenter") || strings.Contains(lower, "vpn") || strings.Contains(lower, "proxy") {
		return false
	}
	return true
}
