package math_test

import (
	"gh_test/math"
	"testing"
)

func TestAdd(t *testing.T) {

	got := math.Add(4, 6)
	want := 11

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
