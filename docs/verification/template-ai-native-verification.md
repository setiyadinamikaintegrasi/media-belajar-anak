# Template AI Native Verification Report

**Project:** Media Belajar Anak
**Date:** 2026-08-09
**Template source:** `setiyadijoko/template-ai-native`
**Repository:** `setiyadinamikaintegrasi/media-belajar-anak`
**PR:** #1
**Bootstrap commit:** `51a9a78`
**Component workflow activation commit:** `0d92819`

## Execution record

### Repository creation

- GitHub Template flag was enabled on `template-ai-native`.
- Repository was created via GitHub **Use this template**.
- Initial repository commit: `1584032`.
- Clone command:

```sh
git clone https://github.com/setiyadinamikaintegrasi/media-belajar-anak.git
cd media-belajar-anak
```

### Initializer

```sh
sh scripts/init-project.sh \
  --name media-belajar-anak \
  --description "Media belajar anak yang interaktif, aman, dan menyenangkan" \
  --stack auto \
  --layout monorepo \
  --primary-path src/backend \
  --component backend=src/backend:go \
  --component frontend=src/frontend:node
```

Generated contract:

```yaml
version: 2
layout: monorepo
primary_stack: auto
primary_path: src/backend
components:
  - id: backend
    path: src/backend
    stack: go
    required: true
    artifact: backend
  - id: frontend
    path: src/frontend
    stack: node
    required: true
    artifact: frontend
```

No symlink was used.

## Local verification

| Command | Result |
|---|---|
| `make setup` | ✅ Exit 0; expected monorepo no-op message |
| `make test-scripts` | ✅ Exit 0 after consumer fixture normalization |
| `make docs-check` | ✅ Exit 0; optional local tools skipped |
| `make ci` | ✅ Exit 0 |
| `make project-config-check` | ✅ Valid |
| `resolve-components.sh --json` | ✅ Two components resolved |
| Go format/test/vet/build | ✅ Pass |
| `npm ci` | ✅ 0 vulnerabilities |
| Frontend format/lint/typecheck/test/build | ✅ Pass |
| No symlink under `src/` | ✅ None |
| `git diff --check` | ✅ Pass |

## Pull request verification

PR #1: https://github.com/setiyadinamikaintegrasi/media-belajar-anak/pull/1

### General template checks

Passed:

- Detect stack
- PR Title Check
- Markdown lint/link/TODO check
- Required docs and metadata
- Production readiness contract
- gitleaks
- OSV scanner
- CodeQL
- actionlint
- zizmor
- Code Review Graph
- OpenCodeReview advisory workflow
- License policy
- AI evaluation advisory workflow

### Component-aware workflow

The generic `ci.yml` reported `component_ready=true` but did not create the
component job. A consumer-owned caller workflow was added:

```text
.github/workflows/media-belajar-components.yml
```

After that direct activation, the reusable component workflow ran successfully
through resolver and build stages.

## Findings

### Finding 1 — initializer contract tests are not consumer-safe

After `init-project.sh` generates the README identity marker, the inherited
`test-init-project.sh` copies that generated README as if it were a fresh
fixture. Three tests fail:

```text
FAIL initializes a fresh README
FAIL writes project identity to README
FAIL preserves README after refused reconfiguration
```

Consumer-side normalization was added so the test reconstructs a marker-only
fixture before testing the initializer. This should be upstreamed to the
original template.

### Finding 2 — generic dispatcher does not activate component job

GitHub log showed:

```text
Detected layout: monorepo
Detected stack: unknown
component_ready=true
```

But the `Component-aware monorepo CI` job was absent from the `ci` workflow.
A direct consumer caller workflow was required to exercise `ci-monorepo.yml`.
The template dispatcher activation should be reviewed upstream.

### Finding 3 — Go component workflow assumes golangci-lint is preinstalled

Backend quality failed on GitHub:

```text
sh: 1: golangci-lint: not found
```

`ci-monorepo.yml` invokes `golangci-lint run`, but does not install or set up
`golangci-lint`. A fresh Go component cannot pass quality checks zero-config.
The template should pin/install the linter or document the requirement.

### Finding 4 — Go coverage gate is strict for bootstrap components

The minimal backend produced:

```text
Go coverage: 50.0%
Go coverage 50.0% < 80%
```

The 80% threshold is reasonable for production, but the template should provide
a documented bootstrap path or fixture with sufficient coverage.

### Finding 5 — frontend integration/e2e directories must contain tests

Frontend unit test passed, but the reusable workflow failed at integration test:

```text
No test files found, exiting with code 1
```

The Node mapper should either no-op empty test categories or the getting-started
guide should require placeholder tests in every test directory.

## Verdict

The template is strong for governance and security controls. The v2 monorepo
contract resolves and builds correctly, but the first real GitHub component run
exposed four adoption gaps around dispatcher activation, Go linter setup,
bootstrap coverage, and empty Node test categories.
