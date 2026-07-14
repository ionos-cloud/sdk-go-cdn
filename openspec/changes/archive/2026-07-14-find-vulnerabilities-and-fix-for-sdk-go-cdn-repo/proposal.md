## Why

The `sdk-go-cdn` module contains three classes of security vulnerabilities: an unsafe TLS configuration in the cert-pinning path, severely outdated Go dependencies with known CVEs (oldest: 2018), and unpinned GitHub Actions workflow steps that create a supply-chain risk. These must be fixed before the SDK can be considered safe for production use.

## What Changes

- **BREAKING** `client.go`: Remove `tlsConfig.InsecureSkipVerify = true` from `addPinnedCertVerification`; replace with a `tls.Config.VerifyConnection` callback that enforces the pinned fingerprint as an *additional* check after the standard TLS handshake. Self-signed certificates that previously succeeded under the `InsecureSkipVerify` bypass will now be rejected by standard hostname and CA-chain validation.
- `go.mod` / `go.sum`: Bump `golang.org/x/oauth2` to latest; `golang.org/x/net` (transitive, CVE-2023-44487, CVE-2022-27664) upgrades automatically via `go mod tidy`; update `go` directive from `1.18` (EOL) to `1.21`.
- `.github/workflows/build.yml`: Pin `actions/checkout` and `SonarSource/sonarcloud-github-action` to immutable 40-character commit SHAs instead of mutable tags/branches to prevent supply-chain attacks.

## Capabilities

### New Capabilities

- `tls-insecure-skip-verify`: Secure the cert-pinning TLS dialer so that standard hostname verification and CA-chain validation remain active while still enforcing the pinned fingerprint.
- `dependency-updates`: Update Go module dependencies to eliminate CVEs in `golang.org/x/net` and related packages, and advance the `go` directive to a supported version.
- `ci-action-pinning`: Pin all third-party GitHub Actions workflow steps to immutable SHA digests to prevent supply-chain attacks.

### Modified Capabilities

*(none — no existing openspec specs exist for this repo)*

## Impact

- **`client.go`**: `addPinnedCertVerification` and `AddPinnedCert` — behaviour change in TLS dial path. **Breaking**: callers using `IONOS_PINNED_CERT` with a self-signed cert that has no matching system CA entry will now receive a TLS error.
- **`go.mod` / `go.sum`**: Dependency graph changes; consumers who vendor will need to re-vendor.
- **`.github/workflows/build.yml`**: CI only; no runtime impact.
- **No public API surface changes** — all changes are internal to transport and security configuration.
