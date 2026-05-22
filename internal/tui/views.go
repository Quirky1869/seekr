package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/Quirky1869/seekr/internal/i18n"
)

// ─── View entry point ─────────────────────────────────────────────────────────

func (m Model) View() string {
	if m.width == 0 {
		return "Loading SEEKR..."
	}

	sections := []string{
		m.renderHeader(),
		m.renderTabBar(),
		m.renderBody(),
		m.renderCmdBar(),
		m.renderHelpBar(),
	}

	full := strings.Join(sections, "\n")

	// Clip to terminal height to avoid overflow
	lines := strings.Split(full, "\n")
	if len(lines) > m.height {
		lines = lines[:m.height]
	}
	return strings.Join(lines, "\n")
}

// ─── Header ──────────────────────────────────────────────────────────────────

func (m Model) renderHeader() string {
	w := m.width

	title1 := headerStyle.Render("█▀ █▀▀ █▀▀ █▄▀ █▀█")
	title2 := headerStyle.Render("▄█ ██▄ ██▄ █░█ █▀▄")
	subtitle := subtitleStyle.Render(m.tr.AppSubtitle)
	lang := langBadgeStyle.Render(i18n.Flag(m.lang))

	// Left: logo + subtitle  Right: lang badge
	left := lipgloss.JoinVertical(lipgloss.Left, title1, title2, subtitle)
	right := lipgloss.NewStyle().Width(12).Align(lipgloss.Right).Render(lang)

	spacer := w - lipgloss.Width(left) - lipgloss.Width(right)
	if spacer < 0 {
		spacer = 0
	}
	gap := strings.Repeat(" ", spacer)

	header := appStyle.Width(w).Render(left + gap + right)

	sep := headerSepStyle.Render(repeat("─", w))
	return header + "\n" + sep
}

// ─── Tab bar ─────────────────────────────────────────────────────────────────

func (m Model) renderTabBar() string {
	labels := []string{
		m.tr.TabBasic,
		m.tr.TabFilters,
		m.tr.TabAdvanced,
		m.tr.TabResults,
	}
	tabs := []string{}
	for i, l := range labels {
		if i == m.tab {
			tabs = append(tabs, tabActiveStyle.Render(l))
		} else {
			tabs = append(tabs, tabInactiveStyle.Render(l))
		}
	}
	bar := tabBarStyle.Width(m.width).Render(strings.Join(tabs, " "))
	sep := headerSepStyle.Render(repeat("─", m.width))
	return bar + "\n" + sep
}

// ─── Body ─────────────────────────────────────────────────────────────────────

func (m Model) renderBody() string {
	switch m.tab {
	case TabBasic:
		return m.renderBasicTab()
	case TabFilters:
		return m.renderFiltersTab()
	case TabAdvanced:
		return m.renderAdvancedTab()
	case TabResults:
		return m.renderResultsTab()
	}
	return ""
}

// ─── Basic tab ───────────────────────────────────────────────────────────────

func (m Model) renderBasicTab() string {
	tr := m.tr
	rows := []string{
		m.renderInputRow(fPath, tr.LabelStartPath, m.inputs[fPath].View()),
		m.renderInputRow(fName, tr.LabelFileName, m.inputs[fName].View()),
		m.renderSelectRow(fType, tr.LabelFileType, m.renderTypeSelect()),
		m.renderInputRow(fMaxDepth, tr.LabelMaxDepth, m.inputs[fMaxDepth].View()),
	}
	return m.wrapPanel(strings.Join(rows, "\n"))
}

func (m Model) renderTypeSelect() string {
	labels := []string{
		m.tr.TypeAny,
		m.tr.TypeFile,
		m.tr.TypeDirectory,
		m.tr.TypeSymlink,
	}
	parts := []string{}
	for i, l := range labels {
		if i == m.fileTypeIdx {
			parts = append(parts, selectActiveStyle.Render(l))
		} else {
			parts = append(parts, selectInactiveStyle.Render(l))
		}
	}
	return strings.Join(parts, " ")
}

// ─── Filters tab ─────────────────────────────────────────────────────────────

func (m Model) renderFiltersTab() string {
	tr := m.tr
	rows := []string{
		m.renderInputRow(fMtime, tr.LabelMtime, m.inputs[fMtime].View()),
		m.renderInputRow(fSize, tr.LabelSize, m.inputs[fSize].View()+" "+
			lipgloss.NewStyle().Foreground(lipgloss.Color(colorGrayMid)).Render(tr.SizeUnit)),
		m.renderInputRow(fPerm, tr.LabelPerm, m.inputs[fPerm].View()),
		m.renderInputRow(fOwner, tr.LabelOwner, m.inputs[fOwner].View()),
		"",
		m.renderToggleRow(fEmpty, tr.LabelEmpty),
		m.renderToggleRow(fExecutable, tr.LabelExecutable),
		m.renderToggleRow(fReadable, tr.LabelReadable),
		m.renderToggleRow(fWritable, tr.LabelWritable),
	}
	return m.wrapPanel(strings.Join(rows, "\n"))
}

