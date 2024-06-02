package assert

import "testing"

func Equal[E comparable](t *testing.T, got, want E) {
	t.Helper()

	if got != want {
		t.Errorf("got: %v; want: %v", got, want)
	}
}