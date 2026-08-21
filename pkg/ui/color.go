package ui

// --- ANSI 256-COLOR PALETTE DICTIONARY ---
const (
	ColorReset     = "\033[0m"
	ColorBold      = "\033[1m"
	ColorDim       = "\033[2m"
	ColorPink      = "\033[38;5;205m"
	ColorPurple    = "\033[38;5;141m"
	ColorOceanBlue = "\033[38;5;33m"
	ColorSysBlue   = "\033[38;5;39m"
	ColorCyan      = "\033[38;5;51m"
	ColorGray      = "\033[38;5;242m"
	ColorYellow    = "\033[38;5;220m"
	ColorGreen     = "\033[38;5;46m"
	ColorRed       = "\033[38;5;196m"
	ColorETAVal    = "\033[38;5;208m"
	ColorSuccess   = "\033[38;5;46m"
)

// Contextual Aliases
const ColorStatusText = ColorPurple

// Purple-to-Blue Gradient Array (Refined)
var BarGradient = []string{
	"\033[38;5;129m", // Deep Purple
	"\033[38;5;99m",  // Violet
	"\033[38;5;93m",  // Dark Violet
	"\033[38;5;63m",  // Light Blue-Violet
	"\033[38;5;33m",  // Deep Blue
	"\033[38;5;27m",  // Vivid Blue
	"\033[38;5;26m",  // Rich Blue
	"\033[38;5;25m",  // Darker Blue
}
