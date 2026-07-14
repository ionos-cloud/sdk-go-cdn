## ADDED Requirements

### Requirement: Cert-pinning dialer SHALL NOT disable standard TLS verification
The `addPinnedCertVerification` function MUST NOT set `InsecureSkipVerify = true` on the `tls.Config`. Standard hostname verification and CA-chain validation SHALL remain active by relying on the default TLS handshake. The pinned-fingerprint check SHALL be applied as an additional constraint via `tls.Config.VerifyConnection`, which runs after the standard TLS handshake succeeds.

#### Scenario: Valid CA-signed cert matching pinned fingerprint
- **WHEN** a client dials a TLS endpoint with a valid CA-signed certificate whose SHA-256 fingerprint matches the configured `IONOS_PINNED_CERT` value
- **THEN** the connection MUST be established successfully

#### Scenario: Valid CA-signed cert NOT matching pinned fingerprint
- **WHEN** a client dials a TLS endpoint with a valid CA-signed certificate whose SHA-256 fingerprint does NOT match the configured `IONOS_PINNED_CERT` value
- **THEN** the connection MUST be rejected with an error indicating fingerprint mismatch

#### Scenario: Self-signed cert matching pinned fingerprint but failing hostname check
- **WHEN** a client dials a TLS endpoint presenting a self-signed certificate whose SHA-256 fingerprint matches the pin but whose subject does not match the dialed hostname
- **THEN** the connection MUST be rejected due to hostname verification failure from standard TLS (not a fingerprint error)

#### Scenario: Cert signed by untrusted CA with matching pinned fingerprint
- **WHEN** a client dials a TLS endpoint presenting a certificate signed by an untrusted CA, even if the SHA-256 fingerprint matches the pin
- **THEN** the connection MUST be rejected due to CA chain validation failure from standard TLS

#### Scenario: No pinned cert configured
- **WHEN** the `IONOS_PINNED_CERT` environment variable is empty or unset
- **THEN** the client SHALL use the default `http.Transport` with standard TLS verification and no custom `DialTLSContext`
