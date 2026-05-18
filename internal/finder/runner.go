package finder

import (
	"bufio"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
)

// LinesMsg carries all find output lines once the command finishes.
type LinesMsg []string

// DoneMsg signals that find has finished with an error.
type DoneMsg struct {
	Err error
}

// RunSync executes find synchronously and returns a tea.Msg.
func RunSync(opts Options) tea.Msg {
	args := Build(opts)
	cmd := exec.Command(args[0], args[1:]...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return DoneMsg{Err: err}
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return DoneMsg{Err: err}
	}

	if err := cmd.Start(); err != nil {
		return DoneMsg{Err: err}
	}

	go func() {
		s := bufio.NewScanner(stderr)
		for s.Scan() {
		}
	}()

	var lines []string
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	_ = cmd.Wait()
	return LinesMsg(lines)
}

// Run returns a BubbleTea Cmd that executes find asynchronously.
func Run(opts Options) tea.Cmd {
	return func() tea.Msg {
		return RunSync(opts)
	}
}
