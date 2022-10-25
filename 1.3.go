package cryptopals

import (
	"encoding/hex"
	"fmt"
	"math"
	"strings"
)

var letterRunes []rune = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func decipherSingleXOR(cyphered string) string {

	buffer, _ := hex.DecodeString(cyphered)

	results := make(map[rune]string)

	smallestScore := math.MaxInt
	bestChar := 'a'

	for _, char := range letterRunes {

		charResult := make([]byte, len(buffer))

		for idx := range buffer {
			charResult[idx] = buffer[idx] ^ byte(char)
		}

		results[char] = string(charResult)
		score := scoreAnswer(string(charResult))

		if score < smallestScore {
			smallestScore = score
			bestChar = char
		}
	}

	fmt.Printf("Best char for decypher is %c", bestChar)

	return results[bestChar]

}

func scoreAnswer(answer string) int {

	words := strings.Split(answer, " ")

	wordLengthSums := 0
	for _, word := range words {
		wordLengthSums += len(word)
	}
	averageWordLength := wordLengthSums / len(words)

	wordLengthScore := 5 - averageWordLength
	if wordLengthScore < 0 {
		wordLengthScore = -wordLengthScore
	}

	letterCharsCount := 0
	for _, char := range answer {
		for _, letterRune := range letterRunes {
			if letterRune == char {
				letterCharsCount++
				break
			}
		}
	}
	letterCountScore := len(answer) - letterCharsCount

	return wordLengthScore * letterCountScore
}
