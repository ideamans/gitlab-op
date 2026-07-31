// Package llmdocs embeds the reference that `gitlab-op llm` prints.
//
// 00-*.md .. 89-*.md are hand-written. 90-commands.md is generated from the
// cobra command tree — run `go generate ./...` after changing commands or
// flags, and commit the result. CI re-runs it and fails on a dirty tree.
package llmdocs

import (
	"embed"

	kit "github.com/ideamans/go-llm-cli-kit/llmdocs"
)

//go:embed *.md
var files embed.FS

// Docs is the embedded reference bundle.
func Docs() *kit.Docs { return kit.New(files, ".") }
