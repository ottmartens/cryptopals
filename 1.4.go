package cryptopals

import (
	"math"
)

func findXORPhrase(phrases []string) string {
	var bestScorePhrase string
	bestScore := math.MinInt

	for _, phrase := range phrases {
		decoded, score := decipherSingleXOR(phrase)

		if score > bestScore {
			bestScore = score
			bestScorePhrase = decoded
		}
	}

	return bestScorePhrase
}
