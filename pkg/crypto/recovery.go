package crypto

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
)

type RecoveryEngine struct {
	KeyLength int
}

func NewRecoveryEngine(keyLength int) *RecoveryEngine {
	if keyLength <= 0 {
		keyLength = 32
	}
	return &RecoveryEngine{KeyLength: keyLength}
}

func (re *RecoveryEngine) GenerateKeyHash(payload []byte) string {
	hash := sha256.Sum256(payload)
	return hex.EncodeToString(hash[:])
}

func (re *RecoveryEngine) VerifyPayload(payload []byte, expectedHash string) error {
	if re.GenerateKeyHash(payload) != expectedHash {
		return errors.New("checksum verification failed: payload hash mismatch")
	}
	return nil
}
