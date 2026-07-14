# ci-action-pinning Specification

## Purpose
TBD - created by archiving change find-vulnerabilities-and-fix-for-sdk-go-cdn-repo. Update Purpose after archive.
## Requirements
### Requirement: Third-party GitHub Actions steps MUST be pinned to immutable SHA digests
All `uses:` references to third-party GitHub Actions in `.github/workflows/*.yml` SHALL specify a full 40-character commit SHA rather than a mutable tag or branch name. A comment on the same line MUST record the human-readable version and pin date for maintainability.

#### Scenario: actions/checkout is pinned to a SHA
- **WHEN** `.github/workflows/build.yml` is read
- **THEN** the `uses: actions/checkout` line MUST reference a full 40-character SHA (e.g., `actions/checkout@<sha> # v4.x.x YYYY-MM-DD`)

#### Scenario: SonarSource/sonarcloud-github-action is pinned to a SHA
- **WHEN** `.github/workflows/build.yml` is read
- **THEN** the `uses: SonarSource/sonarcloud-github-action` line MUST NOT reference `@master` or any other branch/tag; it MUST reference a full 40-character SHA with an inline version comment

#### Scenario: No workflow step uses a mutable branch ref
- **WHEN** all `.github/workflows/*.yml` files are scanned for `uses:` directives
- **THEN** no third-party action (any org other than `actions/`) SHALL use `@master`, `@main`, `@latest`, or a semver tag without a SHA pin

