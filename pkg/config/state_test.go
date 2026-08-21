package config

import (
	"testing"
)

func TestConfigManager(t *testing.T) {
	tmpDir := t.TempDir()
	cm := NewConfigManager(tmpDir)

	validKey := make([]byte, 32)
	for i := range validKey {
		validKey[i] = byte(i + 1)
	}

	tests := []struct {
		name    string
		state   *ConfigState
		key     []byte
		wantErr bool
	}{
		{
			"Valid state save/load",
			&ConfigState{
				Credentials: map[string]string{"key": "value"},
				City:        "London",
			},
			validKey,
			false,
		},
		{
			"Invalid key length",
			&ConfigState{},
			[]byte("short"),
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := cm.SaveState(tt.state, tt.key)
			if (err != nil) != tt.wantErr {
				t.Fatalf("SaveState() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr {
				_, err = cm.LoadState(tt.key)
				if err == nil {
					t.Errorf("Expected error on LoadState with invalid key, got nil")
				}
				return
			}

			loaded, err := cm.LoadState(tt.key)
			if err != nil {
				t.Fatalf("LoadState() error = %v", err)
			}

			if loaded.Credentials["key"] != tt.state.Credentials["key"] {
				t.Errorf("Expected credentials %v, got %v", tt.state.Credentials, loaded.Credentials)
			}
			if loaded.LastHealthCheck.IsZero() {
				t.Errorf("Expected non-zero LastHealthCheck")
			}
		})
	}
}
