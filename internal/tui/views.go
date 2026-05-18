package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/Quirky1869/seekr/internal/finder"
	"github.com/Quirky1869/seekr/internal/i18n"
)

// renderRoot is the main View function.
func renderRoot(m Model) string {
	t := i18n.Get(m.lang)

	header  := renderHeader(m, t)
	tabBar  := renderTabBar(m, t)
	content := renderContent(m, t)
	footer  := renderFooter(m, t)

	return lipgloss.JoinVertical(lipgloss.Left,
		header,
		tabBar,
		content,
		footer,
	)
}

// ── Header ────────────────────────────────────────────────────────────────────

func renderHeader(m Model, t i18n.T) string {
	title := styleAppTitle.Render("▸▸ SEEKR")
	sub   := styleAppSubtitle.Render(t.AppSubtitle)
	lang  := styleLangBadge.Render(fmt.Sprintf("%s %s", i18n.Flag(m.lang), strings.ToUpper(string(m.lang))))

	left  := lipgloss.JoinHorizontal(lipgloss.Center, title, "  ", sub)
	right := lang

	gap := m.width - lipgloss.Width(left) - lipgloss.Width(right)
	if gap < 0 {
		gap = 0
	}
	spacer := strings.Repeat(" ", gap)

	line1 := lipgloss.NewStyle().Background(colorBg).Render(left + spacer + right)
	sep   := styleDivider.Render(strings.Repeat("─", m.width))
	return lipgloss.JoinVertical(lipgloss.Left, line1, sep)
}

// ── Tab bar ───────────────────────────────────────────────────────────────────

func renderTabBar(m Model, t i18n.T) string {
	tabs := []string{t.TabSearch, t.TabOptions, t.TabPreview, t.TabHelp}
	var rendered []string
	for i, tab := range tabs {
		if tabID(i) == m.activeTab {
			rendered = append(rendered, styleTabActive.Render(tab))
		} else {
			rendered = append(rendered, styleTabInactive.Render(tab))
		}
	}
	bar := lipgloss.JoinHorizontal(lipgloss.Center, rendered...)
	sep := styleDivider.Render(strings.Repeat("─", m.width))
	return lipgloss.JoinVertical(lipgloss.Left, bar, sep)
}

// ── Content dispatcher ────────────────────────────────────────────────────────

func renderContent(m Model, t i18n.T) string {
	contentH := m.height - 8 // header(2) + tabbar(2) + footer(2) + margin(2)
	if contentH < 4 {
		contentH = 4
	}
	switch m.activeTab {
	case tabSearch:
		return renderSearchTab(m, t, contentH)
	case tabOptions:
		return renderOptionsTab(m, t, contentH)
	case tabPreview:
		return renderPreviewTab(m, t, contentH)
	case tabHelp:
		return renderHelpTab(m, t, contentH)
	}
	return ""
}

// ── Search tab ────────────────────────────────────────────────────────────────

func renderSearchTab(m Model, t i18n.T, height int) string {
	leftW  := m.width/2 - 2
	rightW := m.width - leftW - 4

	// ── Left: form ──
	form := renderForm(m, t, leftW)

	// ── Right: results ──
	results := renderResults(m, t, rightW, height)

	cols := lipgloss.JoinHorizontal(lipgloss.Top, form, "  ", results)
	return cols
}

func renderForm(m Model, t i18n.T, width int) string {
	var rows []string

	rows = append(rows, styleSectionTitle.Render("// SEARCH PARAMETERS"))

	rows = append(rows, renderInputField(m, t.LabelDirectory, fieldDir, width))
	rows = append(rows, renderInputField(m, t.LabelFilename, fieldName, width))
	rows = append(rows, renderSelectorField(m, t.LabelFileType, fieldType, width))
	rows = append(rows, renderToggleField(m, t.LabelCaseSensitive, fieldCaseSensitive, width))

	rows = append(rows, "")
	rows = append(rows, styleSectionTitle.Render("// SIZE"))
	rows = append(rows, renderSelectorField(m, "", fieldSizeCmp, width))
	rows = append(rows, renderInputField(m, t.LabelSize, fieldSizeVal, width))
	rows = append(rows, renderSelectorField(m, "", fieldSizeUnit, width))

	rows = append(rows, "")
	rows = append(rows, styleSectionTitle.Render("// TIME"))
	rows = append(rows, renderSelectorField(m, t.LabelModified, fieldTimeField, width))
	rows = append(rows, renderSelectorField(m, "", fieldTimePreset, width))
	rows = append(rows, renderInputField(m, "  "+t.TimeCustom, fieldTimeCustom, width))

	rows = append(rows, "")
	rows = append(rows, styleSectionTitle.Render("// DEPTH & PERMISSIONS"))
	rows = append(rows, renderInputField(m, t.LabelDepth, fieldDepth, width))
	rows = append(rows, renderInputField(m, t.LabelPermissions, fieldPerms, width))
	rows = append(rows, renderInputField(m, t.LabelOwner, fieldOwner, width))
	rows = append(rows, renderInputField(m, t.LabelGroup, fieldGroup, width))

	// Preview command line at bottom of form
	rows = append(rows, "")
	rows = append(rows, styleSectionTitle.Render("// COMMAND"))
	cmd := finder.BuildString(m.buildOpts())
	rows = append(rows, stylePreviewBox.Width(width-2).Render(cmd))

	// Run button
	rows = append(rows, "")
	runLabel := t.ActionRun
	if m.searching {
		runLabel = t.ResultsRunning
	}
	rows = append(rows, styleBtnPrimary.Render("▶ "+runLabel)+"  "+
		styleBtnSecondary.Render(t.ActionCopy))

	content := strings.Join(rows, "\n")
	panel := stylePanelNormal.Width(width).Render(content)
	return panel
}

