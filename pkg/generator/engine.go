package generator

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"unicode"

	"unipass/pkg/geo"
)

// ValidateName ensures input name meets requirements
func ValidateName(name string) error {
	if len(name) < 4 || len(name) > 12 {
		return fmt.Errorf("name length must be between 4 and 12 characters")
	}
	for _, char := range name {
		if !unicode.IsLetter(char) {
			return fmt.Errorf("name must contain only letters")
		}
	}
	return nil
}

// GenerateFull produces a complete 64-char SHA-256 password string
func GenerateFull(name string, loc geo.LocationDetails, timestamp int64) string {
	concept := &geo.VectorEngine{IsStationary: true}
	degrees, bearing := concept.ResolvePhysicalVector(loc.Lat, loc.Lon, timestamp)
	vectorSeed := geo.FormatVectorSeed(loc.Lat, loc.Lon, degrees, bearing)

	rawInput := fmt.Sprintf("%s|%s|%d", name, vectorSeed, timestamp)
	hash := sha256.Sum256([]byte(rawInput))
	return hex.EncodeToString(hash[:])
}

// GenerateShort produces a high-entropy 12-char password (first 12 characters of full hash)
func GenerateShort(name string, loc geo.LocationDetails, timestamp int64) string {
	fullHash := GenerateFull(name, loc, timestamp)
	return fullHash[:12]
}
