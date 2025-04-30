package git

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// GetStagedDiff runs `git diff --cached --no-color` and returns the diff as a string.
// If there are no staged changes, it returns an empty string without error.
func GetStagedDiff() (string, error) {
	cmd := exec.Command("git", "diff", "--cached", "--no-color")

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to run git diff: %w", err)
	}

	diff := strings.TrimSpace(out.String())
	return diff, nil
}
