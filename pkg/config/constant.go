package config

import "time"

// Branding Metadata
const (
	AppName    = "UniPass"
	BrandName  = "FJ-Cyberzilla"
	AppVersion = "1.0.0"
)

// Default AES-256 Key Fallback (Exactly 32 Bytes)
// #nosec G101 -- Fallback key for development mode when environment variable is not present.
const DefaultCipherKey = "a-32-byte-secret-key-for-unipass"

// ANSI Color Palette
const (
	ColorReset         = "\033[0m"
	ColorVintagePurple = "\033[38;2;138;43;226m"
	ColorPink          = "\033[38;2;255;105;180m"
	ColorPurple        = "\033[38;2;186;85;211m"
	ColorOceanBlue     = "\033[38;2;0;191;255m"
	ColorRed           = "\033[38;2;255;69;0m"
	ColorGreen         = "\033[38;2;50;205;50m"
)

// Precision Math & Cryptographic Parameters
const (
	EarthVolumetricMeanDiameterMeters = 12742000.0
	ModuloPrimeSeed                   = 1000003
	MinNameLength                     = 4
	MaxNameLength                     = 12
	AESKeyLength                      = 32
	HTTPTimeoutDuration               = 6 * time.Second
	BannerAnimationDelayMs            = 8 * time.Millisecond
)
