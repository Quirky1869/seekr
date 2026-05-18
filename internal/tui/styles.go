package tui

import "github.com/charmbracelet/lipgloss"

// ── Cyberpunk 2077 palette ────────────────────────────────────────────────────
// Primary neons
var (
	colorYellow  = lipgloss.Color("#FFE400") // Night City yellow
	colorCyan    = lipgloss.Color("#00F5FF") // ICE blue
	colorMagenta = lipgloss.Color("#FF003C") // Arasaka red-pink
	colorGreen   = lipgloss.Color("#39FF14") // Netrunner green
	colorOrange  = lipgloss.Color("#FF6B00") // Corpo orange
)

// Backgrounds & neutrals
var (
	colorBg       = lipgloss.Color("#0D0D0F") // near-black
	colorBgPanel  = lipgloss.Color("#111118") // panel bg
	colorBgActive = lipgloss.Color("#1A1A2E") // focused panel
	colorBorder   = lipgloss.Color("#2A2A4A") // dim border
	colorDim      = lipgloss.Color("#3A3A5C") // dimmer text
	colorMuted    = lipgloss.Color("#6C6C9A") // muted text
	colorText     = lipgloss.Color("#C8C8E8") // body text
	colorBright   = lipgloss.Color("#EEEEFF") // bright text
)

// ── Base styles ───────────────────────────────────────────────────────────────

var baseText = lipgloss.NewStyle().
	Foreground(colorText).
	Background(colorBg)

// ── App chrome ────────────────────────────────────────────────────────────────

var styleAppTitle = lipgloss.NewStyle().
	Bold(true).
	Foreground(colorYellow).
	Background(colorBg).
	Padding(0, 1)

var styleAppSubtitle = lipgloss.NewStyle().
	Foreground(colorMuted).
	Background(colorBg).
	Italic(true)

var styleLangBadge = lipgloss.NewStyle().
	Foreground(colorCyan).
	Background(colorBg).
	Bold(true).
	Padding(0, 1)

// ── Tabs ──────────────────────────────────────────────────────────────────────

var styleTabActive = lipgloss.NewStyle().
	Bold(true).
	Foreground(colorBg).
	Background(colorYellow).
	Padding(0, 1)

var styleTabInactive = lipgloss.NewStyle().
	Foreground(colorMuted).
	Background(colorBgPanel).
	Padding(0, 1)

var styleTabBar = lipgloss.NewStyle().
	Background(colorBgPanel).
	BorderBottom(true).
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(colorBorder)

// ── Panels ────────────────────────────────────────────────────────────────────

var stylePanelFocused = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(colorYellow).
	Background(colorBgActive).
	Padding(0, 1)

var stylePanelNormal = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(colorBorder).
	Background(colorBgPanel).
	Padding(0, 1)

// ── Form fields ───────────────────────────────────────────────────────────────

var styleLabelFocused = lipgloss.NewStyle().
	Bold(true).
	Foreground(colorYellow)

var styleLabelNormal = lipgloss.NewStyle().
	Foreground(colorMuted)

var styleInputFocused = lipgloss.NewStyle().
	Foreground(colorCyan).
	Background(colorBgActive).
	Border(lipgloss.NormalBorder()).
	BorderForeground(colorYellow).
	Padding(0, 1)

var styleInputNormal = lipgloss.NewStyle().
	Foreground(colorText).
	Background(colorBgPanel).
	Border(lipgloss.NormalBorder()).
	BorderForeground(colorBorder).
	Padding(0, 1)

// Checkbox / toggle
var styleCheckOn = lipgloss.NewStyle().
	Bold(true).
	Foreground(colorGreen)

var styleCheckOff = lipgloss.NewStyle().
	Foreground(colorDim)

// Selector (←→) highlight
var styleSelectorActive = lipgloss.NewStyle().
	Bold(true).
	Foreground(colorYellow).
	Background(colorBgActive).
	Padding(0, 1)

var styleSelectorInactive = lipgloss.NewStyle().
	Foreground(colorMuted).
	Padding(0, 1)

var styleSelectorArrow = lipgloss.NewStyle().
	Bold(true).
	Foreground(colorMagenta)

// ── Results ───────────────────────────────────────────────────────────────────

var styleResultsTitle = lipgloss.NewStyle().
	Bold(true).
	Foreground(colorCyan)

var styleResultItem = lipgloss.NewStyle().
	Foreground(colorText)

var styleResultItemSelected = lipgloss.NewStyle().
	Bold(true).
	Foreground(colorYellow).
	Background(colorBgActive)

var styleResultsEmpty = lipgloss.NewStyle().
	Italic(true).
	Foreground(colorDim)

var styleResultsRunning = lipgloss.NewStyle().
	Bold(true).
	Foreground(colorOrange)

var styleResultCount = lipgloss.NewStyle().
	Bold(true).
	Foreground(colorGreen)

var styleResultError = lipgloss.NewStyle().
	Bold(true).
	Foreground(colorMagenta)

// ── Command preview ───────────────────────────────────────────────────────────

var stylePreviewTitle = lipgloss.NewStyle().
	Bold(true).
	Foreground(colorCyan)

var stylePreviewBox = lipgloss.NewStyle().
	Foreground(colorGreen).
	Background(colorBgPanel).
	Border(lipgloss.NormalBorder()).
	BorderForeground(colorBorder).
	Padding(0, 1)

var stylePreviewCopied = lipgloss.NewStyle().
	Bold(true).
	Foreground(colorGreen)

// ── Buttons ───────────────────────────────────────────────────────────────────

var styleBtnPrimary = lipgloss.NewStyle().
	Bold(true).
	Foreground(colorBg).
	Background(colorYellow).
	Padding(0, 2).
	MarginRight(1)

var styleBtnSecondary = lipgloss.NewStyle().
	Bold(true).
	Foreground(colorCyan).
	Background(colorBgPanel).
	Border(lipgloss.NormalBorder()).
	BorderForeground(colorCyan).
	Padding(0, 2).
	MarginRight(1)

var styleBtnDanger = lipgloss.NewStyle().
	Bold(true).
	Foreground(colorBg).
	Background(colorMagenta).
	Padding(0, 2).
	MarginRight(1)

// ── Footer ────────────────────────────────────────────────────────────────────

var styleFooter = lipgloss.NewStyle().
	Foreground(colorMuted).
	Background(colorBg).
	BorderTop(true).
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(colorBorder)

var styleFooterKey = lipgloss.NewStyle().
	Bold(true).
	Foreground(colorYellow)

var styleFooterSep = lipgloss.NewStyle().
	Foreground(colorBorder)

// ── Section headers ───────────────────────────────────────────────────────────

var styleSectionTitle = lipgloss.NewStyle().
	Bold(true).
	Foreground(colorCyan).
	MarginBottom(1)

// ── Divider ───────────────────────────────────────────────────────────────────

var styleDivider = lipgloss.NewStyle().
	Foreground(colorBorder)

// ── Help ──────────────────────────────────────────────────────────────────────

var styleHelpTitle = lipgloss.NewStyle().
	Bold(true).
	Foreground(colorYellow)

var styleHelpKey = lipgloss.NewStyle().
	Bold(true).
	Foreground(colorCyan).
	Width(18)

var styleHelpDesc = lipgloss.NewStyle().
	Foreground(colorText)

var styleHelpCategory = lipgloss.NewStyle().
	Bold(true).
	Foreground(colorMagenta).
	MarginTop(1)
