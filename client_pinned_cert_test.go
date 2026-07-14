package ionoscloud

import (
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func tlsTestServer(t *testing.T) *httptest.Server {
	t.Helper()
	srv := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	t.Cleanup(srv.Close)
	return srv
}

func serverLeafCert(srv *httptest.Server) *x509.Certificate {
	cert, _ := x509.ParseCertificate(srv.TLS.Certificates[0].Certificate[0])
	return cert
}

func certFingerprint(cert *x509.Certificate) string {
	fp := sha256.Sum256(cert.Raw)
	return hex.EncodeToString(fp[:])
}

func serverRootCAs(srv *httptest.Server) *x509.CertPool {
	return srv.Client().Transport.(*http.Transport).TLSClientConfig.RootCAs
}

// Scenario 1: valid CA-signed cert whose fingerprint matches → connection succeeds.
func TestPinnedCert_ValidCert_MatchingFingerprint(t *testing.T) {
	srv := tlsTestServer(t)
	leaf := serverLeafCert(srv)
	fingerprint := certFingerprint(leaf)

	tlsCfg := &tls.Config{RootCAs: serverRootCAs(srv)}
	transport := &http.Transport{}
	transport.DialTLSContext = addPinnedCertVerification([]byte(fingerprint), tlsCfg)

	client := &http.Client{Transport: transport}
	resp, err := client.Get(srv.URL)
	if err != nil {
		t.Fatalf("expected connection to succeed, got: %v", err)
	}
	resp.Body.Close()
}

// Scenario 2: valid CA-signed cert whose fingerprint does NOT match → fingerprint mismatch error.
func TestPinnedCert_ValidCert_WrongFingerprint(t *testing.T) {
	srv := tlsTestServer(t)
	wrongFingerprint := strings.Repeat("00", 32)

	tlsCfg := &tls.Config{RootCAs: serverRootCAs(srv)}
	transport := &http.Transport{}
	transport.DialTLSContext = addPinnedCertVerification([]byte(wrongFingerprint), tlsCfg)

	client := &http.Client{Transport: transport}
	_, err := client.Get(srv.URL)
	if err == nil {
		t.Fatal("expected fingerprint mismatch error, but connection succeeded")
	}
	if !strings.Contains(err.Error(), "fingerprint") {
		t.Errorf("expected fingerprint error, got: %v", err)
	}
}

// Scenario 3: correct fingerprint but ServerName forced to a hostname not in the cert SANs →
// TLS hostname verification failure (not a fingerprint error).
func TestPinnedCert_ValidFingerprint_HostnameMismatch(t *testing.T) {
	srv := tlsTestServer(t)
	leaf := serverLeafCert(srv)
	fingerprint := certFingerprint(leaf)

	// ServerName does not match the cert's SANs (127.0.0.1, ::1).
	tlsCfg := &tls.Config{
		RootCAs:    serverRootCAs(srv),
		ServerName: "invalid.hostname.that.does.not.match",
	}
	transport := &http.Transport{}
	transport.DialTLSContext = addPinnedCertVerification([]byte(fingerprint), tlsCfg)

	client := &http.Client{Transport: transport}
	_, err := client.Get(srv.URL)
	if err == nil {
		t.Fatal("expected hostname verification failure, but connection succeeded")
	}
	if strings.Contains(err.Error(), "fingerprint") {
		t.Errorf("expected TLS hostname error, not fingerprint error, got: %v", err)
	}
}

// Scenario 4: cert signed by untrusted CA with correct fingerprint →
// CA chain validation failure (not a fingerprint error).
func TestPinnedCert_UntrustedCA_MatchingFingerprint(t *testing.T) {
	srv := tlsTestServer(t)
	leaf := serverLeafCert(srv)
	fingerprint := certFingerprint(leaf)

	// Empty CertPool: server's CA is not trusted.
	tlsCfg := &tls.Config{RootCAs: x509.NewCertPool()}
	transport := &http.Transport{}
	transport.DialTLSContext = addPinnedCertVerification([]byte(fingerprint), tlsCfg)

	client := &http.Client{Transport: transport}
	_, err := client.Get(srv.URL)
	if err == nil {
		t.Fatal("expected CA chain validation failure, but connection succeeded")
	}
	if strings.Contains(err.Error(), "fingerprint") {
		t.Errorf("expected CA chain error, not fingerprint error, got: %v", err)
	}
}

// Scenario 5: no fingerprint configured → transport.DialTLSContext remains nil.
func TestPinnedCert_NoFingerprintConfigured(t *testing.T) {
	transport := &http.Transport{}
	AddPinnedCert(transport, "")
	if transport.DialTLSContext != nil {
		t.Fatal("expected DialTLSContext to be nil when no fingerprint is configured")
	}
}