func renderResults(m Model, t i18n.T, width, height int) string {
	title := styleResultsTitle.Render(t.ResultsTitle)

	var body string
	switch {
	case m.searching:
		body = styleResultsRunning.Render(t.ResultsRunning)
	case m.searchErr != "":
		body = styleResultError.Render(t.ResultsError + ": " + m.searchErr)
	case len(m.results) == 0:
		body = styleResultsEmpty.Render(t.ResultsEmpty)
	default:
		count := styleResultCount.Render(fmt.Sprintf("%d %s", len(m.results), t.ResultsCount))
		lines := make([]string, len(m.results))
		for i, r := range m.results {
			lines[i] = styleResultItem.Render(r)
		}
		body = count + "\n" + strings.Join(lines, "\n")
	}

	vp := m.resultsVP
	vp.Width  = width - 4
	vp.Height = height - 4
	vp.SetContent(body)

	panel := stylePanelNormal.Width(width).Height(height).
		Render(title + "\n" + vp.View())
	return panel
}

// ── Options tab ───────────────────────────────────────────────────────────────

func renderOptionsTab(m Model, t i18n.T, height int) string {
	w := m.width - 4
	var rows []string

	rows = append(rows, styleSectionTitle.Render("// ADVANCED OPTIONS"))
	rows = append(rows, renderSelectorField(m, t.LabelSymlinks, fieldSymlinks, w))
	rows = append(rows, renderToggleField(m, t.LabelEmpty, fieldEmpty, w))
	rows = append(rows, renderToggleField(m, t.LabelExecutable, fieldExecutable, w))
	rows = append(rows, renderToggleField(m, t.LabelReadable, fieldReadable, w))
	rows = append(rows, renderToggleField(m, t.LabelWritable, fieldWritable, w))

	content := strings.Join(rows, "\n")
	return stylePanelNormal.Width(w).Height(height).Render(content)
}

// ── Preview tab ───────────────────────────────────────────────────────────────

func renderPreviewTab(m Model, t i18n.T, height int) string {
	w := m.width - 4
	cmd := finder.BuildString(m.buildOpts())

	title := stylePreviewTitle.Render(t.PreviewTitle)
	box   := stylePreviewBox.Width(w - 4).Render(cmd)

	notice := ""
	if m.copiedNotice {
		notice = "\n" + stylePreviewCopied.Render(t.PreviewCopied)
	}

	explainTitle := styleSectionTitle.Render("// BREAKDOWN")
	explain      := renderCommandBreakdown(m.buildOpts(), m.lang)

	content := title + "\n\n" + box + notice + "\n\n" + explainTitle + "\n" + explain
	return stylePanelNormal.Width(w).Height(height).Render(content)
}

func renderCommandBreakdown(opts finder.Options, lang i18n.Lang) string {
	t := i18n.Get(lang)
	args := finder.Build(opts)
	if len(args) == 0 {
		return ""
	}
	var lines []string
	for i, a := range args {
		if i == 0 {
			lines = append(lines, styleHelpKey.Render("find")+
				styleHelpDesc.Render(" — GNU file search utility"))
			continue
		}
		desc := argDescription(a, lang, t)
		lines = append(lines, "  "+styleHelpKey.Render(a)+" "+styleHelpDesc.Render(desc))
	}
	return strings.Join(lines, "\n")
}

