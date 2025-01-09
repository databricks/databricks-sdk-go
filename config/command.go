package config

import (
	"context"
	"os/exec"
	"strings"

	"github.com/databricks/databricks-sdk-go/databricks/log"
)

// Run a command and return the output.
func runCommand(ctx context.Context, cmd string, args []string) ([]byte, error) {
	log.DebugContext(ctx, "Running command: %s %v", cmd, strings.Join(args, " "))
	return exec.Command(cmd, args...).Output()
}
