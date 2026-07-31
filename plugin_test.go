package main

import (
	"testing"

	"github.com/ideamans/go-llm-cli-kit/skillcheck"
)

// TestPluginSkills validates the distributed Claude Code plugin: the manifest
// version must track the CLI version, the SKILL.md frontmatter must stay within
// the Agent Skills standard, and the descriptions must still carry the terms a
// user says when they need this tool.
func TestPluginSkills(t *testing.T) {
	report := skillcheck.CheckDir("plugins/gitlab-op", skillcheck.Options{
		Version:             PluginVersion,
		Keywords:            []string{"gitlab", "group", "install"},
		RequireInstallSkill: true,
	})
	for _, problem := range report.Problems {
		t.Error(problem)
	}
}
