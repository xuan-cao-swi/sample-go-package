package zibble

import "testing"

func TestReverse(t *testing.T) {
	got := Reverse("hello")
	want := "olleh"
	if got != want {
		t.Errorf("Reverse(%q) = %q, want %q", "hello", got, want)
	}
}

func TestIsPalindrome(t *testing.T) {
	cases := map[string]bool{
		"racecar": true,
		"Level":   true,
		"golang":  false,
	}
	for input, want := range cases {
		if got := IsPalindrome(input); got != want {
			t.Errorf("IsPalindrome(%q) = %v, want %v", input, got, want)
		}
	}
}

func TestShout(t *testing.T) {
	got := Shout("hi")
	want := "HI!"
	if got != want {
		t.Errorf("Shout(%q) = %q, want %q", "hi", got, want)
	}
}
