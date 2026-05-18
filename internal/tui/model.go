package tui

import (
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Quirky1869/seekr/internal/finder"
	"github.com/Quirky1869/seekr/internal/i18n"
)

// ── Tabs ──────────────────────────────────────────────────────────────────────

type tabID int

const (
	tabSearch tabID = iota
	tabOptions
	tabPreview
	tabHelp
	tabCount
)

// ── Focus fields within the Search tab ───────────────────────────────────────

type fieldID int

const (
	fieldDir fieldID = iota
	fieldName
	fieldType
	fieldSizeCmp
	fieldSizeVal
	fieldSizeUnit
	fieldTimePreset
	fieldTimeField
	fieldTimeCustom
	fieldDepth
	fieldOwner
	fieldGroup
	fieldPerms
	fieldSymlinks
	fieldCaseSensitive
	fieldEmpty
	fieldExecutable
	fieldReadable
	fieldWritable
	fieldCount
)

// ── Internal messages ─────────────────────────────────────────────────────────

type clearCopiedMsg struct{}

// ── Model ─────────────────────────────────────────────────────────────────────

// Model is the root BubbleTea model.
type Model struct {
	// UI state
	width, height int
	activeTab     tabID
	activeField   fieldID
	lang          i18n.Lang

	// Text inputs (fields that need free text entry)
	inputs map[fieldID]textinput.Model

	// Enum selectors (fields cycled with ←→)
	selectors map[fieldID]selector

	// Boolean toggles
	toggles map[fieldID]bool

	// Results viewport
	resultsVP    viewport.Model
	previewVP    viewport.Model
	results      []string
	searching    bool
	searchErr    string
	copiedNotice bool
}

// ── selector helper ───────────────────────────────────────────────────────────

type selector struct {
	choices []string // display labels
	values  []string // actual values sent to finder
	idx     int
}

func (s *selector) next() { s.idx = (s.idx + 1) % len(s.choices) }
func (s *selector) prev() {
	s.idx = (s.idx - 1 + len(s.choices)) % len(s.choices)
}
func (s selector) value() string  { return s.values[s.idx] }
func (s selector) label() string  { return s.choices[s.idx] }

// ── Constructor ───────────────────────────────────────────────────────────────

func NewModel() Model {
	m := Model{
		lang:      i18n.EN,
		activeTab: tabSearch,
		inputs:    make(map[fieldID]textinput.Model),
		selectors: make(map[fieldID]selector),
		toggles:   make(map[fieldID]bool),
	}

	// Text inputs
	textFields := []struct {
		id          fieldID
		placeholder string
	}{
		{fieldDir, "."},
		{fieldName, "*"},
		{fieldSizeVal, "0"},
		{fieldTimeCustom, "7"},
		{fieldDepth, ""},
		{fieldOwner, ""},
		{fieldGroup, ""},
		{fieldPerms, ""},
	}
	for _, tf := range textFields {
		ti := textinput.New()
		ti.Placeholder = tf.placeholder
		ti.CharLimit = 128
		ti.Width = 30
		m.inputs[tf.id] = ti
	}
	// Focus the directory field by default
	inp := m.inputs[fieldDir]
	inp.Focus()
	m.inputs[fieldDir] = inp

	// Selectors – values must match finder constants
	t := i18n.Get(i18n.EN)
	m.selectors[fieldType] = selector{
		choices: []string{t.FileTypeAny, t.FileTypeFile, t.FileTypeDir, t.FileTypeSymlink,
			t.FileTypeSocket, t.FileTypePipe, t.FileTypeBlock, t.FileTypeChar},
		values: []string{"", "f", "d", "l", "s", "p", "b", "c"},
	}
	m.selectors[fieldSizeCmp] = selector{
		choices: []string{t.SizeExact, t.SizeGt, t.SizeLt},
		values:  []string{"", "+", "-"},
	}
	m.selectors[fieldSizeUnit] = selector{
		choices: []string{t.SizeAny, t.SizeB, t.SizeK, t.SizeM, t.SizeG},
		values:  []string{"", "c", "k", "M", "G"},
	}
	m.selectors[fieldTimePreset] = selector{
		choices: []string{t.TimeAny, t.TimeToday, t.TimeYesterday, t.TimeWeek, t.TimeMonth, t.TimeCustom},
		values:  []string{"", "today", "yesterday", "week", "month", "custom"},
	}
	m.selectors[fieldTimeField] = selector{
		choices: []string{t.TimeModified, t.TimeAccessed, t.TimeChanged},
		values:  []string{"mtime", "atime", "ctime"},
	}
	m.selectors[fieldSymlinks] = selector{
		choices: []string{t.SymlinkNever, t.SymlinkFollow, t.SymlinkCLI},
		values:  []string{"", "-L", "-H"},
	}

	// Viewports
	m.resultsVP = viewport.New(60, 20)
	m.previewVP = viewport.New(60, 10)

	// Initialize toggles
	boolFields := []fieldID{fieldCaseSensitive, fieldEmpty, fieldExecutable, fieldReadable, fieldWritable}
	for _, f := range boolFields {
		m.toggles[f] = false
	}

	return m
}

