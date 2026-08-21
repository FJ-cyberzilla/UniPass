package geo

import (
	"testing"
)

func TestFetchCoordinatesValidation(t *testing.T) {
	_, err := FetchCoordinates("", "")
	if err == nil {
		t.Error("Expected error for empty provider string")
	}
}

func TestFetchCoordinatesMockServer(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping network test in short mode")
	}

	loc, err := FetchCoordinates("IP", "")
	if err != nil {
		t.Logf("Network request failed gracefully: %v", err)
	} else if loc.Lat == 0 && loc.Lon == 0 {
		t.Error("Expected non-zero coordinates from location provider")
	}
}

func TestAstronomicalMethods(t *testing.T) {
	loc := &LocationDetails{Lat: 40.7128, Lon: -74.0060} // New York
	timestamp := int64(1724250000) // Arbitrary timestamp

	eot := loc.EquationOfTime(timestamp)
	if eot == 0 {
		t.Log("EquationOfTime is 0, this might be expected at certain times, check if 0 is a valid result")
	}

	dec := loc.SolarDeclination(timestamp)
	if dec > 23.44 || dec < -23.44 {
		t.Errorf("SolarDeclination out of range: %f", dec)
	}

	elev := loc.SunElevation(timestamp)
	if elev > 90 || elev < -90 {
		t.Errorf("SunElevation out of range: %f", elev)
	}

	azi := loc.SunAzimuth(timestamp)
	if azi < 0 || azi >= 360 {
		t.Errorf("SunAzimuth out of range: %f", azi)
	}
}
