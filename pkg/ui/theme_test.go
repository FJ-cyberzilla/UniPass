package ui

import (
	"testing"
)

func TestRenderCentered(t *testing.T) {
	tc := &ThemeController{Width: 80, IsInteractive: true}

	input := "TEST_HEADER"
	visibleLen := len(input)

	centered := tc.RenderCentered(input, visibleLen)
	expectedPadding := 34

	if len(centered)-visibleLen != expectedPadding {
		t.Errorf("Expected %d leading spaces, got string length %d", expectedPadding, len(centered))
	}
}

func TestBuildGradientBar(t *testing.T) {
	bar := BuildGradientBar(0.5)
	if bar == "" {
		t.Errorf("Expected non-empty progress bar string")
	}
}
