package main

import "fmt"

import "strings"
import "encoding/hex"
import "crypto/md5"

const dayFourInput = "bgvyzdsv"

func hashMatch(hash []byte, prefix string) bool {
	base16 := hex.EncodeToString(hash)
	return strings.HasPrefix(base16, prefix)
}

func findPrefix(prefix string) int {
	for i := 0; i < 10000000; i++ {
		testInput := fmt.Sprintf("%v%v", dayFourInput, i)
		hash := md5.Sum([]byte(testInput))

		if hashMatch(hash[:], prefix) {
			return i
		}
	}

	return -1
}

// AdventCoin1 returns the lowest number which when joined with the input,
// creates a hash prefixed with 00000
func AdventCoin1() int {
	return findPrefix("00000")
}

func AdventCoin2() int {
	return findPrefix("000000")
}
