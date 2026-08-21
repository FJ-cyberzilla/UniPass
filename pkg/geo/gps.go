package geo

import (
	"context"
	"os/exec"
	"time"
)

// CommandRunner allows overriding the command execution path during tests
var execCommandContext = exec.CommandContext

func FetchDeviceGPS() ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	cmd := execCommandContext(ctx, "termux-location")
	return cmd.Output()
}
