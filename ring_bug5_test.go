package main

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

// TestWriteErrorHandlesWrappedStatusErr verifies that writeError properly
// extracts the HTTP status code from a wrapped *statusErr using errors.As,
// instead of losing the code when the error is wrapped with fmt.Errorf.
func TestWriteErrorHandlesWrappedStatusErr(t *testing.T) {
	// Simulate a wrapped error (as might happen during refactoring)
	original := badRequest("something went wrong")
	wrapped := fmt.Errorf("context layer: %w", original)

	recorder := httptest.NewRecorder()
	writeError(recorder, wrapped)

	if recorder.Code != 400 {
		t.Fatalf("writeError should extract 400 from wrapped statusErr, got %d", recorder.Code)
	}
}
