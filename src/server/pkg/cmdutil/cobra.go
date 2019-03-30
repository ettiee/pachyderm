package cmdutil

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/pachyderm/pachyderm/src/client/pfs"
	"github.com/spf13/cobra"
)

// RunFixedArgs wraps a function in a function
// that checks its exact argument count.
func RunFixedArgs(numArgs int, run func([]string) error) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		if len(args) != numArgs {
			fmt.Printf("expected %d arguments, got %d\n\n", numArgs, len(args))
			cmd.Usage()
		} else {
			if err := run(args); err != nil {
				ErrorAndExit("%v", err)
			}
		}
	}
}

// RunBoundedArgs wraps a function in a function
// that checks its argument count is within a range.
func RunBoundedArgs(min int, max int, run func([]string) error) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		if len(args) < min || len(args) > max {
			fmt.Printf("expected %d to %d arguments, got %d\n\n", min, max, len(args))
			cmd.Usage()
		} else {
			if err := run(args); err != nil {
				ErrorAndExit("%v", err)
			}
		}
	}
}

// Run makes a new cobra run function that wraps the given function.
func Run(run func(args []string) error) func(*cobra.Command, []string) {
	return func(_ *cobra.Command, args []string) {
		if err := run(args); err != nil {
			ErrorAndExit(err.Error())
		}
	}
}

// ErrorAndExit errors with the given format and args, and then exits.
func ErrorAndExit(format string, args ...interface{}) {
	if errString := strings.TrimSpace(fmt.Sprintf(format, args...)); errString != "" {
		fmt.Fprintf(os.Stderr, "%s\n", errString)
	}
	os.Exit(1)
}

// ParseCommit takes an argument of the form "repo[@branch-or-commit]" and
// returns the corresponding *pfs.Commit.
func ParseCommit(arg string) (*pfs.Commit, error) {
	parts := strings.SplitN(arg, "@", 2)
	if parts[0] == "" {
		return nil, fmt.Errorf("invalid format \"%s\": repo cannot be empty", arg)
	}
	commit := &pfs.Commit{
		Repo: &pfs.Repo{
			Name: parts[0],
		},
		ID: "",
	}
	if len(parts) == 2 {
		commit.ID = parts[1]
	}
	return commit, nil
}

// ParseCommits converts all arguments to *pfs.Commit structs using the
// semantics of ParseCommit
func ParseCommits(args []string) ([]*pfs.Commit, error) {
	var results []*pfs.Commit
	for _, arg := range args {
		commit, err := ParseCommit(arg)
		if err != nil {
			return nil, err
		}
		results = append(results, commit)
	}
	return results, nil
}

// ParseBranch takes an argument of the form "repo[@branch]" and
// returns the corresponding *pfs.Branch.  This uses ParseBranch under the hood
// because a branch name is usually interchangeable with a commit-id.
func ParseBranch(arg string) (*pfs.Branch, error) {
	commit, err := ParseCommit(arg)
	if err != nil {
		return nil, err
	}
	return &pfs.Branch{Repo: commit.Repo, Name: commit.ID}, nil
}

// ParseBranches converts all arguments to *pfs.Commit structs using the
// semantics of ParseBranch
func ParseBranches(args []string) ([]*pfs.Branch, error) {
	var results []*pfs.Branch
	for _, arg := range args {
		branch, err := ParseBranch(arg)
		if err != nil {
			return nil, err
		}
		results = append(results, branch)
	}
	return results, nil
}

// ParseFile takes an argument of the form "repo[@branch-or-commit[:path]]", and
// returns the corresponding *pfs.File.
func ParseFile(arg string) (*pfs.File, error) {
	repoAndRest := strings.SplitN(arg, "@", 2)
	if repoAndRest[0] == "" {
		return nil, fmt.Errorf("invalid format \"%s\": repo cannot be empty", arg)
	}
	file := &pfs.File{
		Commit: &pfs.Commit{
			Repo: &pfs.Repo{
				Name: repoAndRest[0],
			},
			ID: "",
		},
		Path: "",
	}
	if len(repoAndRest) > 1 {
		commitAndPath := strings.SplitN(repoAndRest[1], ":", 2)
		if commitAndPath[0] == "" {
			return nil, fmt.Errorf("invalid format \"%s\": commit cannot be empty", arg)
		}
		file.Commit.ID = commitAndPath[0]
		if len(commitAndPath) > 1 {
			file.Path = commitAndPath[1]
		}
	}
	return file, nil
}

