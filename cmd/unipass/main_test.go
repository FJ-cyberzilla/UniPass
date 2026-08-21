package main

import (
	"testing"
	"unipass/pkg/health"
)

func TestHealthFlagExecution(t *testing.T) {
	hc := health.NewHealthController()
	report := hc.Diagnose()

	if report.OS == "" {
		t.Errorf("Expected valid OS string from health controller")
	}
}
