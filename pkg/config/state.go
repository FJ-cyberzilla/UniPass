package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"unipass/pkg/crypto"
)

type ConfigState struct {
	LastHealthCheck time.Time         `json:"last_health_check"`
	Credentials     map[string]string `json:"credentials"`
}

type ConfigManager struct {
	configPath string
	vault      *crypto.Vault
}

func NewConfigManager(baseDir string) *ConfigManager {
	configPath := filepath.Join(baseDir, "unipass.json")
	vaultPath := filepath.Join(baseDir, "credentials.vault")
	return &ConfigManager{
		configPath: configPath,
		vault:      crypto.NewVault(vaultPath),
	}
}

// EvaluateWeeklyHealthCheck returns true if a health check interval has elapsed (>7 days)
func EvaluateWeeklyHealthCheck(lastCheck time.Time) bool {
	if lastCheck.IsZero() {
		return true
	}
	return time.Since(lastCheck) > 7*24*time.Hour
}

// UpdateHealthCheckTimestamp refreshes the health check timestamp
func (cm *ConfigManager) UpdateHealthCheckTimestamp(state *ConfigState) {
	state.LastHealthCheck = time.Now()
}

func (cm *ConfigManager) SaveState(state *ConfigState, masterKey []byte) error {
	if len(masterKey) != 32 {
		return errors.New("invalid master key: must be 32 bytes")
	}

	cm.UpdateHealthCheckTimestamp(state)

	if len(state.Credentials) > 0 {
		credData, err := json.Marshal(state.Credentials)
		if err != nil {
			return fmt.Errorf("failed to serialize credentials: %w", err)
		}
		defer crypto.ZeroMemory(credData)

		if err := cm.vault.StoreEncrypt(credData, masterKey); err != nil {
			return fmt.Errorf("failed to store credentials in vault: %w", err)
		}
	}

	metadataState := ConfigState{
		LastHealthCheck: state.LastHealthCheck,
		Credentials:     nil,
	}

	data, err := json.MarshalIndent(metadataState, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config metadata: %w", err)
	}

	if err := os.MkdirAll(filepath.Dir(cm.configPath), 0700); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	return os.WriteFile(cm.configPath, data, 0600)
}

func (cm *ConfigManager) LoadState(masterKey []byte) (*ConfigState, error) {
	if len(masterKey) != 32 {
		return nil, errors.New("invalid master key: must be 32 bytes")
	}

	var state ConfigState
	if data, err := os.ReadFile(cm.configPath); err == nil {
		_ = json.Unmarshal(data, &state)
	}

	state.Credentials = make(map[string]string)

	decryptedCreds, err := cm.vault.LoadDecrypt(masterKey)
	if err == nil {
		defer crypto.ZeroMemory(decryptedCreds)
		var creds map[string]string
		if err := json.Unmarshal(decryptedCreds, &creds); err == nil {
			state.Credentials = creds
		}
	}

	return &state, nil
}
