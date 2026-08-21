package ui

import (
	"testing"
)

func TestGetTerminalWidth(t *testing.T) {
	width := GetTerminalWidth()
	if width < 40 {
		t.Errorf("Expected terminal width fallback of at least 40, got %d", width)
	}
}

func TestRenderHeaderAndBanner(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Rendering header panicked: %v", r)
		}
	}()

	RenderHeader("UNIPASS SYSTEM", "V1.0.0")
}
