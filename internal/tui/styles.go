package tui

import "github.com/charmbracelet/lipgloss"

// Cyberpunk 2077 palette
const (
	colorYellow  = "#FCE300" // Neon yellow — primary accent
	colorCyan    = "#00F5FF" // Neon cyan — secondary
	colorPink    = "#FF2079" // Neon pink — danger / delete
	colorPurple  = "#BD00FF" // Deep purple — borders active
	colorDark    = "#0D0D0F" // Near black background
	colorDark2   = "#13131A" // Panel background
	colorDark3   = "#1A1A28" // Slightly lighter panel
	colorGray    = "#3A3A55" // Muted border / inactive
	colorGrayMid = "#6060AA" // Mid gray text
	colorWhite   = "#E8E8F0" // Text
	colorGreen   = "#39FF14" // Neon green — success
	colorOrange  = "#FF6B00" // Warning
)

// ─── Base ──────────────────────────────────────────────────────────────────

var (
	baseStyle = lipgloss.NewStyle().
			Background(lipgloss.Color(colorDark))

	// App container — full terminal window
	appStyle = lipgloss.NewStyle().
			Background(lipgloss.Color(colorDark))
)

// ─── Header ─────────────────────────────────────────────────────────────────

var (
	headerStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(colorYellow)).
			Background(lipgloss.Color(colorDark)).
			PaddingLeft(1)

	subtitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorCyan)).
			Background(lipgloss.Color(colorDark)).
			PaddingLeft(1)

	langBadgeStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorDark)).
			Background(lipgloss.Color(colorYellow)).
			Bold(true).
			PaddingLeft(1).
			PaddingRight(1)

	headerSepStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorYellow)).
			Background(lipgloss.Color(colorDark))
)

// ─── Tabs ───────────────────────────────────────────────────────────────────

var (
	tabActiveStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(colorDark)).
			Background(lipgloss.Color(colorYellow)).
			PaddingLeft(1).
			PaddingRight(1)

	tabInactiveStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color(colorGrayMid)).
				Background(lipgloss.Color(colorDark2)).
				PaddingLeft(1).
				PaddingRight(1)

	tabBarStyle = lipgloss.NewStyle().
			Background(lipgloss.Color(colorDark2))
)

// ─── Panel / Box ────────────────────────────────────────────────────────────

var (
	panelStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(colorPurple)).
			Background(lipgloss.Color(colorDark2)).
			Padding(0, 1)

	panelTitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(colorYellow)).
			Background(lipgloss.Color(colorDark2))
)

// ─── Form fields ────────────────────────────────────────────────────────────

var (
	labelStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorCyan)).
			Background(lipgloss.Color(colorDark2)).
			Bold(true).
			Width(20)

	labelFocusStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorYellow)).
			Background(lipgloss.Color(colorDark2)).
			Bold(true).
			Width(20)

	inputActiveStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color(colorWhite)).
				Background(lipgloss.Color(colorDark3)).
				Border(lipgloss.NormalBorder()).
				BorderForeground(lipgloss.Color(colorYellow))

	inputInactiveStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color(colorGrayMid)).
				Background(lipgloss.Color(colorDark3)).
				Border(lipgloss.NormalBorder()).
				BorderForeground(lipgloss.Color(colorGray))

	// Toggle ON/OFF
	toggleOnStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorDark)).
			Background(lipgloss.Color(colorGreen)).
			Bold(true).
			PaddingLeft(1).
			PaddingRight(1)

	toggleOffStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorGrayMid)).
			Background(lipgloss.Color(colorDark3)).
			PaddingLeft(1).
			PaddingRight(1)

	// Select options
	selectActiveStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color(colorDark)).
				Background(lipgloss.Color(colorCyan)).
				Bold(true).
				PaddingLeft(1).
				PaddingRight(1)

	selectInactiveStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color(colorGrayMid)).
				Background(lipgloss.Color(colorDark3)).
				PaddingLeft(1).
				PaddingRight(1)

	rowStyle = lipgloss.NewStyle().
			Background(lipgloss.Color(colorDark2)).
			PaddingTop(0).
			PaddingBottom(0)
)

// ─── Results ────────────────────────────────────────────────────────────────

var (
	cmdBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(colorCyan)).
			Background(lipgloss.Color(colorDark3)).
			Foreground(lipgloss.Color(colorYellow)).
			Bold(true).
			Padding(0, 1)

	resultItemStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorWhite)).
			Background(lipgloss.Color(colorDark2))

	resultItemSelectedStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color(colorDark)).
				Background(lipgloss.Color(colorCyan)).
				Bold(true)

	resultCountStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color(colorGreen)).
				Background(lipgloss.Color(colorDark2)).
				Bold(true)

	statusRunningStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color(colorOrange)).
				Background(lipgloss.Color(colorDark2)).
				Bold(true)

	statusEmptyStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color(colorGrayMid)).
				Background(lipgloss.Color(colorDark2)).
				Italic(true)

	errorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorPink)).
			Background(lipgloss.Color(colorDark2)).
			Bold(true)
)

// ─── Help bar ───────────────────────────────────────────────────────────────

var (
	helpBarStyle = lipgloss.NewStyle().
			Background(lipgloss.Color(colorDark2)).
			Foreground(lipgloss.Color(colorGrayMid)).
			PaddingLeft(1)

	helpKeyStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorYellow)).
			Background(lipgloss.Color(colorDark2)).
			Bold(true)

	helpSepStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorGray)).
			Background(lipgloss.Color(colorDark2))
)

// ─── Buttons ────────────────────────────────────────────────────────────────

var (
	btnRunStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorDark)).
			Background(lipgloss.Color(colorYellow)).
			Bold(true).
			PaddingLeft(1).
			PaddingRight(1)

	btnCopyStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorDark)).
			Background(lipgloss.Color(colorCyan)).
			Bold(true).
			PaddingLeft(1).
			PaddingRight(1)

	btnClearStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorWhite)).
			Background(lipgloss.Color(colorGray)).
			PaddingLeft(1).
			PaddingRight(1)

	btnDangerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorWhite)).
			Background(lipgloss.Color(colorPink)).
			Bold(true).
			PaddingLeft(1).
			PaddingRight(1)
)

// ─── Helpers ────────────────────────────────────────────────────────────────

// Repeat a string n times
func repeat(s string, n int) string {
	out := ""
	for i := 0; i < n; i++ {
		out += s
	}
	return out
}
