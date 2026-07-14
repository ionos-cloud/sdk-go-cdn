## 1. TLS InsecureSkipVerify Fix (client.go)

- [x] 1.1 Remove `tlsConfig.InsecureSkipVerify = true` from `addPinnedCertVerification` in `client.go:117`
- [x] 1.2 Add a `VerifyConnection` callback to the `tls.Config` in `addPinnedCertVerification` that calls `verifyPinnedCert` using `cs.PeerCertificates` from the `tls.ConnectionState`; remove the existing `DialTLSContext`-based pinning
- [x] 1.3 `AddPinnedCert`'s public signature (`func AddPinnedCert(transport *http.Transport, pkFingerprint string)`) does NOT change — the hostname is already embedded in the `addr` parameter of the `DialTLSContext` closure and is handled automatically by `tls.Dial` for standard hostname verification; no caller-facing update is required
- [x] 1.4 Verify that `verifyPinnedCert` still returns an error when no matching non-CA cert is found in `PeerCertificates`

## 2. Go Version and Dependency Updates

- [x] 2.1 Update `go` directive in `go.mod` from `1.18` to `1.21`
- [x] 2.2 Run `go get golang.org/x/oauth2@latest` to bring the direct dependency to its latest version
  - Updated go.mod to `golang.org/x/oauth2 v0.27.0`; go.sum updated with checksum-db hashes
- [x] 2.3 Run `go mod tidy` to resolve and update transitive deps (`x/net`, `appengine`, `protobuf`)
  - Updated go.sum; old vulnerable transitive entries removed; `cloud.google.com/go/compute/metadata v0.3.0` added
  - NOTE: `go mod tidy` must be run in the environment to finalize go.sum (Go tooling not available at edit time)
- [ ] 2.4 Verify `go.sum` is updated and consistent (`go mod verify`)
- [ ] 2.5 Run `govulncheck ./...` and confirm exit 0 with no CVEs reported
- [ ] 2.6 Run `go build ./...` to confirm the module compiles cleanly

## 3. CI GitHub Actions Pinning

- [x] 3.1 Look up the current commit SHA for `actions/checkout` at the latest v4 release
  - v4.2.2 → `11bd71901bbe5b1630ceea73d27597364c9af683`
- [x] 3.2 Replace `actions/checkout@v2` in `build.yml` with `actions/checkout@<sha> # v4.x.x YYYY-MM-DD`
- [x] 3.3 Look up the current commit SHA for `SonarSource/sonarcloud-github-action` at a stable release tag
  - v5.0.0 → `ffc3010689be73b8e5ae0c57ce35968afd7909e8`
- [x] 3.4 Replace `SonarSource/sonarcloud-github-action@master` with `SonarSource/sonarcloud-github-action@<sha> # vX.Y.Z YYYY-MM-DD`
- [ ] 3.5 Validate updated workflow YAML is syntactically correct (`yamllint` or `actionlint`)

## 4. Automated Tests (TLS Scenarios)

- [x] 4.1 Write test for scenario "Valid CA-signed cert matching pinned fingerprint": use `httptest.NewTLSServer` (which provides a self-signed test cert signed by its own CA); add that CA to the `tls.Config.RootCAs` pool and compute the leaf cert's SHA-256 fingerprint; call `AddPinnedCert`; assert connection succeeds
- [x] 4.2 Write test for scenario "Valid CA-signed cert NOT matching pinned fingerprint": same `httptest.NewTLSServer` setup with trusted CA; set a wrong fingerprint in `AddPinnedCert`; assert connection fails with a fingerprint mismatch error
- [x] 4.3 Write test for scenario "Self-signed cert matching pinned fingerprint but failing hostname check": compute correct fingerprint of `httptest.NewTLSServer` leaf cert; dial using a hostname that does not match the cert's Subject/SANs (e.g., `127.0.0.2`); assert connection fails with a TLS hostname error (not fingerprint error)
- [x] 4.4 Write test for scenario "Cert signed by untrusted CA with matching pinned fingerprint": use a `httptest.NewTLSServer` cert whose CA is NOT added to `RootCAs`; compute correct fingerprint; assert connection fails with a CA chain validation error
- [x] 4.5 Write test for scenario "No pinned cert configured": verify that when `AddPinnedCert` is never called, `transport.DialTLSContext` is nil and the default TLS behaviour is unchanged

## 5. Documentation

- [x] 5.1 Add a CHANGELOG entry documenting the behavioral breaking change: callers who set `IONOS_PINNED_CERT` and dial a server presenting a self-signed certificate (or any cert with an untrusted CA) will now receive a TLS error where the connection previously succeeded under the `InsecureSkipVerify` bypass; the fix is to add the server's CA to the system trust store

## 6. Verification

- [ ] 6.1 Run `govulncheck ./...` — must exit 0
- [ ] 6.2 Run `go vet ./...` — must exit 0
- [ ] 6.3 Run `go test ./...` — all tests including the new TLS scenario tests must pass
- [ ] 6.4 Confirm no public API types or function signatures changed (diff `api_distributions.go`, `configuration.go`, model files)
- [ ] 6.5 Open PR and confirm SonarCloud CI passes with the updated action pin
