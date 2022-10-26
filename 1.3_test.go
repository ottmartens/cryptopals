package cryptopals

import "testing"

func TestDecipherSingleXOR(t *testing.T) {

	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	want := "Cooking MC's like a pound of bacon"

	result, _ := decipherSingleXOR(input)

	assertEquals(t, result, want)
}
