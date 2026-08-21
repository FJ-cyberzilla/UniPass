package ui

import (
	"bufio"
	"strings"
	"testing"
)

func TestPromptWithValidation_ValidInput(t *testing.T) {
	input := "JohnDoe\n"
	reader := bufio.NewReader(strings.NewReader(input))
	emptyCounter := 0

	val, ok := PromptWithValidation(reader, "Prompt: ", &emptyCounter)
	if !ok || val != "JohnDoe" {
		t.Errorf("Expected 'JohnDoe', got '%s'", val)
	}
}

func TestPromptWithValidation_EmptyRetriesExceeded(t *testing.T) {
	input := "\n\n\n\n\n"
	reader := bufio.NewReader(strings.NewReader(input))
	emptyCounter := 0

	_, ok := PromptWithValidation(reader, "Prompt: ", &emptyCounter)
	if ok {
		t.Error("Expected validation to return false after exceeding empty input limit")
	}
}
