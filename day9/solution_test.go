package day9

import "testing"

func TestWinnerResultTrivial(t *testing.T) {
	result := WinnerResult(9, 25)
	if result != 32 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 32)
	}
}

func TestWinnerResult(t *testing.T) {
	result := WinnerResult(10, 1618)
	if result != 8317 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 8317)
	}
}
