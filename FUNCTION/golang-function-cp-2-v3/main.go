package main

import (
	"fmt"
	"strings"
)

func CountVowelConsonant(str string) (int, int, bool) {
	vowels := "aiueoAIUEO"
	vowelCount := 0
	consonantCount := 0

	for _, char := range str {
		if strings.ContainsRune(vowels, char) {
			vowelCount++
		} else if char != ' ' && ((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')) {
			consonantCount++
		}
	}

	noVowelOrConsonant := vowelCount == 0 || consonantCount == 0
	return vowelCount, consonantCount, noVowelOrConsonant
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountVowelConsonant("Hidup Itu Indah"))
}
