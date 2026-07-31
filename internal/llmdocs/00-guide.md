# gitlab-op — reference for AI agents

`gitlab-op` wraps the handful of GitLab administration jobs that get done
repeatedly: creating a group, creating a project inside one, and inviting people
to a group. It is deliberately small — it is not a general GitLab client.

Nothing prompts; it reads arguments and a credentials file only. Errors print
and the exit code is non-zero on failure. This reference is embedded in the
binary, so `gitlab-op llm` always describes the exact version you are running.

## Ground rules

1. **Every command writes to a live GitLab instance.** There is no `--dry-run`
   and no confirmation prompt. Creating a group or inviting people takes effect
   immediately, so confirm the target with the user before running anything.
2. **Credentials come from a profile file, not the environment.** A missing or
   misnamed profile is the most common failure — see below.
3. **Groups are addressed by slug, projects by `group-slug/project-slug`.**
   The slug is the URL path segment, not the display name. `new-group` takes
   both: the slug first, then the human-readable name.
4. **Invitations are by email address.** They reach people who may not have an
   account yet, so treat sending them as an outward-facing action.

## Authentication

Credentials live in an INI file at **`~/.gitlab-op/credentials`**, with one
section per profile:

```ini
[default]
url   = https://gitlab.example.com
token = glpat-xxxxxxxxxxxxxxxxxxxx

[client-a]
url   = https://gitlab.client-a.example
token = glpat-yyyyyyyyyyyyyyyyyyyy
```

The profile is chosen by the **`GITLAB_OP_PROFILE`** environment variable and
defaults to `default`. There is no `--profile` flag; export the variable for the
command or the session:

```bash
GITLAB_OP_PROFILE=client-a gitlab-op new-group demo "Demo Group"
```

The token needs GitLab API scope sufficient for the operation — creating groups
and projects requires more than a read-only token. **Never echo the token back
into the conversation**, and never write it anywhere but that file.

## Commands

| The user wants | Command |
| --- | --- |
| A new group | `gitlab-op new-group <slug> <name>` |
| A new project inside a group | `gitlab-op new-project <group-slug/project-slug> <name>` |
| To add people to a group | `gitlab-op invite <group-slug> <email> [email…]` |

`invite` takes any number of addresses after the group slug.

## Typical flow

```bash
export GITLAB_OP_PROFILE=default

gitlab-op new-group acme "ACME Corporation"
gitlab-op new-project acme/website "ACME Website"
gitlab-op invite acme alice@example.com bob@example.com
```

Group creation has to come first: `new-project` and `invite` both look the group
up by slug and fail if it does not exist.

## Failure modes

| Symptom | Cause | Fix |
| --- | --- | --- |
| `no such file or directory` mentioning `.gitlab-op/credentials` | the credentials file has never been created | write it, with at least a `[default]` section |
| `profile <name> not found` | `GITLAB_OP_PROFILE` names a section that is not in the file | check the section names, or unset the variable to use `default` |
| 401 / 403 from GitLab | the token is wrong, expired, or lacks scope | reissue a token with API scope for the operation |
| group not found | the slug is wrong, or the token cannot see that group | slugs are URL path segments, not display names |
| project creation fails but the group exists | the argument is not in `group-slug/project-slug` form | pass the full path |

## What this CLI will not do

- No deletion, renaming or permission editing. Creation and invitation only.
- No issue, merge request, pipeline or repository operations — use `glab` or the
  API for those.
- No dry-run mode. To check something first, look at it in the GitLab UI.
