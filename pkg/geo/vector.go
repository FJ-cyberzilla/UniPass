package geo

import (
	"fmt"
	"math"
)

const (
	DegreesInCircle = 360.0
	CardinalSegments = 8
	SegmentDegrees  = 45.0
	SegmentOffset   = 22.5
	DriftTimestampMod = 3600000
	DriftScaleFactor  = 10000.0
	LatLonMultiplier  = 111.0
)

// VectorEngine decides physical orientation state strategy
type VectorEngine struct {
	IsStationary bool
}

// CalculateOrientation converts raw heading into 8-point cardinal string
func CalculateOrientation(degrees float64) string {
	cardinals := []string{"N", "NE", "E", "SE", "S", "SW", "W", "NW"}
	normalized := math.Mod(degrees, DegreesInCircle)
	if normalized < 0 {
		normalized += DegreesInCircle
	}
	index := int(math.Floor((normalized+SegmentOffset)/SegmentDegrees)) % CardinalSegments
	return cardinals[index]
}

// ResolvePhysicalVector computes orientation vector considering satellite jitter
func (c *VectorEngine) ResolvePhysicalVector(lat, lon float64, timestamp int64) (float64, string) {
	if c.IsStationary {
		// Calculate artificial satellite orbital drift angle based on lat/lon + microsecond time
		drift := math.Mod(float64(timestamp%DriftTimestampMod)/DriftScaleFactor+lat*LatLonMultiplier+lon*LatLonMultiplier, DegreesInCircle)
		return drift, CalculateOrientation(drift)
	}

	// Dynamic mobile fallback angle
	heading := math.Mod(float64(timestamp%DriftTimestampMod)/DriftScaleFactor, DegreesInCircle)
	return heading, CalculateOrientation(heading)
}

// FormatVectorSeed builds the unified spatial-directional seed string
func FormatVectorSeed(lat, lon float64, degrees float64, bearing string) string {
	return fmt.Sprintf("LAT:%.4f|LON:%.4f|ANGLE:%.1f|BEARING:%s", lat, lon, degrees, bearing)
}
