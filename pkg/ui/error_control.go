package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// PromptWithValidation handles ESC key exits and double empty Enter prompts
func PromptWithValidation(reader *bufio.Reader, promptText string, emptyCount *int) (string, bool) {
	fmt.Printf("%s%s%s", ColorCyan, promptText, ColorReset)

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("\n%sExiting UniPass...%s\n", ColorGray, ColorReset)
		os.Exit(0)
	}

	// Detect ESC key press (ASCII 27 / \x1b)
	if strings.Contains(input, "\x1b") {
		fmt.Printf("\n%s[ESC DETECTED] Exiting UniPass system...%s\n", ColorGray, ColorReset)
		os.Exit(0)
	}

	cleanInput := strings.TrimSpace(input)

	// Handle empty input attempts
	if cleanInput == "" {
		*emptyCount++
		if *emptyCount == 1 {
			// First empty press: Show warning in Yellow
			fmt.Printf("%s[WARNING] Cannot be left empty! Press Enter again or ESC to exit.%s\n", ColorYellow, ColorReset)
			return "", false
		} else if *emptyCount >= 2 {
			// Second empty press: Exit cleanly without error text
			os.Exit(0)
		}
	}

	// Reset empty counter on valid input
	*emptyCount = 0
	return cleanInput, true
}
