package crypto

import (
	"bytes"
	"path/filepath"
	"testing"
)

func TestVaultStoreAndLoad(t *testing.T) {
	tmpDir := t.TempDir()
	vaultPath := filepath.Join(tmpDir, "unipass.vault")
	vault := NewVault(vaultPath)

	masterKey := make([]byte, 32)
	for i := range masterKey {
		masterKey[i] = byte(i + 1)
	}

	sensitiveData := []byte("secret_master_seed_hash_value")

	// Test Store
	if err := vault.StoreEncrypt(sensitiveData, masterKey); err != nil {
		t.Fatalf("Vault store failed: %v", err)
	}

	// Test Load
	decrypted, err := vault.LoadDecrypt(masterKey)
	if err != nil {
		t.Fatalf("Vault load failed: %v", err)
	}

	if !bytes.Equal(decrypted, sensitiveData) {
		t.Errorf("Expected decrypted data %s, got %s", sensitiveData, decrypted)
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

func TestVaultInvalidKeyLength(t *testing.T) {
	tmpDir := t.TempDir()
	vault := NewVault(filepath.Join(tmpDir, "test.vault"))
	invalidKey := []byte("short_key")

	if err := vault.StoreEncrypt([]byte("test"), invalidKey); err == nil {
		t.Errorf("Expected error for invalid key length, got nil")
	}
}
