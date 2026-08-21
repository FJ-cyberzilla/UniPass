package geo

import (
	"os"
	"runtime"
	"testing"
)

func TestFetchSystemIdentity(t *testing.T) {
	provider := &SystemProvider{}
	hostname, _ := os.Hostname()
	expectedSource := "system:" + runtime.GOOS + "-" + runtime.GOARCH + "-" + hostname

	got, _ := provider.Resolve()

	if got.Source != expectedSource {
		t.Errorf("Resolve().Source = %q, want %q", got.Source, expectedSource)
	}
}
