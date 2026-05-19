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

	title := headerStyle.Render("█▀ █▀▀ █▀▀ █▄▀ █▀█")
	title2 := headerStyle.Render("▄█ ██▄ ██▄ █░█ █▀▄")
	subtitle := subtitleStyle.Render(m.tr.AppSubtitle)
	lang := langBadgeStyle.Render(i18n.Flag(m.lang))

	// Left: logo + subtitle  Right: lang badge
	left := lipgloss.JoinVertical(lipgloss.Left, title, title2, subtitle)
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
	rows := []string{
		m.renderInputRow(fMinDepth, tr.LabelMinDepth, m.inputs[fMinDepth].View()),
		"",
		m.renderToggleRow(fFollowSym, tr.LabelFollowSym),
		m.renderToggleRow(fNoMount, tr.LabelNoMount),
		"",
		m.renderInputRow(fRegex, tr.LabelRegex, m.inputs[fRegex].View()),
		m.renderInputRow(fExclude, tr.LabelExclude, m.inputs[fExclude].View()),
		"",
		m.renderToggleRow(fDeleteMode,
			dangerLabel(tr.LabelDeleteMode)),
	}
	return m.wrapPanel(strings.Join(rows, "\n"))
}

func dangerLabel(s string) string {
	return lipgloss.NewStyle().Foreground(lipgloss.Color(colorPink)).Bold(true).Render(s)
}

// ─── Results tab ─────────────────────────────────────────────────────────────

func (m Model) renderResultsTab() string {
	tr := m.tr
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

	status := ""
	if m.statusMsg != "" {
		status = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorGreen)).
			Background(lipgloss.Color(colorDark2)).
			Render("  " + m.statusMsg)
	}

	vp := m.resultVP.View()
	vpStyled := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(colorPurple)).
		Width(m.width - 4).
		Render(vp)

	parts := []string{header}
	if status != "" {
		parts = append(parts, status)
	}
	parts = append(parts, vpStyled)

	return lipgloss.NewStyle().
		Background(lipgloss.Color(colorDark2)).
		Padding(0, 1).
		Render(strings.Join(parts, "\n"))
}

// ─── Command bar ─────────────────────────────────────────────────────────────

func (m Model) renderCmdBar() string {
	cmd := m.buildCmdString()
	// truncate if needed
	maxW := m.width - 6
	if lipgloss.Width(cmd) > maxW && maxW > 3 {
		runes := []rune(cmd)
		cmd = string(runes[:maxW-3]) + "..."
	}
	label := lipgloss.NewStyle().
		Foreground(lipgloss.Color(colorCyan)).
		Background(lipgloss.Color(colorDark2)).
		Bold(true).
		Render("$ ")
	full := cmdBoxStyle.Width(m.width - 4).Render(label + cmd)
	return full
}

// ─── Help bar ────────────────────────────────────────────────────────────────

func (m Model) renderHelpBar() string {
	tr := m.tr
	items := []struct{ key, desc string }{
		{"F5", tr.HelpRun},
		{"F6", tr.HelpCopy},
		{"F7", tr.HelpQuit},
		{"Tab/1-4", tr.HelpTab},
		{"↑↓", tr.HelpNav},
		{"Ctrl+L", tr.HelpLang},
		{"q/Ctrl+C", tr.HelpQuit},
	}
	parts := []string{}
	for _, it := range items {
		k := helpKeyStyle.Render(it.key)
		d := lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorGrayMid)).
			Background(lipgloss.Color(colorDark2)).
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
	if m.focused == fieldID {
		lStyle = labelFocusStyle
	}
	l := lStyle.Render(label)

	var tog string
	if m.toggles[fieldID] {
		tog = toggleOnStyle.Render(" ON ")
	} else {
		tog = toggleOffStyle.Render("OFF")
	}
	row := lipgloss.JoinHorizontal(lipgloss.Top, l, "  ", tog)
	return rowStyle.Render(row)
}

func (m Model) wrapPanel(content string) string {
	return lipgloss.NewStyle().
		Background(lipgloss.Color(colorDark2)).
		Padding(1, 2).
		Width(m.width).
		Render(content)
}

// i18n label helper for delete danger zone
var _ = i18n.Get // keep import used
