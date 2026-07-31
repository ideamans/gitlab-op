package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ideamans/go-llm-cli-kit/catalog"
	"github.com/spf13/cobra"
)

// docsDir is where the embedded reference chapters live, relative to this
// package's directory (go generate runs in the directory of the directive).
const docsDir = "./internal/llmdocs"

// newGenerateCommand returns the hidden command that regenerates the derived
// chapters of the embedded reference.
//
// It lives in package main because the command tree is defined here; an
// independent internal/gen-llmdocs program could not import it without
// splitting the CLI into another package first.
func newGenerateCommand(root *cobra.Command) *cobra.Command {
	return &cobra.Command{
		Use:    "gen-llmdocs",
		Short:  "regenerate the embedded LLM reference (development only)",
		Hidden: true,
		Args:   cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			md := catalog.Markdown(root, catalog.Options{
				Title: "Command catalog",
				Intro: "Generated from the cobra command tree by `go generate ./...`.\n" +
					"Do not edit by hand — edit the command definitions instead.",
				// llm documents itself in 00-guide.md; gen-llmdocs is a
				// development command and must not be advertised to agents.
				Skip: []string{"llm", "gen-llmdocs"},
			})

			path := filepath.Join(docsDir, "90-commands.md")
			if err := os.WriteFile(path, []byte(md), 0o644); err != nil {
				return err
			}
			fmt.Fprintf(cmd.ErrOrStderr(), "wrote %s (%d bytes)\n", path, len(md))
			return nil
		},
	}
}
