---
paths:
  - "main.go"
  - "config.go"
  - "app.go"
  - "internal/llmdocs/0*.md"
  - "plugins/gitlab-op/**"
  - "context7.json"
---

# Regen triggers

You are editing something the embedded LLM reference is derived from. Before
committing:

1. Run `/regen-ai` (or `go generate ./...`) and commit the result.
2. Do not hand-edit `internal/llmdocs/90-commands.md`.
3. If you changed how credentials or profiles resolve, update the
   Authentication section of `00-guide.md` and the context7 rules to match.
4. If you changed a skill description, `go test ./...` — `TestPluginSkills`
   asserts the discovery keywords are still present.
