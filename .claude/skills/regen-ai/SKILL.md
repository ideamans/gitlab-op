---
name: regen-ai
description: Regenerate the embedded LLM reference for gitlab-op and verify it still builds and passes the plugin checks.
---

# regen-ai

```bash
go generate ./...
git diff --stat -- internal/llmdocs
go test ./...
go run . llm | head -5
```

Report what changed in `internal/llmdocs/`, and whether the tests pass.
