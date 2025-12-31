//go:build haiku

package clipboard

import (
	"os/exec"
	"strings"
)

func readAll() (string, error) {
	// Haiku's native 'clipboard' tool: -p prints the current clipboard
	out, err := exec.Command("clipboard", "-p").Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func writeAll(text string) error {
	// Haiku's native 'clipboard' tool: reads from stdin by default
	cmd := exec.Command("clipboard", "-i")
	cmd.Stdin = strings.NewReader(text)
	return cmd.Run()
}

