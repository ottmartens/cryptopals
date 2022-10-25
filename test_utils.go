package cryptopals

import "testing"

func assertEquals(t *testing.T, result, want interface{}) {
	if result != want {
		t.Fatalf("Got '%v', expected '%v'", result, want)
	}
}
