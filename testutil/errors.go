package testutil

import (
	"reflect"
	"testing"
)

func EqualError(t *testing.T, err, wantErr error) bool {
	t.Helper()

	// Check error type
	if reflect.TypeOf(err) != reflect.TypeOf(wantErr) {
		t.Fatalf("Error: got = %T; want = %T", err, wantErr)
	}

	if err == nil || wantErr == nil {
		return err == nil && wantErr == nil
	}
	// Check error content
	return err.Error() == wantErr.Error()
}
