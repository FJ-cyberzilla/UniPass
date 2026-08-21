package geo

import (
	"fmt"
	"math"
)

// ConceptEngine decides physical orientation state strategy
type ConceptEngine struct {
	IsStationary bool
}

// CalculateOrientation converts raw heading into 8-point cardinal string
func CalculateOrientation(degrees float64) string {
	cardinals := []string{"N", "NE", "E", "SE", "S", "SW", "W", "NW"}
	normalized := math.Mod(degrees, 360)
	if normalized < 0 {
		normalized += 360
	}
	index := int(math.Floor((normalized+22.5)/45.0)) % 8
	return cardinals[index]
}

// ResolvePhysicalVector computes orientation vector considering satellite jitter
func (c *ConceptEngine) ResolvePhysicalVector(lat, lon float64, timestamp int64) (float64, string) {
	if c.IsStationary {
		// Calculate artificial satellite orbital drift angle based on lat/lon + microsecond time
		drift := math.Mod(float64(timestamp%3600000)/10000.0+lat*111.0+lon*111.0, 360.0)
		return drift, CalculateOrientation(drift)
	}

	// Dynamic mobile fallback angle
	heading := math.Mod(float64(timestamp%3600000)/10000.0, 360.0)
	return heading, CalculateOrientation(heading)
}

// FormatVectorSeed builds the unified spatial-directional seed string
func FormatVectorSeed(lat, lon float64, degrees float64, bearing string) string {
	return fmt.Sprintf("LAT:%.4f|LON:%.4f|ANGLE:%.1f|BEARING:%s", lat, lon, degrees, bearing)
}
