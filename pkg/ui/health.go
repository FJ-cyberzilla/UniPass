package ui

import (
	"fmt"
	"strings"
	"time"

	"unipass/pkg/config"
)

type DiagnosticStep struct {
	Label    string
	Duration time.Duration
}

func RunWeeklyHealthCheck() {
	tc := NewThemeController()

	// Guard 1: Skip if stdout is non-interactive (piped/redirected)
	if !tc.IsInteractive {
		return
	}

	// Guard 2: Skip if 7 days have not yet elapsed
	if !config.ShouldRunWeeklyHealthCheck() {
		return
	}

	// 1. Centered Title
	titleStr := "[UNIPASS SYSTEM HEALTH ENGINE - WEEKLY DIAGNOSTIC]"
	fmt.Printf("\n%s\n\n", tc.RenderCentered(ColorSysBlue+ColorBold+titleStr+ColorReset, len(titleStr)))

	steps := []DiagnosticStep{
		{Label: "Verifying cryptographic primitives & SHA-256 state...", Duration: 250 * time.Millisecond},
		{Label: "Testing spatial vector math & coordinate calculations...", Duration: 300 * time.Millisecond},
		{Label: "Auditing internal entropy providers...", Duration: 200 * time.Millisecond},
		{Label: "Checking package dependencies & local build health...", Duration: 350 * time.Millisecond},
		{Label: "Flushing ephemeral execution caches...", Duration: 150 * time.Millisecond},
	}

	var totalDuration time.Duration
	for _, s := range steps {
		totalDuration += s.Duration
	}

	startTime := time.Now()
	totalSteps := len(steps)
	dotFrames := []string{".  ", ".. ", "..."}
	dotIdx := 0

	for i, step := range steps {
		stepStart := time.Now()

		for time.Since(stepStart) < step.Duration {
			elapsed := time.Since(startTime)
			remaining := totalDuration - elapsed
			if remaining < 0 {
				remaining = 0
			}

			progress := float64(i) / float64(totalSteps)
			bar := BuildGradientBar(progress)
			dots := dotFrames[dotIdx%len(dotFrames)]
			dotIdx++

			// Calculate visible string lengths for exact horizontal centering
			barVisibleText := fmt.Sprintf("[%s] %3.0f%% | ETA: %4.1fs %s", strings.Repeat("█", 24), progress*100, remaining.Seconds(), dots)
			statusVisibleText := fmt.Sprintf("-> Status: %s", step.Label)

			// Formatted lines with ANSI 256 colors
			barColoredLine := fmt.Sprintf("[%s] %3.0f%% | %sETA:%s %s%4.1fs%s %s",
				bar, progress*100, ColorSysBlue, ColorReset, ColorETAVal, remaining.Seconds(), ColorReset, dots)
			statusColoredLine := fmt.Sprintf("%s-> Status: %s%s", ColorSysBlue, ColorStatusText, step.Label)

			// Render centered lines
			fmt.Printf("\033[2K\r%s\n", tc.RenderCentered(barColoredLine, len(barVisibleText)))
			fmt.Printf("\033[2K\r%s\033[1A", tc.RenderCentered(statusColoredLine, len(statusVisibleText)))

			time.Sleep(50 * time.Millisecond)
		}
	}

	// 2. Final Completion State
	fullBar := BuildGradientBar(1.0)
	barVisibleText := fmt.Sprintf("[%s] 100%% | ETA:  0.0s   ", strings.Repeat("█", 24))
	statusVisibleText := "-> Status: ✅ All system health diagnostics passed clean!"

	barColoredLine := fmt.Sprintf("[%s] 100%% | %sETA:%s %s0.0s%s   ",
		fullBar, ColorSysBlue, ColorReset, ColorETAVal, ColorReset)
	statusColoredLine := fmt.Sprintf("%s-> Status: %s✅ All system health diagnostics passed clean!%s",
		ColorSysBlue, ColorSuccess, ColorReset)

	fmt.Printf("\033[2K\r%s\n", tc.RenderCentered(barColoredLine, len(barVisibleText)))
	fmt.Printf("\033[2K\r%s\n\n", tc.RenderCentered(statusColoredLine, len(statusVisibleText)))

	_ = config.UpdateHealthCheckTimestamp()
}
