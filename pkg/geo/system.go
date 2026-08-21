package geo

import (
	"os"
	"runtime"
)

// FetchSystemIdentity safely collects non-malicious system identifiers as a last-resort fallback
func FetchSystemIdentity() string {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown-host"
	}
	return runtime.GOOS + "-" + runtime.GOARCH + "-" + hostname
}