// ParseFiles converts all arguments to *pfs.Commit structs using the
// semantics of ParseFile
func ParseFiles(args []string) ([]*pfs.File, error) {
	var results []*pfs.File
	for _, arg := range args {
		commit, err := ParseFile(arg)
		if err != nil {
			return nil, err
		}
		results = append(results, commit)
	}
	return results, nil
}

// RepeatedStringArg is an alias for []string
type RepeatedStringArg []string

func (r *RepeatedStringArg) String() string {
	result := "["
	for i, s := range *r {
		if i != 0 {
			result += ", "
		}
		result += s
	}
	return result + "]"
}

// Set adds a string to r
func (r *RepeatedStringArg) Set(s string) error {
	*r = append(*r, s)
	return nil
}

// Type returns the string representation of the type of r
func (r *RepeatedStringArg) Type() string {
	return "[]string"
}

// CreateAliases generates one or many nested command trees for each invocation
// listed in 'invocations', which should be space-delimited as on the command-
// line.  The 'Use' field of 'cmd' should not include the name of the command as
// that will be filled in based on each invocation.  These commands can later be
// merged into the final Command tree using 'MergeCommands' below.
func CreateAliases(cmd *cobra.Command, invocations []string) []*cobra.Command {
	var aliases []*cobra.Command

	// Create logical commands for each substring in each invocation
	for _, invocation := range invocations {
		var root, prev *cobra.Command
		args := strings.Split(invocation, " ")

		for i, arg := range args {
			cur := &cobra.Command{}

			// The leaf command node should include the usage from the given cmd,
			// while logical nodes just need one piece of the invocation.
			if i == len(args) - 1 {
				*cur = *cmd
				if cmd.Use == "" {
					cur.Use = arg
				} else {
					cur.Use = strings.Replace(cmd.Use, "{{alias}}", arg, -1)
				}
				cur.Example = strings.Replace(cmd.Example, "{{alias}}", fmt.Sprintf("%s %s", os.Args[0], invocation), -1)
			} else {
				cur.Use = arg
			}

			if root == nil {
				root = cur
			} else if prev != nil {
				prev.AddCommand(cur)
			}
			prev = cur
		}

		aliases = append(aliases, root)
	}

	return aliases
}

// MergeCommands merges several command aliases (generated by 'CreateAliases'
// above) into a single coherent cobra command tree (with root command 'root').
// Because 'CreateAliases' generates empty commands to preserve the right
// hierarchy, we go through a little extra effort to allow intermediate 'docs'
// commands to be preserved in the final command structure.
func MergeCommands(root *cobra.Command, children []*cobra.Command) {
	// Implement our own 'find' function because Command.Find is not reliable?
	findCommand := func(parent *cobra.Command, name string) *cobra.Command {
		for _, cmd := range parent.Commands() {
			if cmd.Name() == name {
				return cmd
			}
		}
		return nil
	}

	// Sort children by max nesting depth - this will put 'docs' commands first,
	// so they are not overwritten by logical commands added by aliases.
	var depth func(*cobra.Command) int
	depth = func(cmd *cobra.Command) int {
		maxDepth := 0
		for _, subcmd := range cmd.Commands() {
			subcmdDepth := depth(subcmd)
			if subcmdDepth > maxDepth {
				maxDepth = subcmdDepth
			}
		}
		return maxDepth + 1
	}

	sort.Slice(children, func(i, j int) bool {
		return depth(children[i]) < depth(children[j])
	})

	// Move each child command over to the main command tree recursively
	for _, cmd := range children {
		parent := findCommand(root, cmd.Name())
		if parent == nil {
			root.AddCommand(cmd)
		} else {
			MergeCommands(parent, cmd.Commands())
		}
	}
}
