package finder

import (
	"fmt"
	"strings"
)

// FileType maps UI choices to find -type values.
type FileType string

const (
	TypeAny     FileType = ""
	TypeFile    FileType = "f"
	TypeDir     FileType = "d"
	TypeSymlink FileType = "l"
	TypeSocket  FileType = "s"
	TypePipe    FileType = "p"
	TypeBlock   FileType = "b"
	TypeChar    FileType = "c"
)

// SizeCmp is the size comparator.
type SizeCmp string

const (
	SizeExact SizeCmp = ""
	SizeGt    SizeCmp = "+"
	SizeLt    SizeCmp = "-"
)

// SizeUnit maps UI label to find suffix.
type SizeUnit string

const (
	UnitNone SizeUnit = ""
	UnitByte SizeUnit = "c"
	UnitKilo SizeUnit = "k"
	UnitMega SizeUnit = "M"
	UnitGiga SizeUnit = "G"
)

// TimeField selects which timestamp to filter on.
type TimeField string

const (
	TimeModified TimeField = "mtime"
	TimeAccessed TimeField = "atime"
	TimeChanged  TimeField = "ctime"
)

// SymlinkMode mirrors find's -P/-L/-H.
type SymlinkMode string

const (
	SymlinkNever  SymlinkMode = ""   // -P (default)
	SymlinkFollow SymlinkMode = "-L"
	SymlinkCLI    SymlinkMode = "-H"
)

// Options holds all the search parameters collected from the UI.
type Options struct {
	// Starting directory
	Directory string

	// Name matching
	NamePattern   string
	CaseSensitive bool // false → -iname

	// File type
	Type FileType

	// Size
	SizeValue string  // numeric value as string, empty = ignore
	SizeCmp   SizeCmp
	SizeUnit  SizeUnit

	// Time
	TimeValue string    // days as string, empty = ignore
	TimeField TimeField
	TimePreset string   // "today", "yesterday", "week", "month", "custom"

	// Depth
	MaxDepth string // empty = unlimited

	// Permissions / ownership
	Permissions string // e.g. "755", "-u=x"
	Owner       string
	Group       string

	// Boolean flags
	EmptyOnly  bool
	Executable bool
	Readable   bool
	Writable   bool

	// Symlink behaviour
	Symlinks SymlinkMode
}

// Build constructs the full find command from Options.
// Returns the command as a slice of strings (args[0] == "find").
func Build(o Options) []string {
	args := []string{"find"}

	// Symlink flag must come before the path
	if o.Symlinks == SymlinkFollow || o.Symlinks == SymlinkCLI {
		args = append(args, string(o.Symlinks))
	}

	// Starting directory
	dir := o.Directory
	if dir == "" {
		dir = "."
	}
	args = append(args, dir)

	// Max depth
	if o.MaxDepth != "" && o.MaxDepth != "0" {
		args = append(args, "-maxdepth", o.MaxDepth)
	}

	// File type
	if o.Type != TypeAny {
		args = append(args, "-type", string(o.Type))
	}

	// Name pattern
	if o.NamePattern != "" {
		if o.CaseSensitive {
			args = append(args, "-name", o.NamePattern)
		} else {
			args = append(args, "-iname", o.NamePattern)
		}
	}

	// Size
	if o.SizeValue != "" && o.SizeUnit != UnitNone {
		sizeArg := fmt.Sprintf("%s%s%s", string(o.SizeCmp), o.SizeValue, string(o.SizeUnit))
		args = append(args, "-size", sizeArg)
	}

	// Time
	timeVal := resolveTimePreset(o.TimePreset, o.TimeValue)
	if timeVal != "" {
		field := string(o.TimeField)
		if field == "" {
			field = string(TimeModified)
		}
		args = append(args, "-"+field, timeVal)
	}

	// Owner
	if o.Owner != "" {
		args = append(args, "-user", o.Owner)
	}

	// Group
	if o.Group != "" {
		args = append(args, "-group", o.Group)
	}

	// Permissions
	if o.Permissions != "" {
		args = append(args, "-perm", o.Permissions)
	}

	// Boolean predicates
	if o.EmptyOnly {
		args = append(args, "-empty")
	}
	if o.Executable {
		args = append(args, "-executable")
	}
	if o.Readable {
		args = append(args, "-readable")
	}
	if o.Writable {
		args = append(args, "-writable")
	}

	return args
}

// BuildString returns the command as a single formatted string.
func BuildString(o Options) string {
	return strings.Join(Build(o), " ")
}

// resolveTimePreset converts a named preset to a find -mtime numeric value.
func resolveTimePreset(preset, custom string) string {
	switch preset {
	case "today":
		return "-1"
	case "yesterday":
		return "1"
	case "week":
		return "-7"
	case "month":
		return "-30"
	case "custom":
		if custom != "" {
			return "-" + custom
		}
	}
	return ""
}
