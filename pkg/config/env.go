package config

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// LoadEnv reads a .env file and sets environment variables
func LoadEnv(path string) error {
	// G304: Validate path is a regular file
	info, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("could not stat env file: %w", err)
	}
	if info.IsDir() {
		return fmt.Errorf("provided path is a directory, not a file: %s", path)
	}

	file, err := os.Open(filepath.Clean(path))
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		// G104: Handle potential error from Setenv
		if err := os.Setenv(strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])); err != nil {
			return fmt.Errorf("failed to set env var: %w", err)
		}
	}
	return scanner.Err()
}
