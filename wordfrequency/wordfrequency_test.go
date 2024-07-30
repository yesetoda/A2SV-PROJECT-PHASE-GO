package wordfrequency

import "testing"

func TestWordFrequency(t *testing.T) {
	s := "hello world he1llo ? .Wor1lD"
	got := WordFreq(s)
	want := map[string]int{"hello": 2, "world": 2}
	if len(got) != len(want) {
		t.Errorf("got %v, want %v", got, want)
	}
	for k, v := range want {
		if v != got[k] {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}
