# Template AI Native Verification Report

**Project:** Media Belajar Anak
**Date:** 2026-08-09
**Template source:** `setiyadijoko/template-ai-native`
**Repository:** `setiyadinamikaintegrasi/media-belajar-anak`
**PR:** #1
**Bootstrap commit:** `51a9a78`
**Component workflow activation commit:** `0d92819`
**P0 revalidation commits:** `4176c4f`, `d10097e`
**Successful hosted run:**
[31312862775](https://github.com/setiyadinamikaintegrasi/media-belajar-anak/actions/runs/31312862775)

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

PR: [#1 — bootstrap media belajar anak monorepo](https://github.com/setiyadinamikaintegrasi/media-belajar-anak/pull/1)

### General template checks

Passed:

- Detect stack
- PR Title Check
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

Failed on the initial run:

- Markdown lint: bare URL in this verification document.
- Generic dispatcher: reusable component jobs did not materialize.

### Component-aware workflow

The initial generic `ci.yml` run reported `component_ready=true` but did not
create the component job. A temporary consumer-owned caller workflow was added:

```text
.github/workflows/media-belajar-components.yml
```

After that direct activation, the reusable component workflow ran successfully
through resolver and build stages.

### P0 template revalidation

The branch now synchronizes the corrected P0 runtime from template commit
`88b3c5f` and removes the temporary caller. The revalidation patch includes:

- isolated reusable-workflow concurrency groups;
- deterministic `golangci-lint` v2.12.2 setup;
- explicit skip behavior for empty optional Node.js integration/e2e categories;
- fail-closed, visible 80% Go coverage enforcement;
- the dedicated initializer README fixture; and
- backend tests that exercise greeting, writer, and executable output behavior
  at 80% measured coverage.

The generic `ci` workflow successfully materialized the resolver, backend and
frontend quality/test/build jobs, and aggregate job without a consumer-owned
caller. All component jobs passed. The same PR revision also passed CodeQL,
gitleaks, OSV Scanner, actionlint, zizmor, metadata, documentation, and the
production-readiness contract. Single-stack jobs skipped as designed because
the repository declares a version-2 monorepo.

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

**Resolution:** replaced by the dedicated upstream fixture during P0
revalidation.

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

**Resolution:** fixed by upstream concurrency isolation. Hosted run
`31312862775` materialized and passed the resolver, six component jobs, and
aggregate job through the generic dispatcher alone.

### Finding 3 — Go component workflow assumes golangci-lint is preinstalled

Backend quality failed on GitHub:

```text
sh: 1: golangci-lint: not found
```

`ci-monorepo.yml` invokes `golangci-lint run`, but does not install or set up
`golangci-lint`. A fresh Go component cannot pass quality checks zero-config.
The template should pin/install the linter or document the requirement.

**Resolution:** the SHA-pinned upstream installer is included in revalidation.

### Finding 4 — Go coverage gate is strict for bootstrap components

The minimal backend produced:

```text
Go coverage: 50.0%
Go coverage 50.0% < 80%
```

The 80% threshold is reasonable for production, but the template should provide
a documented bootstrap path or fixture with sufficient coverage.

**Resolution:** the 80% gate remains unchanged; the backend now meets it with
behavioral tests and the threshold is printed in CI.

### Finding 5 — frontend integration/e2e directories must contain tests

Frontend unit test passed, but the reusable workflow failed at integration test:

```text
No test files found, exiting with code 1
```

The Node mapper should either no-op empty test categories or the getting-started
guide should require placeholder tests in every test directory.

**Resolution:** optional empty categories now skip explicitly; real failures
remain blocking.

## Verdict

The P0 template fixes are validated on a real Go and Node.js consumer monorepo.
The generic dispatcher now activates the complete component-aware CI path
without a workaround, all required component jobs pass, and the original five
adoption gaps are closed for this pilot. Deployment and real-model evaluation
remain consumer-owned skeletons by design and were not validated here.
