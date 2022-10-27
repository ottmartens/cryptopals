package cryptopals

import (
	"encoding/hex"
)

func encryptWithRepeatingXOR(text string, key string) string {
	buffer := []byte(text)
	encrypted := make([]byte, len(buffer))

	keyIndex := 0
	for i, textByte := range buffer {
		encrypted[i] = textByte ^ key[keyIndex]

		if keyIndex < len(key)-1 {
			keyIndex++
		} else {
			keyIndex = 0
		}
	}
	return hex.EncodeToString(encrypted)
}
