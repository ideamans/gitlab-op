# AI artifact policy

`gitlab-op llm` prints an embedded reference assembled from
`internal/llmdocs/`. One chapter is generated; the rest are hand-written.

## Never hand-edit

| File | Produced by |
| --- | --- |
| `internal/llmdocs/90-commands.md` | `go generate ./...` (the hidden `gen-llmdocs` subcommand walks the cobra tree) |

## Source of truth

| To change… | Edit |
| --- | --- |
| ground rules, credentials, failure modes | `internal/llmdocs/00-guide.md` |
| a command description | the cobra definition in `main.go` |
| what the distributed skills tell an agent to do | `plugins/gitlab-op/skills/*/SKILL.md` |
| pitfalls surfaced through context7 | `context7.json` `rules` |

This CLI has no dry-run and no undo. If you add a command that deletes or
changes permissions, that fact belongs in the ground rules of `00-guide.md`
and in the usage skill before it belongs anywhere else.
