package tui

import (
	"os/exec"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Quirky1869/seekr/internal/finder"
	"github.com/Quirky1869/seekr/internal/i18n"
)

// ─── Tab / field constants ───────────────────────────────────────────────────

const (
	TabBasic    = 0
	TabFilters  = 1
	TabAdvanced = 2
	TabResults  = 3
)

// field IDs (unique per input)
const (
	fPath     = 0
	fName     = 1
	fType     = 2 // cycles, not a text input
	fMaxDepth = 3

	fMtime      = 4
	fSize       = 5
	fPerm       = 6
	fOwner      = 7
	fEmpty      = 8
	fExecutable = 9
	fReadable   = 10
	fWritable   = 11

	fMinDepth   = 12
	fFollowSym  = 13
	fNoMount    = 14
	fRegex      = 15
	fExclude    = 16
	fDeleteMode = 17
)

var tabFields = map[int][]int{
	TabBasic:    {fPath, fName, fType, fMaxDepth},
	TabFilters:  {fMtime, fSize, fPerm, fOwner, fEmpty, fExecutable, fReadable, fWritable},
	TabAdvanced: {fMinDepth, fFollowSym, fNoMount, fRegex, fExclude, fDeleteMode},
	TabResults:  {},
}

// fields that are toggles (bool) or selectors, not text inputs
var nonTextFields = map[int]bool{
	fType: true, fEmpty: true, fExecutable: true, fReadable: true,
	fWritable: true, fFollowSym: true, fNoMount: true, fDeleteMode: true,
}

var fileTypeValues = []string{"", "f", "d", "l"}

// ─── Messages ────────────────────────────────────────────────────────────────

type searchResultMsg struct {
	results []string
	err     error
}

type copiedMsg struct{}

// ─── Model ───────────────────────────────────────────────────────────────────

type Model struct {
	lang   i18n.Lang
	tr     i18n.T
	width  int
	height int

	tab     int
	focused int

	inputs  map[int]textinput.Model
	toggles map[int]bool

	fileTypeIdx int

	results   []string
	resultVP  viewport.Model
	running   bool
	errMsg    string
	statusMsg string
}

// ─── Constructor ─────────────────────────────────────────────────────────────

func InitialModel() Model {
	lang := i18n.EN
	tr := i18n.Get(lang)

	m := Model{
		lang:    lang,
		tr:      tr,
		tab:     TabBasic,
		focused: fPath,
		toggles: make(map[int]bool),
		inputs:  make(map[int]textinput.Model),
	}

	type def struct {
		id int
		ph string
		w  int
	}
	defs := []def{
		{fPath, tr.PlaceholderPath, 40},
		{fName, tr.PlaceholderName, 40},
		{fMaxDepth, "∞", 6},
		{fMtime, "-1  +7  0", 40},
		{fSize, "+5M  -1k  100c", 40},
		{fPerm, "-664  /222", 40},
		{fOwner, "root  or  1000", 40},
		{fMinDepth, "0", 6},
		{fRegex, `.*\.go$`, 40},
		{fExclude, "./.git", 40},
	}
	for _, d := range defs {
		ti := textinput.New()
		ti.Placeholder = d.ph
		ti.CharLimit = 256
		ti.Width = d.w
		m.inputs[d.id] = ti
	}

	m = m.focusField(fPath)
	m.resultVP = viewport.New(80, 10)
	return m
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

// ─── Update ──────────────────────────────────────────────────────────────────

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		inputW := clamp(m.width/2-26, 20, 80)
		for id, ti := range m.inputs {
			if id == fMaxDepth || id == fMinDepth {
				ti.Width = 6
			} else {
				ti.Width = inputW
			}
			m.inputs[id] = ti
		}
		vpH := clamp(m.height-13, 3, 200)
		m.resultVP.Width = clamp(m.width-4, 20, 300)
		m.resultVP.Height = vpH
		return m, nil

	case searchResultMsg:
		m.running = false
		if msg.err != nil {
			m.errMsg = msg.err.Error()
			m.results = nil
		} else {
			m.errMsg = ""
			m.results = msg.results
		}
		m.resultVP.SetContent(strings.Join(m.results, "\n"))
		m.tab = TabResults
		return m, nil

	case copiedMsg:
		return m, nil

	case tea.KeyMsg:
		k := msg.String()

		// Global shortcuts — always active
		switch k {
		case "ctrl+c":
			return m, tea.Quit
		case "ctrl+l":
			m.lang = i18n.Toggle(m.lang)
			m.tr = i18n.Get(m.lang)
			m.refreshPlaceholders()
			return m, nil
		case "f5":
			m.running = true
			m.statusMsg = ""
			m.errMsg = ""
			opts := m.buildOptions()
			return m, func() tea.Msg {
				res, err := finder.Run(opts)
				return searchResultMsg{results: res, err: err}
			}
		case "f6":
			m.statusMsg = m.tr.LabelCopied
			s := m.buildCmdString()
			return m, func() tea.Msg {
				clipboardWrite(s)
				return copiedMsg{}
			}
		case "f7":
			m.results = nil
			m.errMsg = ""
			m.statusMsg = ""
			m.resultVP.SetContent("")
			return m, nil
		case "f1", "f2", "f3", "f4":
			t := int(k[1]-'1')
			return m.switchTab(t), nil
		case "tab":
			return m.switchTab((m.tab + 1) % 4), nil
		}

		// Quit only when not typing
		if k == "q" && (m.tab == TabResults || nonTextFields[m.focused]) {
			return m, tea.Quit
		}

		// Scroll results
		if m.tab == TabResults {
			switch k {
			case "up", "k":
				m.resultVP.LineUp(3)
			case "down", "j":
				m.resultVP.LineDown(3)
			}
			return m, nil
		}

		// Navigation within tab
		switch k {
		case "up", "shift+tab":
			m = m.moveFocus(-1)
			return m, nil
		case "down":
			m = m.moveFocus(1)
			return m, nil
		case "enter":
			if m.focused == fType {
				m.fileTypeIdx = (m.fileTypeIdx + 1) % len(fileTypeValues)
				return m, nil
			}
			if nonTextFields[m.focused] {
				m.toggles[m.focused] = !m.toggles[m.focused]
				return m, nil
			}
			case "left", "right":
			if m.focused == fType {
				if k == "right" {
					m.fileTypeIdx = (m.fileTypeIdx + 1) % len(fileTypeValues)
				} else {
					m.fileTypeIdx = (m.fileTypeIdx - 1 + len(fileTypeValues)) % len(fileTypeValues)
				}
				return m, nil
			}
		}

		// Text input forwarding
		if ti, ok := m.inputs[m.focused]; ok {
			var cmd tea.Cmd
			ti, cmd = ti.Update(msg)
			m.inputs[m.focused] = ti
			cmds = append(cmds, cmd)
		}
	}

	var vpCmd tea.Cmd
	m.resultVP, vpCmd = m.resultVP.Update(msg)
	cmds = append(cmds, vpCmd)

	return m, tea.Batch(cmds...)
}

