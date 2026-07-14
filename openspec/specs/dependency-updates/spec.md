# dependency-updates Specification

## Purpose
TBD - created by archiving change find-vulnerabilities-and-fix-for-sdk-go-cdn-repo. Update Purpose after archive.
## Requirements
### Requirement: Module MUST declare a supported Go version
The `go` directive in `go.mod` SHALL be set to `1.21` or higher. Go 1.18 is end-of-life and no longer receives security backports.

#### Scenario: go.mod specifies a maintained Go version
- **WHEN** `go.mod` is read
- **THEN** the `go` directive MUST be `1.21` or higher

### Requirement: Direct dependencies MUST be free of known CVEs
All packages listed under `require` in `go.mod` SHALL be at versions with no outstanding CVEs in the Go vulnerability database (`vuln.go.dev`). `govulncheck ./...` MUST exit 0.

#### Scenario: govulncheck reports no vulnerabilities
- **WHEN** `govulncheck ./...` is run against the module
- **THEN** the tool MUST exit with code 0 and report no known vulnerabilities

#### Scenario: golang.org/x/net is at a CVE-free version
- **WHEN** `go.sum` is inspected for the resolved version of `golang.org/x/net`
- **THEN** it MUST be `v0.23.0` or higher (first version resolving CVE-2023-44487 and CVE-2022-27664)

### Requirement: Transitive dependencies MUST be resolved to current safe versions
After `go mod tidy`, the indirect dependencies `google.golang.org/appengine` and `github.com/golang/protobuf` MUST be updated to their latest compatible versions or removed if no longer needed by the dependency graph.

#### Scenario: go.sum does not contain known-vulnerable dependency hashes
- **WHEN** `go.sum` is checked against the Go checksum database after `go mod tidy`
- **THEN** no entry SHALL correspond to a version with a known CVE in the Go vulnerability database

