package ui

import (
	"os"
	"strings"
)

type ThemeController struct {
	Width         int
	IsInteractive bool
}

func NewThemeController() *ThemeController {
	fileInfo, err := os.Stdout.Stat()
	isTTY := err == nil && (fileInfo.Mode()&os.ModeCharDevice) != 0

	// Default fallback width
	width := 80

	return &ThemeController{
		Width:         width,
		IsInteractive: isTTY,
	}
}

func (tc *ThemeController) RenderCentered(text string, visibleLen int) string {
	if visibleLen >= tc.Width {
		return text
	}
	leftPadding := (tc.Width - visibleLen) / 2
	return strings.Repeat(" ", leftPadding) + text
}

func BuildGradientBar(progress float64) string {
	barLength := 24
	filled := int(progress * float64(barLength))
	var sb strings.Builder

	for i := 0; i < barLength; i++ {
		if i < filled {
			colorIdx := (i * len(BarGradient)) / barLength
			sb.WriteString(BarGradient[colorIdx] + "█")
		} else {
			sb.WriteString(ColorDim + "░")
		}
	}
	sb.WriteString(ColorReset)
	return sb.String()
}
