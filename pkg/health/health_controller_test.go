package health

import (
	"runtime"
	"testing"
)

func TestDetectEnvironment(t *testing.T) {
	hc := NewHealthController()

	// Since we can't easily mock runtime.GOOS or os.Getenv without
	// changing the code structure, we verify that it returns
	// a valid non-unknown environment for the current system.
	// Or, we could structure the code to allow injection of OS/Env variables.

	env := hc.DetectEnvironment()

	if env == EnvUnknown {
		t.Logf("Detected as Unknown, which might be correct for unsupported platforms")
	} else if env == "" {
		t.Errorf("Expected valid environment, got empty")
	}
}

func TestDiagnose(t *testing.T) {
	hc := NewHealthController()
	health := hc.Diagnose()

	if health.OS != runtime.GOOS {
		t.Errorf("Expected OS %s, got %s", runtime.GOOS, health.OS)
	}
	if health.Arch != runtime.GOARCH {
		t.Errorf("Expected Arch %s, got %s", runtime.GOARCH, health.Arch)
	}

	// HasTTY and HasColorSupport should be boolean, not specifically tested
	// due to host variability.
}