// ─── Advanced tab ────────────────────────────────────────────────────────────

func (m Model) renderAdvancedTab() string {
	tr := m.tr
	var deleteLabel string
	if m.focused == fDeleteMode {
		deleteLabel = lipgloss.NewStyle().Foreground(lipgloss.Color(colorYellow)).Bold(true).Render(tr.LabelDeleteMode)
	} else {
		deleteLabel = dangerLabel(tr.LabelDeleteMode)
	}

	rows := []string{
		m.renderInputRow(fMinDepth, tr.LabelMinDepth, m.inputs[fMinDepth].View()),
		"",
		m.renderToggleRow(fFollowSym, tr.LabelFollowSym),
		m.renderToggleRow(fNoMount, tr.LabelNoMount),
		"",
		m.renderInputRow(fRegex, tr.LabelRegex, m.inputs[fRegex].View()),
		m.renderInputRow(fExclude, tr.LabelExclude, m.inputs[fExclude].View()),
		"",
		m.renderToggleRow(fGrepMode, tr.LabelGrepMode),
	}

	if m.toggles[fGrepMode] {
		rows = append(rows,
			m.renderSubInputRow(fGrepPattern, tr.LabelGrepPattern, m.inputs[fGrepPattern].View()),
			m.renderSubToggleRow(fGrepIgnoreCase, tr.LabelGrepIgnoreCase),
		)
	}

	rows = append(rows, "", m.renderToggleRow(fDeleteMode, deleteLabel))

	return m.wrapPanel(strings.Join(rows, "\n"))
}

func dangerLabel(s string) string {
	return lipgloss.NewStyle().Foreground(lipgloss.Color(colorPink)).Bold(true).Render(s)
}

// ─── Results tab ─────────────────────────────────────────────────────────────

func (m Model) renderResultsTab() string {
	tr := m.tr

	if m.toggles[fGrepMode] {
		return m.renderSplitResultsTab()
	}

	var header string
	switch {
	case m.running:
		header = statusRunningStyle.Render(tr.LabelRunning)
	case m.errMsg != "":
		header = errorStyle.Render("⚠  " + m.errMsg)
	case len(m.results) == 0:
		header = statusEmptyStyle.Render(tr.LabelNoResults)
	default:
		count := fmt.Sprintf("  %d %s", len(m.results), tr.LabelResultCount)
		header = resultCountStyle.Render(count)
	}

	vp := m.resultVP.View()
	vpStyled := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(colorPurple)).
		Width(m.width - 4).
		Render(vp)

	return lipgloss.NewStyle().
		Padding(0, 1).
		Render(strings.Join([]string{header, vpStyled}, "\n"))
}

func (m Model) renderSplitResultsTab() string {
	tr := m.tr

	// ── left header
	var leftHeader string
	switch {
	case m.running:
		leftHeader = statusRunningStyle.Render(tr.LabelRunning)
	case m.errMsg != "":
		leftHeader = errorStyle.Render("⚠  " + m.errMsg)
	case len(m.results) == 0:
		leftHeader = statusEmptyStyle.Render(tr.LabelNoResults)
	default:
		leftHeader = resultCountStyle.Render(fmt.Sprintf("  %d %s", len(m.results), tr.LabelResultCount))
	}

	// ── right header
	var rightHeader string
	switch {
	case m.running:
		rightHeader = statusRunningStyle.Render(tr.LabelRunning)
	case len(m.occurrences) == 0:
		rightHeader = statusEmptyStyle.Render(tr.LabelNoOccurrences)
	default:
		rightHeader = resultCountStyle.Render(fmt.Sprintf("  %d %s", len(m.occurrences), tr.LabelOccurrenceCount))
	}

	halfW := m.resultVP.Width

	leftVP := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(colorPurple)).
		Width(halfW).
		Render(m.resultVP.View())

	rightVP := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(colorPurple)).
		Width(halfW).
		Render(m.occurrenceVP.View())

	leftPanel := lipgloss.JoinVertical(lipgloss.Left, leftHeader, leftVP)
	rightPanel := lipgloss.JoinVertical(lipgloss.Left, rightHeader, rightVP)

	split := lipgloss.JoinHorizontal(lipgloss.Top, leftPanel, "  ", rightPanel)

	return lipgloss.NewStyle().Padding(0, 1).Render(split)
}

// ─── Command bar ─────────────────────────────────────────────────────────────

func (m Model) renderCmdBar() string {
	statusPart := ""
	if m.statusMsg != "" {
		statusPart = "  " + lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorGreen)).
			Bold(true).
			Render(m.statusMsg)
	}

	cmd := m.buildCmdString()
	maxW := m.width - 6 - lipgloss.Width(statusPart)
	if lipgloss.Width(cmd) > maxW && maxW > 3 {
		runes := []rune(cmd)
		cmd = string(runes[:maxW-3]) + "..."
	}
	label := lipgloss.NewStyle().
		Foreground(lipgloss.Color(colorCyan)).
		Bold(true).
		Render("$ ")
	full := cmdBoxStyle.Width(m.width - 4).Render(label + cmd + statusPart)
	return full
}

