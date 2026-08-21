package crypto

import (
	"testing"
)

func TestNewRecoveryEngineKeyLength(t *testing.T) {
	validKey := "a-32-byte-secret-key-for-unipass" // 32 bytes
	shortKey := "too-short"                        // 9 bytes

	_, err := NewRecoveryEngine(validKey)
	if err != nil {
		t.Fatalf("Expected valid 32-byte key to succeed, got: %v", err)
	}

	_, err = NewRecoveryEngine(shortKey)
	if err == nil {
		t.Fatal("Expected short key to fail cipher initialization")
	}
}

func TestEncryptDecryptPayload(t *testing.T) {
	key := "a-32-byte-secret-key-for-unipass"
	engine, err := NewRecoveryEngine(key)
	if err != nil {
		t.Fatalf("Failed creating engine: %v", err)
	}

	name := "Mina"
	country := "UK"
	city := "London"
	ts := int64(1700000000000000)

	payload, err := engine.EncryptPayload(name, country, city, ts)
	if err != nil {
		t.Fatalf("Encryption failed: %v", err)
	}

	decName, decCountry, decCity, decTS, err := engine.DecryptPayload(payload)
	if err != nil {
		t.Fatalf("Decryption failed: %v", err)
	}

	if decName != name || decCountry != country || decCity != city || decTS != ts {
		t.Fatalf("Decrypted payload mismatch! Got (%s, %s, %s, %d)", decName, decCountry, decCity, decTS)
	}
}

func TestDecryptInvalidPayload(t *testing.T) {
	key := "a-32-byte-secret-key-for-unipass"
	engine, _ := NewRecoveryEngine(key)

	_, _, _, _, err := engine.DecryptPayload("invalid-base64-payload!!!")
	if err == nil {
		t.Fatal("Expected error when decrypting garbage payload")
	}
}
