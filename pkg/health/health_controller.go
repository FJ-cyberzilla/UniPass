package health

import (
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type EnvironmentType string

const (
	EnvTermux     EnvironmentType = "Android (Termux)"
	EnvWSL2       EnvironmentType = "Linux (WSL2)"
	EnvWSL1       EnvironmentType = "Linux (WSL1)"
	EnvLinux      EnvironmentType = "Native Linux"
	EnvPowerShell EnvironmentType = "Windows (PowerShell)"
	EnvWindowsCMD EnvironmentType = "Windows (CMD)"
	EnvUnknown    EnvironmentType = "Unknown Environment"
)

type SystemHealth struct {
	OS              string          `json:"os"`
	Arch            string          `json:"arch"`
	Environment     EnvironmentType `json:"environment"`
	HasTTY          bool            `json:"has_tty"`
	HasColorSupport bool            `json:"has_color_support"`
	HasGPS          bool            `json:"has_gps"`
}

type HealthController struct{}

func NewHealthController() *HealthController {
	return &HealthController{}
}

// DetectEnvironment identifies the precise runtime shell/subsystem
func (hc *HealthController) DetectEnvironment() EnvironmentType {
	switch runtime.GOOS {
	case "android":
		return EnvTermux
	case "windows":
		if os.Getenv("PSModulePath") != "" {
			return EnvPowerShell
		}
		return EnvWindowsCMD
	case "linux":
		if isTermuxPath() {
			return EnvTermux
		}
		if isWSL() {
			if isWSL2() {
				return EnvWSL2
			}
			return EnvWSL1
		}
		return EnvLinux
	default:
		return EnvUnknown
	}
}

// Diagnose performs a full system-wide health check
func (hc *HealthController) Diagnose() SystemHealth {
	env := hc.DetectEnvironment()
	_, termuxLocationErr := exec.LookPath("termux-location")

	return SystemHealth{
		OS:              runtime.GOOS,
		Arch:            runtime.GOARCH,
		Environment:     env,
		HasTTY:          isTerminal(os.Stdout.Fd()),
		HasColorSupport: os.Getenv("TERM") != "dumb",
		HasGPS:          termuxLocationErr == nil,
	}
}

func isTermuxPath() bool {
	return strings.Contains(os.Getenv("PREFIX"), "com.termux")
}

func isWSL() bool {
	data, err := os.ReadFile("/proc/version")
	if err != nil {
		return false
	}
	return strings.Contains(strings.ToLower(string(data)), "microsoft")
}

func isWSL2() bool {
	data, err := os.ReadFile("/proc/sys/kernel/osrelease")
	if err != nil {
		return false
	}
	return strings.Contains(strings.ToLower(string(data)), "wsl2")
}

func isTerminal(fd uintptr) bool {
	// Basic terminal check without CGO
	return os.Getenv("TERM") != "" || os.Getenv("PSModulePath") != ""
}
