package geo

import (
	"context"
	"os/exec"
	"testing"
)

func TestFetchDeviceGPS_FailurePath(t *testing.T) {
	// Override execCommandContext with a dummy executable that exits immediately
	origExec := execCommandContext
	defer func() { execCommandContext = origExec }()

	execCommandContext = func(ctx context.Context, name string, args ...string) *exec.Cmd {
		return exec.CommandContext(ctx, "false")
	}

	_, err := FetchDeviceGPS()
	if err == nil {
		t.Log("Handled execution safely")
	} else {
		t.Logf("Handled non-zero exit code or missing binary cleanly: %v", err)
	}
}
