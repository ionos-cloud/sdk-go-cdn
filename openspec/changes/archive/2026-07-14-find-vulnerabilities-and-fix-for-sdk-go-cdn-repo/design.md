## Context

`sdk-go-cdn` is a generated Go SDK (OpenAPI Generator) for the IONOS Cloud CDN API. It has not been updated since v1.2.0 (early 2024). Three independent vulnerability classes were found during a security audit:

1. **CWE-295 — Improper Certificate Validation** (`client.go:117`): The cert-pinning dialer sets `tlsConfig.InsecureSkipVerify = true` and manually checks the SHA-256 fingerprint of the server certificate. This disables hostname verification and CA-chain validation entirely. If the pinning logic fails for any reason (empty cert list edge case, future code change), there is no TLS safety net.

2. **Outdated dependencies with CVEs**:
   - `golang.org/x/net v0.0.0-20190108225652-1e06a53dbb7e` (2019) — affected by CVE-2023-44487 (HTTP/2 Rapid Reset), CVE-2022-27664, and others.
   - `google.golang.org/appengine v1.4.0` (2019), `github.com/golang/protobuf v1.2.0` (2018) — transitive, pulled in via `golang.org/x/oauth2`.
   - `go 1.18` directive — EOL since Aug 2023; stdlib security fixes are not backported.

3. **CI supply-chain risk** (`.github/workflows/build.yml`): `SonarSource/sonarcloud-github-action@master` uses a mutable ref; if the upstream repo is compromised, arbitrary code runs with `SONAR_TOKEN` and `GITHUB_TOKEN` access. `actions/checkout@v2` is also a mutable tag.

## Goals / Non-Goals

**Goals:**
- Eliminate the three vulnerability classes without changing any public API surface.
- Keep the cert-pinning feature fully functional after the TLS fix.
- Bring transitive dependencies to a clean `govulncheck`-passing state.
- Pin CI actions to SHA digests.
- Document the behavioral breaking change for consumers.

**Non-Goals:**
- Regenerating the SDK from a new OpenAPI spec.
- Adding new API endpoints or capabilities.
- Changing authentication mechanisms beyond the TLS transport fix.
- Addressing `issue-creation.yml` (already fixed in PR #4).

## Decisions

### D1 — Use `tls.Config.VerifyConnection` instead of `InsecureSkipVerify`

**Decision**: Replace `InsecureSkipVerify = true` with a `VerifyConnection` callback. With `InsecureSkipVerify` absent, the standard TLS handshake validates hostname and CA chain. `VerifyConnection` is then called as an *additional* step to assert the pinned fingerprint.

**Implementation sketch**:
```go
tlsConfig := &tls.Config{
    VerifyConnection: func(cs tls.ConnectionState) error {
        return verifyPinnedCert(fingerprint, cs.PeerCertificates)
    },
}
```
The `AddPinnedCert` public signature (`func AddPinnedCert(transport *http.Transport, pkFingerprint string)`) does not change. The hostname for standard TLS verification is already embedded in the `addr` parameter passed to the `DialTLSContext` closure and handled automatically by `tls.Dial`.

**Rationale**: `VerifyConnection` was introduced in Go 1.15 and gives full access to the verified `ConnectionState`. Setting `InsecureSkipVerify = true` disables *all* built-in TLS checks — chain, hostname, expiry — leaving cert pinning as the *only* gate. Defense-in-depth requires standard checks to also pass.

**Alternative considered**: Keep `InsecureSkipVerify = true` but add explicit `x509.Certificate.Verify` calls inside `verifyPinnedCert`. Rejected: would require constructing a verification options struct with hostname, duplicating stdlib logic and introducing new error surfaces.

### D2 — Bump `golang.org/x/oauth2` to latest; let `go mod tidy` pull updated transitive deps

**Decision**: Update only the direct dependency (`golang.org/x/oauth2`). The CVE-bearing packages (`x/net`, `appengine`, `protobuf`) are transitive. Running `go get golang.org/x/oauth2@latest && go mod tidy` will resolve them to current safe versions.

**Rationale**: Direct dependency on `x/net` is not declared in `go.mod`; forcing it as a direct dep would add unnecessary noise. The transitive pull is correct Go module hygiene.

**Alternative considered**: Pin `golang.org/x/net` as a direct `require` entry at a safe version. Rejected: indirect deps should stay indirect.

### D3 — Update `go` directive to `1.21`

**Decision**: Bump `go 1.18` → `go 1.21` in `go.mod`.

**Rationale**: Go 1.21 is the oldest maintained release. It also enables the `toolchain` directive for reproducible builds. The SDK uses no APIs removed between 1.18 and 1.21.

### D4 — Pin GitHub Actions to full commit SHA

**Decision**: Replace `actions/checkout@v2` and `SonarSource/sonarcloud-github-action@master` with pinned SHA digests.

**Rationale**: Tags and branch names are mutable refs. A compromised upstream repo can push malicious code to `master` or retag `v2`. SHA pins are immutable. Add a comment with the human-readable version for maintainability.

## Risks / Trade-offs

- **[Risk] `VerifyConnection` behaviour change**: With `InsecureSkipVerify = false`, the standard TLS handshake now validates the server cert's hostname and CA chain. Callers using `IONOS_PINNED_CERT` with a self-signed certificate (no matching system CA) will now get a connection error where previously they succeeded. This is a **behavioral breaking change**.
  → **Mitigation**: Intentional and correct. Document in CHANGELOG. Self-signed cert users must either add the CA to the system trust store or accept the stricter posture.

- **[Risk] Transitive dep bumps may introduce breaking changes**: `x/net`, `appengine`, and `protobuf` are internal to oauth2; unlikely to affect the SDK's public surface.
  → **Mitigation**: Run `go build ./...` and `go vet ./...` after `go mod tidy`.

- **[Risk] SHA-pinned CI actions need periodic manual bumps**: Unlike tags, SHA pins don't auto-update.
  → **Mitigation**: Accept as the correct security trade-off. Recommend adding Dependabot config as a follow-up.

## Migration Plan

1. Apply TLS fix in `client.go` — remove `InsecureSkipVerify = true`, add `VerifyConnection` callback.
2. Run `go get golang.org/x/oauth2@latest && go mod tidy` — update `go.mod` / `go.sum`.
3. Update `go` directive in `go.mod` to `1.21`.
4. Update `build.yml` with pinned SHA digests; add inline comments with tag name + date.
5. Write automated tests covering all five TLS scenarios from `specs/tls-insecure-skip-verify/spec.md` using `httptest.NewTLSServer`.
6. Add CHANGELOG entry documenting the behavioral breaking change: callers using `IONOS_PINNED_CERT` with a self-signed certificate that is not in the system CA trust store will now receive a TLS error (hostname/chain validation failure) where connections previously succeeded under the `InsecureSkipVerify` bypass.
7. Run `govulncheck ./...` to confirm zero known CVEs remain.
8. Open PR; CI (SonarCloud) will gate merge.

**Rollback**: All changes are isolated files. Revert `client.go`, `go.mod`, `go.sum`, `build.yml`, and `CHANGELOG.md` independently if needed.

## Open Questions

- Should `govulncheck` be added as a CI step to prevent future regressions? (Out of scope for this change but recommended as a follow-up.)
- Should Dependabot be configured for Go modules and GitHub Actions? (Out of scope; recommend as follow-up.)
