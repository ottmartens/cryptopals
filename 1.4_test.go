package cryptopals

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestFindXORPhrase(t *testing.T) {

	fileContents, _ := ioutil.ReadFile("1.4.data.txt")

	lines := strings.Split(string(fileContents), "\n")

	phrase := findXORPhrase(lines)

	assertEquals(t, phrase, "Now that the party is jumping\n")
}
