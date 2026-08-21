package geo

import (
	"testing"
)

func TestEvaluateTrust(t *testing.T) {
	validator := &TrustValidatorImpl{}
	tests := []struct {
		name    string
		payload string
		want    bool
	}{
		{"Valid location", "45.0,90.0", true},
		{"VPN marker", "45.0,90.0-VPN", false},
		{"Proxy marker", "45.0,90.0-PROXY", false},
		{"Datacenter marker", "45.0,90.0-DATACENTER", false},
		{"Mixed case", "45.0,90.0-vPn", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validator.Evaluate(tt.payload)
			if got != tt.want {
				t.Errorf("Evaluate() = %v, want %v", got, tt.want)
			}
		})
	}
}
