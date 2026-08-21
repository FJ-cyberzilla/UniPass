package ui

import (
	"fmt"
	"strings"
)

// FormatBannerHeader formats a clean header string without printing to stdout
func FormatBannerHeader(version string) string {
	cleanVersion := strings.TrimPrefix(strings.TrimPrefix(version, "v"), "V")
	return fmt.Sprintf(">>> UniPass CLI | UNIPASS SYSTEM v%s <<<", cleanVersion)
}
