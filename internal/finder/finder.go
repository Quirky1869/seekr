package finder

import (
	"fmt"
	"os/exec"
	"strings"
)

// Options holds all find parameters
type Options struct {
	StartPath   string
	FileName    string
	FileType    string // "", "f", "d", "l"
	MaxDepth    string
	MinDepth    string
	Mtime       string
	Size        string
	Perm        string
	Owner       string
	Empty       bool
	Executable  bool
	Readable    bool
	Writable    bool
	FollowSym   bool
	NoMount     bool
	Regex       string
	Exclude     string
	DeleteMode  bool
}

// BuildCommand constructs the find command arguments from options
func BuildCommand(o Options) []string {
	args := []string{}

	// Symlink flag before path
	if o.FollowSym {
		args = append(args, "-L")
	}

	// Starting path
	path := strings.TrimSpace(o.StartPath)
	if path == "" {
		path = "."
	}
	args = append(args, path)

	// Depth options
	if o.MaxDepth != "" {
		args = append(args, "-maxdepth", o.MaxDepth)
	}
	if o.MinDepth != "" {
		args = append(args, "-mindepth", o.MinDepth)
	}

	// No mount
	if o.NoMount {
		args = append(args, "-xdev")
	}

	// Exclude via prune
	if excl := strings.TrimSpace(o.Exclude); excl != "" {
		args = append(args, "-path", excl, "-prune", "-o")
	}

	// File type
	if o.FileType != "" {
		args = append(args, "-type", o.FileType)
	}

	// Name pattern
	if name := strings.TrimSpace(o.FileName); name != "" {
		if strings.ContainsAny(name, ".*?[") {
			args = append(args, "-name", name)
		} else {
			args = append(args, "-name", name)
		}
	}

	// Regex
	if rx := strings.TrimSpace(o.Regex); rx != "" {
		args = append(args, "-regex", rx)
	}

	// Mtime
	if m := strings.TrimSpace(o.Mtime); m != "" {
		args = append(args, "-mtime", m)
	}

	// Size
	if s := strings.TrimSpace(o.Size); s != "" {
		args = append(args, "-size", s)
	}

	// Permissions
	if p := strings.TrimSpace(o.Perm); p != "" {
		args = append(args, "-perm", p)
	}

	// Owner
	if ow := strings.TrimSpace(o.Owner); ow != "" {
		args = append(args, "-user", ow)
	}

	// Boolean predicates
	if o.Empty {
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

	// Delete (dangerous — add last)
	if o.DeleteMode {
		args = append(args, "-delete")
	} else {
		args = append(args, "-print")
	}

	return args
}

// CommandString returns the full command as a shell string
func CommandString(o Options) string {
	args := BuildCommand(o)
	// Quote args that contain spaces or special chars
	quoted := make([]string, len(args)+1)
	quoted[0] = "find"
	for i, a := range args {
		if strings.ContainsAny(a, " \t*?[]{}") {
			quoted[i+1] = fmt.Sprintf("'%s'", a)
		} else {
			quoted[i+1] = a
		}
	}
	return strings.Join(quoted, " ")
}

// Run executes find and returns the lines of output
func Run(o Options) ([]string, error) {
	args := BuildCommand(o)
	cmd := exec.Command("find", args...)
	out, err := cmd.Output()
	if err != nil {
		// find exits non-zero on permission errors but may still have output
		if exitErr, ok := err.(*exec.ExitError); ok {
			lines := splitLines(string(out))
			if len(lines) > 0 {
				_ = exitErr
				return lines, nil
			}
		}
		return nil, err
	}
	return splitLines(string(out)), nil
}

func splitLines(s string) []string {
	lines := strings.Split(strings.TrimRight(s, "\n"), "\n")
	result := []string{}
	for _, l := range lines {
		if l != "" {
			result = append(result, l)
		}
	}
	return result
}
