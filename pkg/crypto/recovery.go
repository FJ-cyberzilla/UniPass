package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"strconv"
	"strings"

	"unipass/pkg/config"
)

type RecoveryEngine struct {
	key []byte
}

func NewRecoveryEngine(secretKey string) (*RecoveryEngine, error) {
	if len(secretKey) != config.AESKeyLength {
		return nil, fmt.Errorf("secret key length must be exactly %d bytes", config.AESKeyLength)
	}
	return &RecoveryEngine{key: []byte(secretKey)}, nil
}

func (r *RecoveryEngine) EncryptPayload(name, country, city string, timestampMicroseconds int64) (string, error) {
	plainText := fmt.Sprintf("%s:%s:%s:%d", name, country, city, timestampMicroseconds)

	block, err := aes.NewCipher(r.key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	cipherText := gcm.Seal(nonce, nonce, []byte(plainText), nil)
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func (r *RecoveryEngine) DecryptPayload(payload string) (string, string, string, int64, error) {
	data, err := base64.StdEncoding.DecodeString(payload)
	if err != nil {
		return "", "", "", 0, fmt.Errorf("invalid base64 structure: %w", err)
	}

	block, err := aes.NewCipher(r.key)
	if err != nil {
		return "", "", "", 0, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", "", "", 0, err
	}

	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return "", "", "", 0, fmt.Errorf("corrupted recovery payload")
	}

	nonce, cipherText := data[:nonceSize], data[nonceSize:]
	plainText, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", "", "", 0, fmt.Errorf("decryption rejected (tampered payload or invalid key)")
	}

	parts := strings.Split(string(plainText), ":")
	if len(parts) != 4 {
		return "", "", "", 0, fmt.Errorf("malformed payload string")
	}

	ts, _ := strconv.ParseInt(parts[3], 10, 64)
	return parts[0], parts[1], parts[2], ts, nil
}
