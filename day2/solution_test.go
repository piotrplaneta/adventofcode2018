package day2

import "testing"

func TestChecksum0Strings(t *testing.T) {
	input := make([]string, 0)
	checksum := Checksum(input)
	if checksum != 0 {
		t.Errorf("Checksum was incorrect, got: %d, want: %d.", checksum, 0)
	}
}

func TestChecksum1StringNoMatches(t *testing.T) {
	input := []string{"abcd"}
	checksum := Checksum(input)
	if checksum != 0 {
		t.Errorf("Checksum was incorrect, got: %d, want: %d.", checksum, 0)
	}
}

func TestChecksum1StringBothMatches(t *testing.T) {
	input := []string{"ababa"}
	checksum := Checksum(input)
	if checksum != 1 {
		t.Errorf("Checksum was incorrect, got: %d, want: %d.", checksum, 1)
	}
}

func TestChecksumMultipleStrings(t *testing.T) {
	input := []string{"abcdef", "bababc", "abbcde", "abcccd", "aabcdd", "abcdee", "ababab"}
	checksum := Checksum(input)
	if checksum != 12 {
		t.Errorf("Checksum was incorrect, got: %d, want: %d.", checksum, 12)
	}
}

func TestSimilarId(t *testing.T) {
	input := []string{"abcde", "fghij", "klmno", "pqrst", "fguij", "axcye", "wvxyz"}
	similarID := SimilarID(input)
	if similarID != "fgij" {
		t.Errorf("Similar id was incorrect, got: %s, want: %s.", similarID, "fgij")
	}
}
