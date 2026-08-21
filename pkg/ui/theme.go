package ui

import (
	"golang.org/x/term"
	"os"
)

// GetTerminalWidth retrieves the current terminal column width, defaulting to 80 if non-interactive or unavailable.
func GetTerminalWidth() int {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil || width <= 0 {
		return 80
	}
	return width
}
