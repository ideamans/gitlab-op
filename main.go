package main

import (
	"fmt"
	"os"

	"github.com/ideamans/go-llm-cli-kit/llmcmd"
	"github.com/spf13/cobra"

	"github.com/ideamans/gitlab-op/internal/llmdocs"
)

//go:generate go run . gen-llmdocs

// PluginVersion is the version the distributed Claude Code plugin claims.
// The release workflow refuses a tag that disagrees with it, and
// TestPluginSkills asserts plugin.json carries the same value.
const PluginVersion = "1.1.0"

var version = PluginVersion

// llmConfig wires the embedded reference into the llm subcommand and the
// deprecated --llm flag.
func llmConfig() llmcmd.Config { return llmcmd.Config{Docs: llmdocs.Docs()} }

// newRootCmd assembles the command tree without executing it, so gen-llmdocs
// can walk it to produce the command catalog chapter.
func newRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:     "gitlab-op",
		Short:   "A CLI tool to operate GitLab",
		Version: version,
	}

	rootCmd.AddCommand(&cobra.Command{
		Use:   "new-group <slug> <name>",
		Short: "Add a new GitLab group",
		Args:  cobra.MaximumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			if err := CreateGroup(args[0], args[1]); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "invite <group-slug> <...emails>",
		Short: "Invite users to a GitLab group",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			if err := InviteEmails(args[0], args[1:]); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "new-project <group-slug/project-slug> <name>",
		Short: "Add a new GitLab project",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			if err := CreateProject(args[0], args[1]); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	})

	// `gitlab-op llm` prints the embedded reference for AI agents.
	llmcmd.AddTo(rootCmd, llmConfig())
	rootCmd.AddCommand(newGenerateCommand(rootCmd))

	return rootCmd
}

func main() {
	// The deprecated --llm flag has to be handled before cobra parses, because
	// cobra resolves the subcommand first and would reject it on a leaf.
	if handled, err := llmcmd.HandleLegacy(os.Args[1:], llmConfig(), os.Stdout); handled {
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		return
	}

	if err := newRootCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
