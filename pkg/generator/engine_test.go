package generator

import (
	"testing"

	"unipass/pkg/geo"
)

func TestValidateName(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"Valid Name", "Mina", false},
		{"Valid Max Length Alphabetic", "Alexander", false},
		{"Too Short", "Bob", true},
		{"Too Long", "Supercalifragilistic", true},
		{"Contains Digits", "Alexander123", true},
		{"Empty", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateName(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateName(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
		})
	}
}

func TestDualPasswordGeneration(t *testing.T) {
	loc := geo.LocationDetails{
		Lat:           51.507446,
		Lon:           -0.127765,
		EarthDiameter: 12742000.0,
	}
	timestamp := int64(1700000000000000)

	shortPass := GenerateShort("Mina", loc, timestamp)
	fullPass := GenerateFull("Mina", loc, timestamp)

	if len(shortPass) != 12 {
		t.Fatalf("Expected short password length 12, got %d", len(shortPass))
	}

	if len(fullPass) != 64 {
		t.Fatalf("Expected full SHA-256 hash length 64, got %d", len(fullPass))
	}

	if fullPass[:12] != shortPass {
		t.Fatal("Expected short password to be prefix of full hash")
	}
}
