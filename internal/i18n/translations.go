package i18n

// Lang represents a language code.
type Lang string

const (
	EN Lang = "en"
	FR Lang = "fr"
)

// T holds all translatable strings.
type T struct {
	// Header / branding
	AppSubtitle string

	// Navigation tabs
	TabSearch  string
	TabOptions string
	TabPreview string
	TabHelp    string

	// Search panel
	LabelDirectory   string
	LabelFilename    string
	LabelFileType    string
	LabelSize        string
	LabelModified    string
	LabelPermissions string
	LabelDepth       string
	LabelOwner       string
	LabelGroup       string
	LabelEmpty       string
	LabelExecutable  string
	LabelReadable    string
	LabelWritable    string
	LabelSymlinks    string
	LabelCaseSensitive string

	// File type options
	FileTypeAny       string
	FileTypeFile      string
	FileTypeDir       string
	FileTypeSymlink   string
	FileTypeSocket    string
	FileTypePipe      string
	FileTypeBlock     string
	FileTypeChar      string

	// Size units
	SizeAny string
	SizeB   string
	SizeK   string
	SizeM   string
	SizeG   string
	// Size comparators
	SizeExact  string
	SizeGt     string
	SizeLt     string

	// Time options
	TimeAny       string
	TimeToday     string
	TimeYesterday string
	TimeWeek      string
	TimeMonth     string
	TimeCustom    string
	TimeModified  string
	TimeAccessed  string
	TimeChanged   string

	// Symlink options
	SymlinkNever  string
	SymlinkFollow string
	SymlinkCLI    string

	// Depth
	DepthUnlimited string

	// Results panel
	ResultsTitle       string
	ResultsEmpty       string
	ResultsRunning     string
	ResultsCount       string
	ResultsError       string

	// Command preview
	PreviewTitle   string
	PreviewCopied  string

	// Actions
	ActionRun    string
	ActionCopy   string
	ActionClear  string
	ActionQuit   string
	ActionExport string

	// Help
	HelpTitle      string
	HelpKeys       string
	HelpAbout      string

	// Footer hints
	HintRun    string
	HintCopy   string
	HintQuit   string
	HintTab    string
	HintLang   string
	HintScroll string
	HintSelect string

	// Errors
	ErrNoDir     string
	ErrFindExec  string
	ErrCopyClip  string
}

