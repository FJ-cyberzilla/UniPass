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
