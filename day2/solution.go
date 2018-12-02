package day2

import (
	"github.com/piotrplaneta/adventofcode2018/utils"
)

type letterFreqencies map[rune]int
type pair struct {
	s1 string
	s2 string
}

//SolvePart1 returns the answer for part 1 of day 2
func SolvePart1() int {
	return Checksum(adventInput())
}

//SolvePart2 returns the answer for part 1 of day 2
func SolvePart2() string {
	return SimilarID(adventInput())
}

//Checksum returns the product of words with 2 and 3 same letters
func Checksum(strings []string) int {
	listOfFrequencies := make([]letterFreqencies, len(strings))

	for i, s := range strings {
		listOfFrequencies[i] = frequencyMap(s)
	}

	return wordsWithLetterOfFrequency(listOfFrequencies, 2) * wordsWithLetterOfFrequency(listOfFrequencies, 3)
}

//SimilarID returns the same letters in similar id
func SimilarID(strings []string) string {
	combinations := generate2ElemCombinations(strings)

	for combination := range combinations {
		if similar(combination) {
			return sameLetters(combination)
		}
	}

	return ""
}

func wordsWithLetterOfFrequency(listOfFrequencies []letterFreqencies, frequency int) int {
	result := 0

	for _, v := range listOfFrequencies {
		if containsFrequence(v, frequency) {
			result++
		}
	}

	return result
}

func containsFrequence(m letterFreqencies, frequency int) bool {
	for _, v := range m {
		if v == frequency {
			return true
		}
	}

	return false
}

func frequencyMap(word string) map[rune]int {
	frequencies := make(letterFreqencies)

	for _, v := range word {
		frequencies[v]++
	}

	return frequencies
}

func adventInput() []string {
	lines, _ := utils.ReadLinesFromFile("day2/input")
	return lines
}

func similar(combination pair) bool {
	errors := 0

	for i, v := range combination.s1 {
		if byte(v) != combination.s2[i] {
			if errors == 1 {
				return false
			}

			errors++
		}
	}

	return true
}

func sameLetters(combination pair) string {
	result := ""

	for i, v := range combination.s1 {
		if byte(v) == combination.s2[i] {
			result += string(v)
		}
	}

	return result
}

func generate2ElemCombinations(list []string) <-chan pair {
	c := make(chan pair)

	go func(c chan pair) {
		defer close(c)

		for i, s1 := range list {
			for _, s2 := range list[i+1:] {
				c <- pair{s1: s1, s2: s2}
			}
		}
	}(c)

	return c
}