// ── Init ──────────────────────────────────────────────────────────────────────

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

// ── Update ────────────────────────────────────────────────────────────────────

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.resultsVP.Width = msg.Width/2 - 4
		m.resultsVP.Height = msg.Height - 12
		m.previewVP.Width = msg.Width - 4
		m.previewVP.Height = 6

	case clearCopiedMsg:
		m.copiedNotice = false

	case finder.LinesMsg:
		m.searching = false
		m.searchErr = ""
		m.results = []string(msg)
		m.resultsVP.SetContent(strings.Join(m.results, "\n"))

	case tea.KeyMsg:
		// ── Global shortcuts ──────────────────────────────────────────────
		switch msg.String() {
		case "ctrl+q":
			return m, tea.Quit

		case "ctrl+l":
			m.lang = i18n.Toggle(m.lang)
			m.refreshSelectorLabels()
			return m, nil

		case "ctrl+c":
			if m.activeTab == tabPreview || m.activeTab == tabSearch {
				// Copy command to clipboard (best-effort)
				cmd := finder.BuildString(m.buildOpts())
				_ = copyToClipboard(cmd)
				m.copiedNotice = true
				cmds = append(cmds, tea.Tick(2*time.Second, func(t time.Time) tea.Msg {
					return clearCopiedMsg{}
				}))
				return m, tea.Batch(cmds...)
			}

		case "tab", "shift+tab":
			if m.activeTab == tabSearch || m.activeTab == tabOptions {
				m.cycleField(msg.String() == "tab")
			}

		case "f5", "enter":
			if !m.searching {
				m.searching = true
				m.results = nil
				m.searchErr = ""
				opts := m.buildOpts()
				cmds = append(cmds, runSearch(opts))
			}

		// Tab navigation
		case "f1":
			m.activeTab = tabSearch
		case "f2":
			m.activeTab = tabOptions
		case "f3":
			m.activeTab = tabPreview
		case "f4", "?":
			m.activeTab = tabHelp
		}

		// ── Per-tab key handling ──────────────────────────────────────────
		switch m.activeTab {
		case tabSearch, tabOptions:
			// Selector navigation (←→ when focused on a selector field)
			switch msg.String() {
			case "left":
				if sel, ok := m.selectors[m.activeField]; ok {
					sel.prev()
					m.selectors[m.activeField] = sel
					return m, nil
				}
			case "right":
				if sel, ok := m.selectors[m.activeField]; ok {
					sel.next()
					m.selectors[m.activeField] = sel
					return m, nil
				}
			case " ":
				if _, ok := m.toggles[m.activeField]; ok {
					m.toggles[m.activeField] = !m.toggles[m.activeField]
					return m, nil
				}
			}
			// Route keypresses to active text input
			if inp, ok := m.inputs[m.activeField]; ok {
				var c tea.Cmd
				inp, c = inp.Update(msg)
				m.inputs[m.activeField] = inp
				cmds = append(cmds, c)
			}

		case tabPreview:
			var c tea.Cmd
			m.previewVP, c = m.previewVP.Update(msg)
			cmds = append(cmds, c)

		case tabHelp:
			// scrollable help - no special handling needed
		}

		// Results viewport always scrollable
		var c tea.Cmd
		m.resultsVP, c = m.resultsVP.Update(msg)
		cmds = append(cmds, c)
	}

	return m, tea.Batch(cmds...)
}

// ── Helpers ───────────────────────────────────────────────────────────────────

