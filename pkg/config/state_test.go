package config

import (
	"path/filepath"
	"testing"
)

func TestConfigManager_SaveAndLoadWithVault(t *testing.T) {
	tmpDir := t.TempDir()
	cm := NewConfigManager(tmpDir)

	masterKey := make([]byte, 32)
	for i := range masterKey {
		masterKey[i] = byte(i + 42)
	}

	initialState := &ConfigState{
		Credentials: map[string]string{
			"primary_seed":  "super_secret_seed_value",
			"recovery_hash": "a1b2c3d4e5f6",
		},
	}

	// 1. Test Save
	if err := cm.SaveState(initialState, masterKey); err != nil {
		t.Fatalf("Failed to save state with vault: %v", err)
	}

	// 2. Test Load
	loadedState, err := cm.LoadState(masterKey)
	if err != nil {
		t.Fatalf("Failed to load state from vault: %v", err)
	}

	if loadedState.Credentials["primary_seed"] != "super_secret_seed_value" {
		t.Errorf("Expected primary_seed 'super_secret_seed_value', got '%s'", loadedState.Credentials["primary_seed"])
	}

	if loadedState.LastHealthCheck.IsZero() {
		t.Errorf("Expected non-zero LastHealthCheck timestamp")
	}
}

func TestConfigManager_InvalidKey(t *testing.T) {
	tmpDir := t.TempDir()
	cm := NewConfigManager(filepath.Join(tmpDir, "invalid"))
	shortKey := []byte("invalid_key")

	if err := cm.SaveState(&ConfigState{}, shortKey); err == nil {
		t.Errorf("Expected error for invalid key length on save, got nil")
	}

	if _, err := cm.LoadState(shortKey); err == nil {
		t.Errorf("Expected error for invalid key length on load, got nil")
	}
}
