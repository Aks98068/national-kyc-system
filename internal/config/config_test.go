package config

import (
	"testing"
)

func TestLoadDefaults(t *testing.T) {
	t.Setenv("APP_NAME", "")
	t.Setenv("APP_ENV", "")
	t.Setenv("HTTP_PORT", "")
	t.Setenv("HTTP_MAX_BODY_BYTES", "")
	t.Setenv("DATABASE_PORT", "")
	t.Setenv("REDIS_PORT", "")
	t.Setenv("SECURITY_ENABLE_TLS", "")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load() returned error: %v", err)
	}

	if cfg.App.Name != "national-kyc-system" {
		t.Fatalf(
			"expected default app name, got %q",
			cfg.App.Name,
		)
	}

	if cfg.HTTP.Port != 8080 {
		t.Fatalf(
			"expected default HTTP port 8080, got %d",
			cfg.HTTP.Port,
		)
	}

	if cfg.HTTP.MaxBodyBytes != 1048576 {
		t.Fatalf(
			"expected default max body size 1048576, got %d",
			cfg.HTTP.MaxBodyBytes,
		)
	}
}

func TestInvalidHTTPPort(t *testing.T) {
	t.Setenv("HTTP_PORT", "invalid")

	_, err := Load()
	if err == nil {
		t.Fatal("expected error for invalid HTTP_PORT")
	}
}

func TestTLSRequiresCertificate(t *testing.T) {
	t.Setenv("SECURITY_ENABLE_TLS", "true")
	t.Setenv("SECURITY_TLS_CERTIFICATE_PATH", "")
	t.Setenv("SECURITY_TLS_PRIVATE_KEY_PATH", "")

	_, err := Load()
	if err == nil {
		t.Fatal("expected TLS configuration validation error")
	}
}