// ─── Help bar ────────────────────────────────────────────────────────────────

func (m Model) renderHelpBar() string {
	tr := m.tr
	items := []struct{ key, desc string }{
		{"F5", tr.HelpRun},
		{"F6", tr.HelpCopy},
		{"F7", tr.HelpQuit},
		{"Tab/F1-F4", tr.HelpTab},
		{"↑↓", tr.HelpNav},
		{"Ctrl+L", tr.HelpLang},
		{"q/Ctrl+C", tr.HelpQuit},
	}
	parts := []string{}
	for _, it := range items {
		k := helpKeyStyle.Render(it.key)
		d := lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorGrayMid)).
			Render(" " + it.desc)
		sep := helpSepStyle.Render("  │  ")
		parts = append(parts, k+d+sep)
	}
	bar := helpBarStyle.Width(m.width).Render(strings.Join(parts, ""))
	return bar
}

// ─── Row helpers ──────────────────────────────────────────────────────────────

func (m Model) renderInputRow(fieldID int, label, input string) string {
	lStyle := labelStyle
	if m.focused == fieldID {
		lStyle = labelFocusStyle
	}
	l := lStyle.Render(label)
	row := lipgloss.JoinHorizontal(lipgloss.Top, l, "  ", input)
	return rowStyle.Render(row)
}

func (m Model) renderSelectRow(fieldID int, label, content string) string {
	lStyle := labelStyle
	if m.focused == fieldID {
		lStyle = labelFocusStyle
	}
	l := lStyle.Render(label)
	hint := lipgloss.NewStyle().
		Foreground(lipgloss.Color(colorGrayMid)).
		Render(" ← Enter →")
	row := lipgloss.JoinHorizontal(lipgloss.Top, l, "  ", content, hint)
	return rowStyle.Render(row)
}

func (m Model) renderToggleRow(fieldID int, label string) string {
	lStyle := labelStyle
	isFocused := m.focused == fieldID

	if isFocused && fieldID != fDeleteMode {
		lStyle = labelFocusStyle
	}

	var l string
	if fieldID == fDeleteMode {
		l = labelStyle.Width(20).Render(label)
	} else {
		l = lStyle.Render(label)
	}

	var tog string
	if m.toggles[fieldID] {
		if isFocused {
			tog = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#000000")).
				Background(lipgloss.Color(colorYellow)).
				Bold(true).
				Width(5).
				Align(lipgloss.Center).
				Render("ON")
		} else {
			tog = toggleOnStyle.Render("ON")
		}
	} else {
		if isFocused {
			tog = lipgloss.NewStyle().
				Foreground(lipgloss.Color(colorYellow)).
				Bold(true).
				Width(5).
				Align(lipgloss.Center).
				Render("OFF")
		} else {
			tog = toggleOffStyle.Render("OFF")
		}
	}
	row := lipgloss.JoinHorizontal(lipgloss.Top, l, "  ", tog)
	return rowStyle.Render(row)
}

func (m Model) renderSubInputRow(fieldID int, label, input string) string {
	lStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(colorGrayMid)).Bold(true).Width(24)
	if m.focused == fieldID {
		lStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(colorCyan)).Bold(true).Width(24)
	}
	l := lStyle.Render(label)
	row := lipgloss.JoinHorizontal(lipgloss.Top, l, "  ", input)
	return rowStyle.Render(row)
}

func (m Model) renderSubToggleRow(fieldID int, label string) string {
	isFocused := m.focused == fieldID
	lStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(colorGrayMid)).Bold(true).Width(24)
	if isFocused {
		lStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(colorCyan)).Bold(true).Width(24)
	}
	l := lStyle.Render(label)
	var tog string
	if m.toggles[fieldID] {
		if isFocused {
			tog = lipgloss.NewStyle().Foreground(lipgloss.Color("#000000")).Background(lipgloss.Color(colorYellow)).Bold(true).Width(5).Align(lipgloss.Center).Render("ON")
		} else {
			tog = toggleOnStyle.Render("ON")
		}
	} else {
		if isFocused {
			tog = lipgloss.NewStyle().Foreground(lipgloss.Color(colorYellow)).Bold(true).Width(5).Align(lipgloss.Center).Render("OFF")
		} else {
			tog = toggleOffStyle.Render("OFF")
		}
	}
	row := lipgloss.JoinHorizontal(lipgloss.Top, l, "  ", tog)
	return rowStyle.Render(row)
}

func (m Model) wrapPanel(content string) string {
	return lipgloss.NewStyle().
		Padding(1, 2).
		Width(m.width).
		Render(content)
}

// i18n label helper for delete danger zone
var _ = i18n.Get // keep import used