// cycleField moves focus between form fields.
func (m *Model) cycleField(forward bool) {
	// Blur current input
	if inp, ok := m.inputs[m.activeField]; ok {
		inp.Blur()
		m.inputs[m.activeField] = inp
	}

	if forward {
		m.activeField = (m.activeField + 1) % fieldCount
	} else {
		m.activeField = (m.activeField - 1 + fieldCount) % fieldCount
	}

	// Focus new input
	if inp, ok := m.inputs[m.activeField]; ok {
		inp.Focus()
		m.inputs[m.activeField] = inp
	}
}

// buildOpts reads all UI state into a finder.Options struct.
func (m *Model) buildOpts() finder.Options {
	get := func(f fieldID) string {
		if inp, ok := m.inputs[f]; ok {
			return inp.Value()
		}
		return ""
	}
	selVal := func(f fieldID) string {
		if s, ok := m.selectors[f]; ok {
			return s.value()
		}
		return ""
	}

	return finder.Options{
		Directory:     get(fieldDir),
		NamePattern:   get(fieldName),
		CaseSensitive: m.toggles[fieldCaseSensitive],
		Type:          finder.FileType(selVal(fieldType)),
		SizeValue:     get(fieldSizeVal),
		SizeCmp:       finder.SizeCmp(selVal(fieldSizeCmp)),
		SizeUnit:      finder.SizeUnit(selVal(fieldSizeUnit)),
		TimePreset:    selVal(fieldTimePreset),
		TimeField:     finder.TimeField(selVal(fieldTimeField)),
		TimeValue:     get(fieldTimeCustom),
		MaxDepth:      get(fieldDepth),
		Owner:         get(fieldOwner),
		Group:         get(fieldGroup),
		Permissions:   get(fieldPerms),
		Symlinks:      finder.SymlinkMode(selVal(fieldSymlinks)),
		EmptyOnly:     m.toggles[fieldEmpty],
		Executable:    m.toggles[fieldExecutable],
		Readable:      m.toggles[fieldReadable],
		Writable:      m.toggles[fieldWritable],
	}
}

// refreshSelectorLabels updates selector display labels after a lang switch.
func (m *Model) refreshSelectorLabels() {
	t := i18n.Get(m.lang)
	typeS := m.selectors[fieldType]
	typeS.choices = []string{t.FileTypeAny, t.FileTypeFile, t.FileTypeDir, t.FileTypeSymlink,
		t.FileTypeSocket, t.FileTypePipe, t.FileTypeBlock, t.FileTypeChar}
	m.selectors[fieldType] = typeS

	cmpS := m.selectors[fieldSizeCmp]
	cmpS.choices = []string{t.SizeExact, t.SizeGt, t.SizeLt}
	m.selectors[fieldSizeCmp] = cmpS

	unitS := m.selectors[fieldSizeUnit]
	unitS.choices = []string{t.SizeAny, t.SizeB, t.SizeK, t.SizeM, t.SizeG}
	m.selectors[fieldSizeUnit] = unitS

	tpS := m.selectors[fieldTimePreset]
	tpS.choices = []string{t.TimeAny, t.TimeToday, t.TimeYesterday, t.TimeWeek, t.TimeMonth, t.TimeCustom}
	m.selectors[fieldTimePreset] = tpS

	tfS := m.selectors[fieldTimeField]
	tfS.choices = []string{t.TimeModified, t.TimeAccessed, t.TimeChanged}
	m.selectors[fieldTimeField] = tfS

	slS := m.selectors[fieldSymlinks]
	slS.choices = []string{t.SymlinkNever, t.SymlinkFollow, t.SymlinkCLI}
	m.selectors[fieldSymlinks] = slS
}

// ── Search command ────────────────────────────────────────────────────────────

// runSearch wraps finder.Run into a BubbleTea Cmd.
func runSearch(opts finder.Options) tea.Cmd {
	return func() tea.Msg {
		return finder.RunSync(opts)
	}
}

// copyToClipboard is a best-effort clipboard write.
func copyToClipboard(s string) error {
	// Uses xclip / xsel / pbcopy / clip.exe depending on OS.
	// We delegate to the platform helper; if unavailable we silently fail.
	return platformCopy(s)
}

// ── View ──────────────────────────────────────────────────────────────────────

func (m Model) View() string {
	if m.width == 0 {
		return "loading…"
	}
	return renderRoot(m)
}

// LinesMsg re-exported so runner.go can use the same type across packages.
// (defined in runner.go as finder.LinesMsg — we just reference it here)
