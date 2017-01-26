package prefix

import (
	"math/rand"
	"strings"
	"testing"
)

func BenchmarkTwoLoopsPrefixMatch(b *testing.B) {
	prefixesArray := generateRandomStrings(80000, 3)
	toCheckArray := generateRandomStrings(7000, 10)

	prefixesSlice := prefixesArray[:]
	toCheckSlice := toCheckArray[:]
	for i := 0; i < b.N; i++ {
		matchByTwoLoops(prefixesSlice, toCheckSlice)
	}
}

func BenchmarkPrefixTreeMatch(b *testing.B) {
	prefixesArray := generateRandomStrings(80000, 3)
	toCheckArray := generateRandomStrings(7000, 10)

	prefixesSlice := prefixesArray[:]
	toCheckSlice := toCheckArray[:]
	for i := 0; i < b.N; i++ {
		Match(prefixesSlice, toCheckSlice)
	}
}

func matchByTwoLoops(prefixes []string, stringsToCheck []string) []string {
	matchingStrings := []string{}
	for i := 0; i < len(stringsToCheck); i++ {
		for j := 0; j < len(prefixes); j++ {
			stringToCheck := stringsToCheck[i]
			prefix := prefixes[j]
			if strings.HasPrefix(stringToCheck, prefix) {
				matchingStrings = append(matchingStrings, stringToCheck)
			}
		}
	}
	return matchingStrings
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
