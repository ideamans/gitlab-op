---
name: gitlab-op-usage
description: Create GitLab groups and projects and invite people to them, using the gitlab-op CLI. Use when the user asks to set up a new GitLab group or project, onboard someone onto a GitLab group, invite an email address to GitLab, or provision GitLab structure for a new client or team — on a self-hosted GitLab as well as gitlab.com.
license: MIT
compatibility: Requires the `gitlab-op` binary on PATH — run the gitlab-op-install skill if it is missing — and a credentials profile at ~/.gitlab-op/credentials containing a GitLab URL and an API-scoped token. Every command writes to the live instance; there is no dry-run.
allowed-tools: Bash(gitlab-op:*) Bash(command:*) Bash(test:*) Read
---

# gitlab-op-usage

Provisioning GitLab groups, projects and memberships.

## 1. Confirm the tool and the profile

```bash
command -v gitlab-op && test -f ~/.gitlab-op/credentials && echo "credentials present"
```

Missing binary? Run the `gitlab-op-install` skill. Missing credentials file?
Tell the user to create `~/.gitlab-op/credentials` as INI, with a `[default]`
section holding `url` and `token`. **Never ask for the token in the
conversation and never echo it back** — it belongs only in that file.

The active profile comes from `GITLAB_OP_PROFILE` and defaults to `default`.
There is no `--profile` flag. If the user works across several GitLab
instances, confirm which profile before running anything.

## 2. Everything here is a live write

There is no `--dry-run`, no confirmation prompt and no undo in this CLI.
Creating a group takes effect immediately, and `invite` sends real email to
real people — some of whom may not have an account yet.

Before running: state the instance (from the profile), the exact slug, and for
`invite` the full list of addresses. Get agreement. This matters most when the
user has more than one profile, because the wrong `GITLAB_OP_PROFILE` puts a
group on the wrong client's GitLab.

## 3. Pick the command

| The user wants | Command |
| --- | --- |
| a new group | `gitlab-op new-group <slug> <name>` |
| a new project in a group | `gitlab-op new-project <group-slug/project-slug> <name>` |
| to add people to a group | `gitlab-op invite <group-slug> <email> [email…]` |

Slugs are URL path segments, not display names. `new-group` takes the slug
first and the human-readable name second.

Order matters: create the group before creating a project in it or inviting
anyone to it — both look it up by slug and fail if it does not exist.

## 4. Read the reference for anything else

```bash
gitlab-op llm
gitlab-op new-group --help
```

## 5. Report what was created

Say the group or project path and the instance it landed on. For `invite`,
list who was invited. If a command failed partway through a sequence, say
exactly which steps succeeded — there is no rollback, so a half-provisioned
group is a state the user needs to know about.

## Failure modes

| Symptom | Fix |
| --- | --- |
| `command not found: gitlab-op` | run the `gitlab-op-install` skill |
| error mentioning `.gitlab-op/credentials` | the file does not exist; the user must create it |
| `profile <name> not found` | `GITLAB_OP_PROFILE` names a missing section — check the file, or unset it for `default` |
| 401 / 403 | token wrong, expired, or lacking API scope |
| group not found | wrong slug, or the token cannot see that group |
| project creation fails, group exists | pass the full `group-slug/project-slug`, not just the project slug |
