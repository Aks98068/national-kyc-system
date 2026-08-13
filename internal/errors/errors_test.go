package errors

import (
	"errors"
	"net/http"
	"testing"
)

func TestValidationError(t *testing.T) {
	err := Validation("invalid document")

	if err.Code != CodeValidation {
		t.Fatalf(
			"expected code %q, got %q",
			CodeValidation,
			err.Code,
		)
	}

	if err.HTTPStatus != http.StatusBadRequest {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusBadRequest,
			err.HTTPStatus,
		)
	}

	if err.Message != "invalid document" {
		t.Fatalf(
			"unexpected message: %q",
			err.Message,
		)
	}
}

func TestUnauthorizedError(t *testing.T) {
	err := Unauthorized()

	if err.Code != CodeUnauthorized {
		t.Fatalf(
			"expected code %q, got %q",
			CodeUnauthorized,
			err.Code,
		)
	}

	if err.HTTPStatus != http.StatusUnauthorized {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusUnauthorized,
			err.HTTPStatus,
		)
	}
}

func TestForbiddenError(t *testing.T) {
	err := Forbidden()

	if err.Code != CodeForbidden {
		t.Fatalf(
			"expected code %q, got %q",
			CodeForbidden,
			err.Code,
		)
	}

	if err.HTTPStatus != http.StatusForbidden {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusForbidden,
			err.HTTPStatus,
		)
	}
}

func TestInternalErrorDoesNotExposeCause(t *testing.T) {
	original := errors.New(
		"database password=super-secret internal connection failure",
	)

	err := Internal(original)

	if err.SafeMessage() != "An internal error occurred." {
		t.Fatalf(
			"unsafe message exposed: %q",
			err.SafeMessage(),
		)
	}

	if err.Cause == nil {
		t.Fatal("expected internal cause to be preserved")
	}
}

func TestDatabaseError(t *testing.T) {
	original := errors.New("connection refused")

	err := Database(original)

	if err.Code != CodeDatabase {
		t.Fatalf(
			"expected code %q, got %q",
			CodeDatabase,
			err.Code,
		)
	}

	if err.HTTPStatus != http.StatusInternalServerError {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusInternalServerError,
			err.HTTPStatus,
		)
	}

	if err.SafeMessage() == "connection refused" {
		t.Fatal("internal database error was exposed")
	}
}

func TestBlockchainError(t *testing.T) {
	original := errors.New(
		"endorsement failed: peer0.org1.example.com",
	)

	err := Blockchain(original)

	if err.Code != CodeBlockchain {
		t.Fatalf(
			"expected code %q, got %q",
			CodeBlockchain,
			err.Code,
		)
	}

	if err.HTTPStatus != http.StatusInternalServerError {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusInternalServerError,
			err.HTTPStatus,
		)
	}

	if err.SafeMessage() == original.Error() {
		t.Fatal("blockchain internal error was exposed")
	}
}
