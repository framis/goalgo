package substring

import "testing"

func TestSearch(t *testing.T) {
	text := "Hello world"
	pattern := "wo"
	rk := NewRabinKarp(pattern)
	if (rk.search(text) != 6) {
		t.Logf("Test failed")
		t.Fail()
	}
}