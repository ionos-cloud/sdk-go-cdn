# Changelog

## [Unreleased]

### Security

- **BREAKING: TLS cert-pinning no longer bypasses standard verification** (`client.go`)
  - `addPinnedCertVerification` previously set `InsecureSkipVerify = true`, disabling hostname
    and CA-chain validation entirely. The pinned-fingerprint check was the only TLS gate.
  - The fix replaces `InsecureSkipVerify` with a `tls.Config.VerifyConnection` callback.
    Standard TLS verification (hostname, CA chain, expiry) now runs first; fingerprint
    matching is applied as an additional constraint afterwards.
  - **Migration**: callers who set `IONOS_PINNED_CERT` and connect to a server presenting a
    self-signed certificate (or any certificate whose CA is not in the system trust store)
    will now receive a TLS error where the connection previously succeeded. To restore
    connectivity, add the server's CA certificate to the system trust store, or supply a
    custom `RootCAs` pool in the `tls.Config` passed to `addPinnedCertVerification`.

- **CVE remediation: dependency updates** (`go.mod`)
  - Updated `golang.org/x/oauth2` to v0.27.0 (was v0.21.0).
  - Transitive dependencies (`golang.org/x/net`, `google.golang.org/appengine`,
    `github.com/golang/protobuf`) updated via `go mod tidy`, resolving CVE-2023-44487
    (HTTP/2 Rapid Reset), CVE-2022-27664, and related issues in `golang.org/x/net`.
  - `go` directive bumped from `1.18` (EOL) to `1.21`.

- **CI supply-chain hardening** (`.github/workflows/build.yml`)
  - `actions/checkout@v2` (mutable tag) replaced with
    `actions/checkout@11bd71901bbe5b1630ceea73d27597364c9af683` (v4.2.2, immutable SHA).
  - `SonarSource/sonarcloud-github-action@master` (mutable branch) replaced with
    `SonarSource/sonarcloud-github-action@ffc3010689be73b8e5ae0c57ce35968afd7909e8`
    (v5.0.0, immutable SHA).
