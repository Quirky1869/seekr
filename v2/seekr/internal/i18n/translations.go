package i18n

// Lang represents a supported language
type Lang string

const (
	EN Lang = "en"
	FR Lang = "fr"
)

// T holds all UI strings
type T struct {
	// Header
	AppTitle    string
	AppSubtitle string

	// Tabs
	TabBasic    string
	TabFilters  string
	TabAdvanced string
	TabResults  string

	// Basic fields
	LabelStartPath  string
	LabelFileName   string
	LabelFileType   string
	LabelMaxDepth   string
	PlaceholderPath string
	PlaceholderName string

	// File types
	TypeAny       string
	TypeFile      string
	TypeDirectory string
	TypeSymlink   string

	// Filters
	LabelMtime      string
	LabelSize       string
	LabelPerm       string
	LabelOwner      string
	LabelEmpty      string
	LabelExecutable string
	LabelReadable   string
	LabelWritable   string
	SizeUnit        string

	// Advanced
	LabelMinDepth   string
	LabelFollowSym  string
	LabelNoMount    string
	LabelRegex      string
	LabelExclude    string
	LabelDeleteMode string

	// Results
	LabelResults     string
	LabelCommand     string
	LabelNoResults   string
	LabelRunning     string
	LabelCopied      string
	LabelResultCount string

	// Buttons / actions
	BtnRun    string
	BtnCopy   string
	BtnClear  string
	BtnQuit   string

	// Errors
	ErrNoPath  string
	ErrBadPath string

	// Help bar
	HelpRun      string
	HelpCopy     string
	HelpTab      string
	HelpLang     string
	HelpQuit     string
	HelpNav      string
	HelpConfirm  string
}

var translations = map[Lang]T{
	EN: {
		AppTitle:    "SEEKR",
		AppSubtitle: "// find interface v1.0 //",

		TabBasic:    "[1] BASIC",
		TabFilters:  "[2] FILTERS",
		TabAdvanced: "[3] ADVANCED",
		TabResults:  "[4] RESULTS",

		LabelStartPath:  "START PATH",
		LabelFileName:   "FILE NAME",
		LabelFileType:   "TYPE",
		LabelMaxDepth:   "MAX DEPTH",
		PlaceholderPath: "/home/user  (leave empty for .)",
		PlaceholderName: "*.log  or  myfile.txt",

		TypeAny:       "ANY",
		TypeFile:      "FILE (f)",
		TypeDirectory: "DIR  (d)",
		TypeSymlink:   "LINK (l)",

		LabelMtime:      "MODIFIED (days)",
		LabelSize:       "SIZE",
		LabelPerm:       "PERMISSIONS",
		LabelOwner:      "OWNER",
		LabelEmpty:      "EMPTY FILES ONLY",
		LabelExecutable: "EXECUTABLE",
		LabelReadable:   "READABLE",
		LabelWritable:   "WRITABLE",
		SizeUnit:        "unit: +5M  -1k  100c",

		LabelMinDepth:   "MIN DEPTH",
		LabelFollowSym:  "FOLLOW SYMLINKS (-L)",
		LabelNoMount:    "NO MOUNT (-xdev)",
		LabelRegex:      "REGEX PATTERN",
		LabelExclude:    "EXCLUDE PATH",
		LabelDeleteMode: "DELETE MATCHED (!)",

		LabelResults:     "RESULTS",
		LabelCommand:     "GENERATED COMMAND",
		LabelNoResults:   "[ no results — run a search ]",
		LabelRunning:     "[ scanning... ]",
		LabelCopied:      "[ command copied to clipboard ]",
		LabelResultCount: "results",

		BtnRun:   "[ F5  RUN ]",
		BtnCopy:  "[ F6  COPY CMD ]",
		BtnClear: "[ F7  CLEAR ]",
		BtnQuit:  "[ Q  QUIT ]",

		ErrNoPath:  "⚠  path is empty — using current directory",
		ErrBadPath: "⚠  path does not exist",

		HelpRun:     "F5 run",
		HelpCopy:    "F6 copy",
		HelpTab:     "Tab/1-4 navigate",
		HelpLang:    "Ctrl+L lang",
		HelpQuit:    "q quit",
		HelpNav:     "↑↓ scroll",
		HelpConfirm: "Enter confirm",
	},
	FR: {
		AppTitle:    "SEEKR",
		AppSubtitle: "// interface find v1.0 //",

		TabBasic:    "[1] BASE",
		TabFilters:  "[2] FILTRES",
		TabAdvanced: "[3] AVANCÉ",
		TabResults:  "[4] RÉSULTATS",

		LabelStartPath:  "CHEMIN DE DÉPART",
		LabelFileName:   "NOM DU FICHIER",
		LabelFileType:   "TYPE",
		LabelMaxDepth:   "PROFONDEUR MAX",
		PlaceholderPath: "/home/user  (vide = .)",
		PlaceholderName: "*.log  ou  monfichier.txt",

		TypeAny:       "TOUT",
		TypeFile:      "FICHIER (f)",
		TypeDirectory: "DOSSIER (d)",
		TypeSymlink:   "LIEN  (l)",

		LabelMtime:      "MODIFIÉ (jours)",
		LabelSize:       "TAILLE",
		LabelPerm:       "PERMISSIONS",
		LabelOwner:      "PROPRIÉTAIRE",
		LabelEmpty:      "FICHIERS VIDES SEULEMENT",
		LabelExecutable: "EXÉCUTABLE",
		LabelReadable:   "LISIBLE",
		LabelWritable:   "MODIFIABLE",
		SizeUnit:        "unité : +5M  -1k  100c",

		LabelMinDepth:   "PROFONDEUR MIN",
		LabelFollowSym:  "SUIVRE LIENS (-L)",
		LabelNoMount:    "SANS MONTAGE (-xdev)",
		LabelRegex:      "MOTIF REGEX",
		LabelExclude:    "EXCLURE CHEMIN",
		LabelDeleteMode: "SUPPRIMER RÉSULTATS (!)",

		LabelResults:     "RÉSULTATS",
		LabelCommand:     "COMMANDE GÉNÉRÉE",
		LabelNoResults:   "[ aucun résultat — lancez une recherche ]",
		LabelRunning:     "[ analyse en cours... ]",
		LabelCopied:      "[ commande copiée dans le presse-papier ]",
		LabelResultCount: "résultats",

		BtnRun:   "[ F5  LANCER ]",
		BtnCopy:  "[ F6  COPIER CMD ]",
		BtnClear: "[ F7  EFFACER ]",
		BtnQuit:  "[ Q  QUITTER ]",

		ErrNoPath:  "⚠  chemin vide — utilisation du répertoire courant",
		ErrBadPath: "⚠  chemin inexistant",

		HelpRun:     "F5 lancer",
		HelpCopy:    "F6 copier",
		HelpTab:     "Tab/1-4 naviguer",
		HelpLang:    "Ctrl+L langue",
		HelpQuit:    "q quitter",
		HelpNav:     "↑↓ défiler",
		HelpConfirm: "Entrée valider",
	},
}

// Get returns the translation struct for the given language
func Get(l Lang) T {
	if t, ok := translations[l]; ok {
		return t
	}
	return translations[EN]
}

// Toggle switches between EN and FR
func Toggle(current Lang) Lang {
	if current == EN {
		return FR
	}
	return EN
}

// Flag returns a small flag emoji for the language
func Flag(l Lang) string {
	if l == FR {
		return "🇫🇷 FR"
	}
	return "🇬🇧 EN"
}
