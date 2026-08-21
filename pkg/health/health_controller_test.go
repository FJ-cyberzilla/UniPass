package health

import (
	"testing"
)

func TestDetectEnvironment(t *testing.T) {
	hc := NewHealthController()
	env := hc.DetectEnvironment()

	if env == "" {
		t.Errorf("Expected valid environment type, got empty string")
	}
	t.Logf("Detected Runtime Environment: %s", env)
}

func TestDiagnose(t *testing.T) {
	hc := NewHealthController()
	health := hc.Diagnose()

	if health.OS == "" || health.Arch == "" {
		t.Errorf("System health report missing OS or Architecture details")
	}
}
