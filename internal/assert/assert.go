package assert

import (
	"strings"
	"testing"
)

func Equal[E comparable](t *testing.T, got, want E) {
	t.Helper()

	if got != want {
		t.Errorf("got: %v; want: %v", got, want)
	}
}

func StringContains(t *testing.T, got, want string) {
	t.Helper()

	if !strings.Contains(got, want) {
		t.Errorf("got: %v, expected to contain: %v", got, want)
	}
}

func NilError(t *testing.T, got error) {
	t.Helper()

	if got != nil {
		t.Errorf("got: %v; want: nil", got)
	}
}
