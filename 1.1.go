package cryptopals

import (
	"encoding/base64"
	"encoding/hex"
)

func hexToBase64(hexString string) string {

	hexValue, _ := hex.DecodeString(hexString)

	return base64.RawStdEncoding.EncodeToString(hexValue)
}
