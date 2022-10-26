package cryptopals

import (
	"encoding/hex"
	"math"
	"strings"
)

var letterFrequencies = map[rune]float64{
	'e': 12,
	'a': 8,
	'i': 8,
	'o': 8,
	'u': 7.4,
	's': 7,
	'h': 6.4,
	'r': 6.2,
	't': 5,
	'n': 5,
	'd': 4.4,
	'l': 4,
	'c': 3.3,
	'm': 3.3,
	'f': 2.5,
	'g': 1.7,
	'p': 1.7,
	'b': 1.6,
	'v': 1.2,
	'k': 0.8,
	'q': 0.5,
	'j': 0.4,
	'x': 0.4,
	'z': 0.2,
}

func decipherSingleXOR(cyphered string) (bestResult string, score int) {
	buffer, _ := hex.DecodeString(cyphered)

	largestScore := math.MinInt

	for i := 0; i < 256; i++ {
		char := rune(i)

		resultBuffer := make([]byte, len(buffer))

		for idx := range buffer {
			resultBuffer[idx] = buffer[idx] ^ byte(char)
		}

		result := string(resultBuffer)

		if score := scoreAnswer(result); score > largestScore {
			largestScore = score
			bestResult = result
		}
	}

	return bestResult, largestScore
}

func scoreAnswer(answer string) int {
	answer = strings.ToLower(answer)

	charScoreSum := 0.0
	for _, char := range answer {
		charScoreSum += letterFrequencies[char]
	}

	words := strings.Split(answer, " ")
	wordCountScore := len(words)

	return int(charScoreSum) * wordCountScore
}
