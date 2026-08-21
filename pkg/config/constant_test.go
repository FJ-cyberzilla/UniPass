package config

import (
	"os"
	"testing"
)

func TestHealthCheckTimestamp(t *testing.T) {
	tmpDir := t.TempDir()
	origHome := os.Getenv("HOME")
	os.Setenv("HOME", tmpDir)
	defer os.Setenv("HOME", origHome)

	// Test 1: Should run when state file doesn't exist
	if !ShouldRunWeeklyHealthCheck() {
		t.Errorf("Expected ShouldRunWeeklyHealthCheck to return true when state file is missing")
	}

	// Test 2: Update timestamp and verify it shouldn't run immediately
	if err := UpdateHealthCheckTimestamp(); err != nil {
		t.Fatalf("Failed to update timestamp: %v", err)
	}

	if ShouldRunWeeklyHealthCheck() {
		t.Errorf("Expected ShouldRunWeeklyHealthCheck to return false right after updating")
	}
}