func argDescription(arg string, _ i18n.Lang, t i18n.T) string {
	_ = t
	switch arg {
	case "-L":
		return "follow symbolic links"
	case "-H":
		return "follow symlinks on CLI args only"
	case "-maxdepth":
		return "limit directory depth"
	case "-type":
		return "filter by file type"
	case "-name":
		return "match filename (case sensitive)"
	case "-iname":
		return "match filename (case insensitive)"
	case "-size":
		return "filter by file size"
	case "-mtime":
		return "filter by modification time"
	case "-atime":
		return "filter by access time"
	case "-ctime":
		return "filter by inode change time"
	case "-user":
		return "filter by owner"
	case "-group":
		return "filter by group"
	case "-perm":
		return "filter by permissions"
	case "-empty":
		return "match empty files/directories"
	case "-executable":
		return "match executable files"
	case "-readable":
		return "match readable files"
	case "-writable":
		return "match writable files"
	}
	return ""
}

// ── Help tab ──────────────────────────────────────────────────────────────────

func renderHelpTab(m Model, t i18n.T, height int) string {
	w := m.width - 4

	title := styleHelpTitle.Render(t.HelpTitle)

	keybindings := [][]string{
		{"Ctrl+Q",      "Quit seekr"},
		{"Ctrl+L",      "Toggle EN/FR language"},
		{"Ctrl+C",      "Copy command to clipboard"},
		{"F5 / Enter",  "Run find command"},
		{"Tab",         "Next field"},
		{"Shift+Tab",   "Previous field"},
		{"← →",         "Cycle selector values"},
		{"Space",       "Toggle boolean option"},
		{"F1",          "Search tab"},
		{"F2",          "Options tab"},
		{"F3",          "Command preview tab"},
		{"F4",          "Help tab"},
		{"↑ ↓",         "Scroll results"},
	}

	kRows := []string{styleHelpCategory.Render("⌨  " + t.HelpKeys)}
	for _, kb := range keybindings {
		kRows = append(kRows, "  "+styleHelpKey.Render(kb[0])+" "+styleHelpDesc.Render(kb[1]))
	}

	about := []string{
		styleHelpCategory.Render("◈  " + t.HelpAbout),
		"  " + styleHelpDesc.Render("seekr is a TUI wrapper around GNU find."),
		"  " + styleHelpDesc.Render("It builds valid find commands from a guided interface,"),
		"  " + styleHelpDesc.Render("making all find options discoverable and safe to use."),
		"",
		"  " + styleHelpDesc.Render("github.com/Quirky1869/seekr"),
	}

	content := title + "\n\n" +
		strings.Join(kRows, "\n") + "\n" +
		strings.Join(about, "\n")

	return stylePanelNormal.Width(w).Height(height).Render(content)
}

// ── Footer ────────────────────────────────────────────────────────────────────

func renderFooter(m Model, t i18n.T) string {
	hints := []string{
		styleFooterKey.Render("F5") + " " + t.ActionRun,
		styleFooterKey.Render("Ctrl+C") + " " + t.ActionCopy,
		styleFooterKey.Render("Ctrl+L") + " lang",
		styleFooterKey.Render("Tab") + " next",
		styleFooterKey.Render("Ctrl+Q") + " quit",
	}
	sep  := styleFooterSep.Render("  │  ")
	line := strings.Join(hints, sep)
	return styleFooter.Width(m.width).Render("  " + line)
}

// ── Field renderers ───────────────────────────────────────────────────────────

func renderInputField(m Model, label string, id fieldID, width int) string {
	focused := m.activeField == id
	var lStyle, iStyle lipgloss.Style
	if focused {
		lStyle = styleLabelFocused
		iStyle = styleInputFocused
	} else {
		lStyle = styleLabelNormal
		iStyle = styleInputNormal
	}

	inp, ok := m.inputs[id]
	if !ok {
		return ""
	}
	inp.Width = width - 6
	lbl := lStyle.Render(label)
	field := iStyle.Render(inp.View())
	return lbl + "\n" + field
}

func renderSelectorField(m Model, label string, id fieldID, width int) string {
	focused := m.activeField == id
	sel, ok := m.selectors[id]
	if !ok {
		return ""
	}

	var lStyle lipgloss.Style
	if focused {
		lStyle = styleLabelFocused
	} else {
		lStyle = styleLabelNormal
	}

	arrowL := styleSelectorArrow.Render("◀")
	arrowR := styleSelectorArrow.Render("▶")
	var valueStyle lipgloss.Style
	if focused {
		valueStyle = styleSelectorActive
	} else {
		valueStyle = styleSelectorInactive
	}
	val := valueStyle.Render(sel.label())

	row := arrowL + val + arrowR
	if label == "" {
		return row
	}
	return lStyle.Render(label) + "\n" + row
}

func renderToggleField(m Model, label string, id fieldID, _ int) string {
	focused := m.activeField == id
	on := m.toggles[id]

	var lStyle lipgloss.Style
	if focused {
		lStyle = styleLabelFocused
	} else {
		lStyle = styleLabelNormal
	}

	var check string
	if on {
		check = styleCheckOn.Render("[✓]")
	} else {
		check = styleCheckOff.Render("[ ]")
	}

	return check + " " + lStyle.Render(label)
}
