package crypto

import (
	"testing"
)

func TestNewRecoveryEngineKeyLength(t *testing.T) {
	engine := NewRecoveryEngine(32)
	if engine.KeyLength != 32 {
		t.Errorf("Expected key length 32, got %d", engine.KeyLength)
	}
}

func TestGenerateAndVerifyHash(t *testing.T) {
	engine := NewRecoveryEngine(32)
	payload := []byte("secret_payload_data")

	hash := engine.GenerateKeyHash(payload)
	if err := engine.VerifyPayload(payload, hash); err != nil {
		t.Fatalf("Payload verification failed: %v", err)
	}

	if err := engine.VerifyPayload(payload, "invalid_hash"); err == nil {
		t.Errorf("Expected error for invalid hash verification, got nil")
	}
}
