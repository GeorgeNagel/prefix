package prefix

import (
	"math/rand"
	"strings"
	"testing"
)

func BenchmarkLoopPrefixMatch(b *testing.B) {
	prefixesArray := generateRandomStrings(80000, 30)
	prefixesSlice := prefixesArray[:]

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		toCheckString := generateRandomString(30)
		matchByLoop(prefixesSlice, toCheckString)
	}
}

func BenchmarkPrefixTreeMatch(b *testing.B) {
	prefixesArray := generateRandomStrings(80000, 30)
	prefixesSlice := prefixesArray[:]

	prefixTree := BuildPrefixTree(prefixesSlice)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		toCheckString := generateRandomString(30)
		StringHasPrefixMatch(prefixTree, toCheckString)
	}
}

func matchByLoop(prefixes []string, stringToCheck string) bool {
	for i := 0; i < len(prefixes); i++ {
		prefix := prefixes[i]
		if strings.HasPrefix(stringToCheck, prefix) {
			return true
		}
	}
	return false
}

func generateRandomStrings(numStrings int, maxLength int) []string {
	strings := []string{}
	for i := 0; i < numStrings; i++ {
		strings = append(strings, generateRandomString(maxLength))
	}
	return strings
}

func generateRandomString(maxLength int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	length := rand.Intn(maxLength)
	bytes := make([]rune, length)
	for i := range bytes {
		bytes[i] = letters[rand.Intn(len(letters))]
	}
	return string(bytes)
}
