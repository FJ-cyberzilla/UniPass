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

// Purple-to-Blue Gradient Array
var BarGradient = []string{
	"\033[38;5;129m", "\033[38;5;128m", "\033[38;5;127m", "\033[38;5;126m",
	"\033[38;5;93m", "\033[38;5;63m", "\033[38;5;33m", "\033[38;5;27m",
}
