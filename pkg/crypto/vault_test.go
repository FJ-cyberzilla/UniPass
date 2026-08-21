package crypto

import (
	"bytes"
	"path/filepath"
	"testing"
)

func TestVault(t *testing.T) {
	tmpDir := t.TempDir()
	vaultPath := filepath.Join(tmpDir, "unipass.vault")
	vault := NewVault(vaultPath)

	key := make([]byte, 32)
	for i := range key {
		key[i] = byte(i + 1)
	}

	tests := []struct {
		name      string
		data      []byte
		key       []byte
		wantErr   bool
	}{
		{"Valid data", []byte("secret"), key, false},
		{"Empty data", []byte(""), key, false},
		{"Large data", bytes.Repeat([]byte("A"), 1024*1024), key, false},
		{"Invalid key length", []byte("secret"), []byte("short"), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := vault.StoreEncrypt(tt.data, tt.key)
			if (err != nil) != tt.wantErr {
				t.Fatalf("StoreEncrypt() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr {
				return
			}

			decrypted, err := vault.LoadDecrypt(tt.key)
			if err != nil {
				t.Fatalf("LoadDecrypt() error = %v", err)
			}

			if !bytes.Equal(decrypted, tt.data) {
				t.Errorf("Decrypted data does not match original")
			}
		})
	}
}

func TestZeroMemory(t *testing.T) {
	sensitive := []byte{0xFF, 0xEE, 0xDD, 0xCC}
	ZeroMemory(sensitive)

	for _, b := range sensitive {
		if b != 0 {
			t.Errorf("Expected memory byte to be 0, got %X", b)
		}
	}
}
