package geo

import (
	"strings"
	"testing"
)

func TestCalculateOrientation(t *testing.T) {
	tests := []struct {
		degrees  float64
		expected string
	}{
		{0.0, "N"},
		{45.0, "NE"},
		{90.0, "E"},
		{135.0, "SE"},
		{180.0, "S"},
		{225.0, "SW"},
		{270.0, "W"},
		{315.0, "NW"},
		{360.0, "N"},
		{-45.0, "NW"},
	}

	for _, tt := range tests {
		result := CalculateOrientation(tt.degrees)
		if result != tt.expected {
			t.Errorf("For %.1f deg expected %s, got %s", tt.degrees, tt.expected, result)
		}
	}
}

func TestConceptEngine_ResolvePhysicalVector(t *testing.T) {
	stationaryEngine := &ConceptEngine{IsStationary: true}
	drift, bearing := stationaryEngine.ResolvePhysicalVector(51.5074, -0.1278, 1787302400000000)
	if bearing == "" {
		t.Error("Expected non-empty bearing string for stationary engine")
	}
	if drift < 0 || drift >= 360 {
		t.Errorf("Expected drift angle between 0 and 360, got %.2f", drift)
	}

	mobileEngine := &ConceptEngine{IsStationary: false}
	mAngle, mBearing := mobileEngine.ResolvePhysicalVector(51.5074, -0.1278, 1787302400000000)
	if mBearing == "" {
		t.Error("Expected non-empty bearing string for mobile engine")
	}
	if mAngle < 0 || mAngle >= 360 {
		t.Errorf("Expected mobile angle between 0 and 360, got %.2f", mAngle)
	}
}

func TestFormatVectorSeed(t *testing.T) {
	seed := FormatVectorSeed(51.5074, -0.1278, 225.0, "SW")
	if !strings.Contains(seed, "LAT:51.5074") || !strings.Contains(seed, "BEARING:SW") {
		t.Errorf("Unexpected formatted vector seed output: %s", seed)
	}
}
