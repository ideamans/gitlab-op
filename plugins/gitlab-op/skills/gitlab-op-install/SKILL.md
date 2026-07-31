---
name: gitlab-op-install
description: Make the gitlab-op command available, installing it only if it is missing. Use when another skill reports that `gitlab-op` is not on PATH, or when the user asks to install, update or upgrade the ideamans GitLab operations CLI. Prefers an already-installed binary, then the latest GitHub release, then a build from source with go install.
license: MIT
compatibility: Requires curl and tar to install from a release, or a Go toolchain for the source fallback. Standalone — does not need gitlab-op to be present already. Installs from the public repository github.com/ideamans/gitlab-op, so no GitHub authentication is needed.
allowed-tools: Bash(curl:*) Bash(wget:*) Bash(tar:*) Bash(unzip:*) Bash(go:*) Bash(uname:*) Bash(command:*) Bash(which:*) Bash(mkdir:*) Bash(mv:*) Bash(cp:*) Bash(rm:*) Bash(chmod:*) Bash(ls:*) Bash(test:*) Bash(echo:*) Read
---

# gitlab-op-install

Make the `gitlab-op` command usable, doing the least work that achieves it.

## Route 1 — an existing installation on PATH

```bash
command -v gitlab-op && gitlab-op --version
```

If that resolves, **use it and stop here.** Do not check for a newer release —
it costs an API call and the user did not ask for an upgrade.

Two checks before trusting the hit:

- **It is the right tool.** `gitlab-op llm | head -1` must read
  `# gitlab-op — reference for AI agents`. If something else owns the name,
  say so and use an explicit path rather than shadowing theirs.
- **It is recent enough.** If `llm` is not a known command, the binary predates
  the embedded reference — continue to route 2 to upgrade it.

## Route 2 — the latest GitHub release

The repository is public, so no authentication is needed.

```bash
VERSION=$(curl -fsSL https://api.github.com/repos/ideamans/gitlab-op/releases/latest \
  | grep '"tag_name"' | head -1 | cut -d'"' -f4)
OS=$(uname -s)                                          # Darwin | Linux
ARCH=$(uname -m); [ "$ARCH" = "amd64" ] && ARCH=x86_64  # x86_64 | arm64
curl -fsSL -o /tmp/gitlab-op.tar.gz \
  "https://github.com/ideamans/gitlab-op/releases/download/${VERSION}/gitlab-op_${OS}_${ARCH}.tar.gz"
```

The archive name follows goreleaser's `uname`-compatible template, so the OS is
capitalised (`Darwin`, `Linux`) and the architecture is `x86_64` rather than
`amd64`. Windows ships a `.zip`.

**Verify the pattern against the release page before trusting it** — this
project's first release is recent, so if the download 404s, list the actual
assets rather than retrying variations:

```bash
curl -fsSL https://api.github.com/repos/ideamans/gitlab-op/releases/latest \
  | grep '"name"' | grep -E 'tar.gz|zip'
```

### Install onto PATH

```bash
tar -xzf /tmp/gitlab-op.tar.gz -C /tmp
mkdir -p ~/.local/bin && mv /tmp/gitlab-op ~/.local/bin/ && chmod +x ~/.local/bin/gitlab-op
```

Prefer the first writable directory already on PATH — `~/.local/bin`, then
`/usr/local/bin`. Two things not to do on your own initiative:

- If nothing on PATH is writable, leave the binary in `/tmp`, print the exact
  `sudo mv` command and let the user run it. Do not run `sudo` yourself.
- If `~/.local/bin` is not on PATH, give the user the line for their shell
  profile. Do not edit the profile for them.

## Route 3 — build from source

```bash
go install github.com/ideamans/gitlab-op@latest
```

Needs a Go toolchain. The module path matches the repository, so this produces
a command called `gitlab-op`.

## Verify

```bash
command -v gitlab-op && gitlab-op --version && gitlab-op llm | head -1
```

Report the version and the path. Then say what is still needed: a credentials
file at `~/.gitlab-op/credentials` with a `[default]` section containing `url`
and an API-scoped `token`. Without it every command fails. Do not create the
file with a placeholder token — let the user supply it.