var translations = map[Lang]T{
	EN: {
		AppSubtitle: "[ FILE SEARCH INTERFACE // find wrapper ]",

		TabSearch:  " SEARCH ",
		TabOptions: " OPTIONS ",
		TabPreview: " COMMAND ",
		TabHelp:    " HELP ",

		LabelDirectory:   "Directory",
		LabelFilename:    "Filename pattern",
		LabelFileType:    "File type",
		LabelSize:        "Size",
		LabelModified:    "Last modified",
		LabelPermissions: "Permissions",
		LabelDepth:       "Max depth",
		LabelOwner:       "Owner",
		LabelGroup:       "Group",
		LabelEmpty:       "Empty files only",
		LabelExecutable:  "Executable only",
		LabelReadable:    "Readable only",
		LabelWritable:    "Writable only",
		LabelSymlinks:    "Symlink handling",
		LabelCaseSensitive: "Case sensitive",

		FileTypeAny:     "any",
		FileTypeFile:    "file",
		FileTypeDir:     "directory",
		FileTypeSymlink: "symlink",
		FileTypeSocket:  "socket",
		FileTypePipe:    "pipe",
		FileTypeBlock:   "block device",
		FileTypeChar:    "char device",

		SizeAny:   "any",
		SizeB:     "bytes",
		SizeK:     "kilobytes",
		SizeM:     "megabytes",
		SizeG:     "gigabytes",
		SizeExact: "exactly",
		SizeGt:    "greater than",
		SizeLt:    "less than",

		TimeAny:       "any time",
		TimeToday:     "today",
		TimeYesterday: "yesterday",
		TimeWeek:      "last 7 days",
		TimeMonth:     "last 30 days",
		TimeCustom:    "custom (days)",
		TimeModified:  "modified",
		TimeAccessed:  "accessed",
		TimeChanged:   "changed",

		SymlinkNever:  "never follow",
		SymlinkFollow: "always follow (-L)",
		SymlinkCLI:    "CLI args only (-H)",

		DepthUnlimited: "unlimited",

		ResultsTitle:   "// RESULTS",
		ResultsEmpty:   "[ no results — adjust your filters ]",
		ResultsRunning: "[ scanning... ]",
		ResultsCount:   "results",
		ResultsError:   "ERROR",

		PreviewTitle:  "// GENERATED COMMAND",
		PreviewCopied: "[ copied to clipboard ]",

		ActionRun:    "RUN",
		ActionCopy:   "COPY",
		ActionClear:  "CLEAR",
		ActionQuit:   "QUIT",
		ActionExport: "EXPORT",

		HelpTitle: "// SEEKR — HELP",
		HelpKeys:  "KEYBINDINGS",
		HelpAbout: "ABOUT",

		HintRun:    "F5/Enter: run",
		HintCopy:   "Ctrl+C: copy cmd",
		HintQuit:   "Ctrl+Q: quit",
		HintTab:    "Tab: next field",
		HintLang:   "Ctrl+L: lang",
		HintScroll: "↑↓: scroll",
		HintSelect: "←→: select",

		ErrNoDir:    "Directory does not exist",
		ErrFindExec: "Failed to execute find",
		ErrCopyClip: "Failed to copy to clipboard",
	},

	FR: {
		AppSubtitle: "[ INTERFACE DE RECHERCHE // find wrapper ]",

		TabSearch:  " RECHERCHE ",
		TabOptions: " OPTIONS ",
		TabPreview: " COMMANDE ",
		TabHelp:    " AIDE ",

		LabelDirectory:   "Répertoire",
		LabelFilename:    "Motif de nom",
		LabelFileType:    "Type de fichier",
		LabelSize:        "Taille",
		LabelModified:    "Modifié le",
		LabelPermissions: "Permissions",
		LabelDepth:       "Profondeur max",
		LabelOwner:       "Propriétaire",
		LabelGroup:       "Groupe",
		LabelEmpty:       "Fichiers vides uniquement",
		LabelExecutable:  "Exécutables uniquement",
		LabelReadable:    "Lisibles uniquement",
		LabelWritable:    "Modifiables uniquement",
		LabelSymlinks:    "Liens symboliques",
		LabelCaseSensitive: "Casse sensible",

		FileTypeAny:     "tous",
		FileTypeFile:    "fichier",
		FileTypeDir:     "répertoire",
		FileTypeSymlink: "lien symbolique",
		FileTypeSocket:  "socket",
		FileTypePipe:    "pipe",
		FileTypeBlock:   "périph. bloc",
		FileTypeChar:    "périph. car.",

		SizeAny:   "toute taille",
		SizeB:     "octets",
		SizeK:     "kilooctets",
		SizeM:     "mégaoctets",
		SizeG:     "gigaoctets",
		SizeExact: "exactement",
		SizeGt:    "supérieur à",
		SizeLt:    "inférieur à",

		TimeAny:       "à tout moment",
		TimeToday:     "aujourd'hui",
		TimeYesterday: "hier",
		TimeWeek:      "7 derniers jours",
		TimeMonth:     "30 derniers jours",
		TimeCustom:    "personnalisé (jours)",
		TimeModified:  "modifié",
		TimeAccessed:  "accédé",
		TimeChanged:   "changé",

		SymlinkNever:  "ne pas suivre",
		SymlinkFollow: "toujours suivre (-L)",
		SymlinkCLI:    "args CLI seulement (-H)",

		DepthUnlimited: "illimité",

		ResultsTitle:   "// RÉSULTATS",
		ResultsEmpty:   "[ aucun résultat — affinez vos filtres ]",
		ResultsRunning: "[ analyse en cours... ]",
		ResultsCount:   "résultats",
		ResultsError:   "ERREUR",

		PreviewTitle:  "// COMMANDE GÉNÉRÉE",
		PreviewCopied: "[ copié dans le presse-papiers ]",

		ActionRun:    "LANCER",
		ActionCopy:   "COPIER",
		ActionClear:  "EFFACER",
		ActionQuit:   "QUITTER",
		ActionExport: "EXPORTER",

		HelpTitle: "// SEEKR — AIDE",
		HelpKeys:  "RACCOURCIS",
		HelpAbout: "À PROPOS",

		HintRun:    "F5/Entrée: lancer",
		HintCopy:   "Ctrl+C: copier cmd",
		HintQuit:   "Ctrl+Q: quitter",
		HintTab:    "Tab: champ suivant",
		HintLang:   "Ctrl+L: langue",
		HintScroll: "↑↓: défiler",
		HintSelect: "←→: sélectionner",

		ErrNoDir:    "Le répertoire n'existe pas",
		ErrFindExec: "Impossible d'exécuter find",
		ErrCopyClip: "Impossible de copier",
	},
}

// Get returns the translation set for the given language.
// Falls back to English if unknown.
func Get(l Lang) T {
	if t, ok := translations[l]; ok {
		return t
	}
	return translations[EN]
}

// Toggle switches between EN and FR.
func Toggle(l Lang) Lang {
	if l == EN {
		return FR
	}
	return EN
}

// Flag returns a small flag emoji for the language.
func Flag(l Lang) string {
	switch l {
	case FR:
		return "🇫🇷"
	default:
		return "🇬🇧"
	}
}
