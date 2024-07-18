package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "gitlab-op",
		Short: "A CLI tool to operate GitLab",
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

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
