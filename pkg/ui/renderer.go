package ui

import (
	"fmt"
	"strings"
)

// RenderHeader prints a centered, full-width formatted header banner to stdout
func RenderHeader(title, version string) {
	width := GetTerminalWidth()
	cleanVersion := strings.TrimPrefix(strings.TrimPrefix(version, "v"), "V")
	headerText := fmt.Sprintf(">>> %s | UNIPASS SYSTEM v%s <<<", title, cleanVersion)

	divider := strings.Repeat("=", width)
	fmt.Println(divider)

	tc := &ThemeController{Width: width, IsInteractive: true}
	fmt.Println(tc.RenderCentered(headerText, len(headerText)))
	fmt.Println(divider)
}
