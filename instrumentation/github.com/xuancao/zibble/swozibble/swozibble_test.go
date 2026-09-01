package swozibble

import "testing"

func TestProcess(t *testing.T) {
	got := Process("hello")
	want := "OLLEH!"
	if got != want {
		t.Errorf("Process(%q) = %q, want %q", "hello", got, want)
	}
}
