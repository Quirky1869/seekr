package tui

import (
	"os/exec"
	"runtime"
	"strings"
)

// platformCopy copies s to the system clipboard using available tools.
func platformCopy(s string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("pbcopy")
	case "windows":
		cmd = exec.Command("clip")
	default: // Linux / BSD
		if _, err := exec.LookPath("xclip"); err == nil {
			cmd = exec.Command("xclip", "-selection", "clipboard")
		} else if _, err := exec.LookPath("xsel"); err == nil {
			cmd = exec.Command("xsel", "--clipboard", "--input")
		} else if _, err := exec.LookPath("wl-copy"); err == nil {
			cmd = exec.Command("wl-copy")
		} else {
			return nil // silently skip — no clipboard tool available
		}
	}
	cmd.Stdin = strings.NewReader(s)
	return cmd.Run()
}
