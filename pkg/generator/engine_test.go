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

func TestGeneratePassword(t *testing.T) {
	loc := geo.LocationDetails{
		Lat:           51.507446,
		Lon:           -0.127765,
		EarthDiameter: 12742000.0,
	}
	timestamp := int64(1700000000000000)

	tests := []struct {
		name      string
		inputName string
	}{
		{"Valid Name", "Mina"},
		{"Min Length", "ABCD"},
		{"Max Length", "AlexanderMax"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shortPass := GenerateShort(tt.inputName, loc, timestamp)
			fullPass := GenerateFull(tt.inputName, loc, timestamp)

			if len(shortPass) != 12 {
				t.Errorf("Expected short password length 12, got %d", len(shortPass))
			}

			if len(fullPass) != 64 {
				t.Errorf("Expected full SHA-256 hash length 64, got %d", len(fullPass))
			}

			if fullPass[:12] != shortPass {
				t.Errorf("Expected short password to be prefix of full hash")
			}
		})
	}
}