// ─── Focus / tab helpers ─────────────────────────────────────────────────────

func (m Model) switchTab(t int) Model {
	m = m.blurAll()
	m.tab = t
	fields := tabFields[t]
	if len(fields) > 0 {
		m.focused = fields[0]
		m = m.focusField(m.focused)
	}
	return m
}

func (m Model) moveFocus(dir int) Model {
	fields := tabFields[m.tab]
	if len(fields) == 0 {
		return m
	}
	idx := 0
	for i, f := range fields {
		if f == m.focused {
			idx = i
			break
		}
	}
	idx = (idx + dir + len(fields)) % len(fields)
	m = m.blurAll()
	m.focused = fields[idx]
	m = m.focusField(m.focused)
	return m
}

func (m Model) blurAll() Model {
	for id, ti := range m.inputs {
		ti.Blur()
		m.inputs[id] = ti
	}
	return m
}

func (m Model) focusField(id int) Model {
	if ti, ok := m.inputs[id]; ok {
		ti.Focus()
		m.inputs[id] = ti
	}
	return m
}

// ─── Build / run ─────────────────────────────────────────────────────────────

func (m Model) val(id int) string {
	if ti, ok := m.inputs[id]; ok {
		return strings.TrimSpace(ti.Value())
	}
	return ""
}

func (m Model) buildOptions() finder.Options {
	return finder.Options{
		StartPath:  m.val(fPath),
		FileName:   m.val(fName),
		FileType:   fileTypeValues[m.fileTypeIdx],
		MaxDepth:   m.val(fMaxDepth),
		MinDepth:   m.val(fMinDepth),
		Mtime:      m.val(fMtime),
		Size:       m.val(fSize),
		Perm:       m.val(fPerm),
		Owner:      m.val(fOwner),
		Empty:      m.toggles[fEmpty],
		Executable: m.toggles[fExecutable],
		Readable:   m.toggles[fReadable],
		Writable:   m.toggles[fWritable],
		FollowSym:  m.toggles[fFollowSym],
		NoMount:    m.toggles[fNoMount],
		Regex:      m.val(fRegex),
		Exclude:    m.val(fExclude),
		DeleteMode: m.toggles[fDeleteMode],
	}
}

func (m Model) buildCmdString() string {
	return finder.CommandString(m.buildOptions())
}

// ─── Helpers ─────────────────────────────────────────────────────────────────

func (m *Model) refreshPlaceholders() {
	set := func(id int, ph string) {
		if ti, ok := m.inputs[id]; ok {
			ti.Placeholder = ph
			m.inputs[id] = ti
		}
	}
	set(fPath, m.tr.PlaceholderPath)
	set(fName, m.tr.PlaceholderName)
}

func clipboardWrite(s string) {
	for _, c := range [][]string{
		{"xclip", "-selection", "clipboard"},
		{"xsel", "--clipboard", "--input"},
		{"wl-copy"},
		{"pbcopy"},
	} {
		cmd := exec.Command(c[0], c[1:]...)
		cmd.Stdin = strings.NewReader(s)
		if err := cmd.Run(); err == nil {
			return
		}
	}
}

func clamp(v, lo, hi int) int {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}
