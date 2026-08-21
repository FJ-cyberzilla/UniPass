package ui

import (
	"testing"
)

func TestColorConstants(t *testing.T) {
	tests := []struct {
		name     string
		got      string
		expected string
	}{
		{"Reset", ColorReset, "\033[0m"},
		{"Red", ColorRed, "\033[38;5;196m"},
		{"Green", ColorGreen, "\033[38;5;46m"},
		{"Pink", ColorPink, "\033[38;5;205m"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.expected {
				t.Errorf("%s = %q, want %q", tt.name, tt.got, tt.expected)
			}
		})
	}
}

func TestBarGradient(t *testing.T) {
	if len(BarGradient) != 8 {
		t.Errorf("BarGradient length = %d, want 8", len(BarGradient))
	}
	
	// Check if all are non-empty
	for i, color := range BarGradient {
		if color == "" {
			t.Errorf("BarGradient[%d] is empty", i)
		}
	}
}
