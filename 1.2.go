package cryptopals

import "encoding/hex"

func fixedXOR(a, b string) string {
	buf1, _ := hex.DecodeString(a)
	buf2, _ := hex.DecodeString(b)

	res := make([]byte, len(buf1))

	for idx := range buf1 {
		res[idx] = buf1[idx] ^ buf2[idx]
	}

	return hex.EncodeToString(res)
}
