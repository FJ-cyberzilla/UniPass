package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

type SystemState struct {
	LastHealthCheck time.Time `json:"last_health_check"`
}

func getMinStateFilePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ".unipass_state.json"
	}
	return filepath.Join(home, ".unipass_state.json")
}

func ShouldRunWeeklyHealthCheck() bool {
	path := getMinStateFilePath()
	data, err := os.ReadFile(path)
	if err != nil {
		return true // File missing or corrupt; run health check
	}

	var state SystemState
	if err := json.Unmarshal(data, &state); err != nil {
		return true
	}

	return time.Since(state.LastHealthCheck) > 7*24*time.Hour
}

func UpdateHealthCheckTimestamp() error {
	path := getMinStateFilePath()
	state := SystemState{LastHealthCheck: time.Now()}
	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0600)
}
