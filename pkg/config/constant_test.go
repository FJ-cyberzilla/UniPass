package config

import (
	"testing"
	"time"
)

func TestEvaluateWeeklyHealthCheck(t *testing.T) {
	if !EvaluateWeeklyHealthCheck(time.Time{}) {
		t.Errorf("Expected true for zero-value time")
	}

	recent := time.Now().Add(-24 * time.Hour)
	if EvaluateWeeklyHealthCheck(recent) {
		t.Errorf("Expected false for health check conducted 24 hours ago")
	}

	old := time.Now().Add(-8 * 24 * time.Hour)
	if !EvaluateWeeklyHealthCheck(old) {
		t.Errorf("Expected true for health check conducted 8 days ago")
	}
}
