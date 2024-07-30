package palindrome

import "testing"

func TestPalindroemChecker(t *testing.T) {
	s := "Hell1LL ?Eh"
	got := PalindromeChecker(s)
	want := true
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}

}
