package httpserver

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/national-kyc-system/national-kyc-system/internal/logger"
)

func TestLivenessEndpoint(t *testing.T) {
	router := NewRouter()

	req := httptest.NewRequest(
		http.MethodGet,
		"/health/live",
		nil,
	)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusOK,
			recorder.Code,
		)
	}

	if contentType := recorder.Header().Get("Content-Type"); contentType != "application/json" {
		t.Fatalf(
			"expected JSON content type, got %q",
			contentType,
		)
	}

	expected := `{"status":"ok"}`

	if !strings.Contains(
		recorder.Body.String(),
		expected,
	) {
		t.Fatalf(
			"expected response to contain %q, got %q",
			expected,
			recorder.Body.String(),
		)
	}
}

func TestReadinessEndpoint(t *testing.T) {
	router := NewRouter()

	req := httptest.NewRequest(
		http.MethodGet,
		"/health/ready",
		nil,
	)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusOK,
			recorder.Code,
		)
	}

	if !strings.Contains(
		recorder.Body.String(),
		`{"status":"ready"}`,
	) {
		t.Fatalf(
			"unexpected readiness response: %q",
			recorder.Body.String(),
		)
	}
}

func TestRequestIDGenerated(t *testing.T) {
	router := NewRouter()

	handler := requestIDMiddleware(router)

	req := httptest.NewRequest(
		http.MethodGet,
		"/health/live",
		nil,
	)

	recorder := httptest.NewRecorder()

	handler.ServeHTTP(recorder, req)

	requestID := recorder.Header().Get(
		requestIDHeader,
	)

	if requestID == "" {
		t.Fatal("expected request ID to be generated")
	}

	if len(requestID) != 32 {
		t.Fatalf(
			"expected 32-character request ID, got %d",
			len(requestID),
		)
	}
}

func TestRequestIDPreserved(t *testing.T) {
	router := NewRouter()

	handler := requestIDMiddleware(router)

	req := httptest.NewRequest(
		http.MethodGet,
		"/health/live",
		nil,
	)

	req.Header.Set(
		requestIDHeader,
		"test-request-id",
	)

	recorder := httptest.NewRecorder()

	handler.ServeHTTP(recorder, req)

	got := recorder.Header().Get(
		requestIDHeader,
	)

	if got != "test-request-id" {
		t.Fatalf(
			"expected request ID %q, got %q",
			"test-request-id",
			got,
		)
	}
}

func TestSecurityHeaders(t *testing.T) {
	handler := securityHeadersMiddleware(
		http.HandlerFunc(
			func(
				w http.ResponseWriter,
				r *http.Request,
			) {
				w.WriteHeader(http.StatusOK)
			},
		),
	)

	req := httptest.NewRequest(
		http.MethodGet,
		"/",
		nil,
	)

	recorder := httptest.NewRecorder()

	handler.ServeHTTP(recorder, req)

	tests := map[string]string{
		"X-Content-Type-Options": "nosniff",
		"X-Frame-Options":        "DENY",
		"Referrer-Policy":        "no-referrer",
		"Cache-Control":          "no-store",
		"Permissions-Policy":     "camera=(), microphone=(), geolocation=()",
	}

	for header, expected := range tests {
		got := recorder.Header().Get(header)

		if got != expected {
			t.Fatalf(
				"header %s: expected %q, got %q",
				header,
				expected,
				got,
			)
		}
	}
}

func TestRecoveryMiddleware(t *testing.T) {
	log := logger.New(
		logger.Config{
			Level:  "error",
			Format: "json",
		},
	)

	panicHandler := http.HandlerFunc(
		func(
			w http.ResponseWriter,
			r *http.Request,
		) {
			panic("test panic")
		},
	)

	handler := requestIDMiddleware(
		recoveryMiddleware(
			log,
			panicHandler,
		),
	)

	req := httptest.NewRequest(
		http.MethodGet,
		"/test",
		nil,
	)

	recorder := httptest.NewRecorder()

	handler.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusInternalServerError {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusInternalServerError,
			recorder.Code,
		)
	}

	body := recorder.Body.String()

	if strings.Contains(body, "test panic") {
		t.Fatal("panic details were exposed to client")
	}

	if !strings.Contains(
		body,
		"INTERNAL_ERROR",
	) {
		t.Fatal("expected INTERNAL_ERROR response")
	}
}

func TestGenerateRequestID(t *testing.T) {
	first := generateRequestID()
	second := generateRequestID()

	if first == "" || second == "" {
		t.Fatal("request ID must not be empty")
	}

	if first == second {
		t.Fatal("request IDs should be unique")
	}

	if len(first) != 32 {
		t.Fatalf(
			"expected request ID length 32, got %d",
			len(first),
		)
	}
}
