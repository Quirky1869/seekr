package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/Quirky1869/seekr/internal/tui"
)

func main() {
	m := tui.InitialModel()
	p := tea.NewProgram(
		m,
		tea.WithAltScreen(),       // full terminal, restores on exit
		tea.WithMouseCellMotion(), // optional: mouse support
	)
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "seekr: %v\n", err)
		os.Exit(1)
	}
}
