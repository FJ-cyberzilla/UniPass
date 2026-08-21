package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type Vault struct {
	filePath string
}

func NewVault(filePath string) *Vault {
	return &Vault{filePath: filePath}
}

// ZeroMemory overwrites sensitive byte slices in memory to prevent cold-boot/dump extraction
func ZeroMemory(b []byte) {
	for i := range b {
		b[i] = 0
	}
}

// StoreEncrypt encrypts data with AES-GCM and atomically writes it to disk
func (v *Vault) StoreEncrypt(data []byte, key []byte) error {
	if len(key) != 32 {
		return errors.New("invalid key length: must be 32 bytes for AES-256")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return fmt.Errorf("failed to create cipher: %w", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return fmt.Errorf("failed to create GCM: %w", err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return fmt.Errorf("failed to generate nonce: %w", err)
	}

	ciphertext := gcm.Seal(nonce, nonce, data, nil)

	// Ensure destination directory exists
	dir := filepath.Dir(v.filePath)
	if err := os.MkdirAll(dir, 0700); err != nil {
		return fmt.Errorf("failed to create vault directory: %w", err)
	}

	// Atomic write: write to temp file first, then rename
	tmpFile := v.filePath + ".tmp"
	if err := os.WriteFile(tmpFile, ciphertext, 0600); err != nil {
		return fmt.Errorf("failed to write temporary vault file: %w", err)
	}

	if err := os.Rename(tmpFile, v.filePath); err != nil {
		_ = os.Remove(tmpFile)
		return fmt.Errorf("failed to commit vault file: %w", err)
	}

	return nil
}

// LoadDecrypt reads encrypted data from disk and decrypts it with AES-GCM
func (v *Vault) LoadDecrypt(key []byte) ([]byte, error) {
	if len(key) != 32 {
		return nil, errors.New("invalid key length: must be 32 bytes for AES-256")
	}

	ciphertext, err := os.ReadFile(v.filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read vault file: %w", err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("failed to create cipher: %w", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCM: %w", err)
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("vault data corrupted: payload too short")
	}

	nonce, encryptedPayload := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, encryptedPayload, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt vault: invalid key or corrupted payload")
	}

	return plaintext, nil
}
