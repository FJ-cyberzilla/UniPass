package geo

import (
	"context"
	"os/exec"
	"testing"
)

func TestFetchDeviceGPS_SuccessPath(t *testing.T) {
	// Override execCommandContext to return valid JSON
	origExec := execCommandContext
	defer func() { execCommandContext = origExec }()

	execCommandContext = func(ctx context.Context, name string, args ...string) *exec.Cmd {
		// Mocking a successful execution that outputs JSON
		return exec.Command("echo", `{"latitude": 12.34, "longitude": 56.78}`)
	}

	provider := &GPSProvider{}
	res, err := provider.Resolve()
	if err != nil {
		t.Fatalf("Expected success, got error: %v", err)
	}

	if res.Latitude != 12.34 || res.Longitude != 56.78 {
		t.Errorf("Expected lat 12.34, lon 56.78, got lat %f, lon %f", res.Latitude, res.Longitude)
	}
}

func TestFetchDeviceGPS_FailurePath(t *testing.T) {
	// Override execCommandContext with a dummy executable that exits immediately
	origExec := execCommandContext
	defer func() { execCommandContext = origExec }()

	execCommandContext = func(ctx context.Context, name string, args ...string) *exec.Cmd {
		return exec.CommandContext(ctx, "false")
	}

	provider := &GPSProvider{}
	_, err := provider.Resolve()
	if err == nil {
		t.Error("Expected error for non-zero exit code, got nil")
	} else {
		t.Logf("Handled non-zero exit code: %v", err)
	}
}
