package prefix

import (
	"reflect"
	"testing"
)

func TestStringIsPrefixOfItself(t *testing.T) {
	prefixes := []string{"abba"}
	toCheck := []string{"abba"}
	prefixTree := BuildPrefixTree(prefixes)
	result := Match(prefixTree, toCheck)
	if !stringSliceEq(result, toCheck) {
		t.Fatalf("%s not in %s", toCheck, result)
	}
}

func TestStringIsPrefixOfLargerString(t *testing.T) {
	prefixes := []string{"ab"}
	toCheck := []string{"abba"}
	prefixTree := BuildPrefixTree(prefixes)
	result := Match(prefixTree, toCheck)
	if !stringSliceEq(result, toCheck) {
		t.Fail()
	}
}

func TestStringIsNotPrefix(t *testing.T) {
	prefixes := []string{"ab"}
	toCheck := []string{"academia"}
	prefixTree := BuildPrefixTree(prefixes)
	result := Match(prefixTree, toCheck)
	if !stringSliceEq(result, []string{}) {
		t.Fatalf("Found invalid matches: %s", result)
	}
}

func TestEmptyStringIsPrefixOfEverything(t *testing.T) {
	prefixes := []string{""}
	toCheck := []string{"academia", "sandstorm"}
	prefixTree := BuildPrefixTree(prefixes)
	result := Match(prefixTree, toCheck)
	if !stringSliceEq(result, toCheck) {
		t.Fatalf("%s not in %s", toCheck, result)
	}
}

func TestEmptyStringDoesntMatchPrefixes(t *testing.T) {
	prefixes := []string{"foot", "bart"}
	toCheck := []string{""}
	prefixTree := BuildPrefixTree(prefixes)
	result := Match(prefixTree, toCheck)
	if !stringSliceEq(result, []string{}) {
		t.Fatalf("Found invalid matches: %s", result)
	}
}

func TestBuildPrefixTreeSimpleStrings(t *testing.T) {
	prefixes := []string{"ab", "po"}
	actualPrefixTree := BuildPrefixTree(prefixes)
	expectedPrefixTree := map[interface{}]interface{}{
		byte('a'): map[interface{}]interface{}{
			byte('b'): map[interface{}]interface{}{
				nil: nil,
			},
		},
		byte('p'): map[interface{}]interface{}{
			byte('o'): map[interface{}]interface{}{
				nil: nil,
			},
		},
	}
	eq := reflect.DeepEqual(actualPrefixTree, expectedPrefixTree)
	if !eq {
		t.Fatalf("%s != %s", actualPrefixTree, expectedPrefixTree)
	}
}

func TestBuildPrefixTreeSameStart(t *testing.T) {
	prefixes := []string{"ab", "ac"}
	actualPrefixTree := BuildPrefixTree(prefixes)
	expectedPrefixTree := map[interface{}]interface{}{
		byte('a'): map[interface{}]interface{}{
			byte('b'): map[interface{}]interface{}{
				nil: nil,
			},
			byte('c'): map[interface{}]interface{}{
				nil: nil,
			},
		},
	}
	eq := reflect.DeepEqual(actualPrefixTree, expectedPrefixTree)
	if !eq {
		t.Fatalf("%s != %s", actualPrefixTree, expectedPrefixTree)
	}
}

func TestStringHasPrefixMatch(t *testing.T) {
	prefixTree := map[interface{}]interface{}{
		byte('a'): map[interface{}]interface{}{
			byte('b'): map[interface{}]interface{}{
				nil: nil,
			},
		},
	}
	stringToCheck := "abc"
	hasMatch := StringHasPrefixMatch(prefixTree, stringToCheck)
	if !hasMatch {
		t.Fail()
	}
}

func TestStringDoesNotHavePrefixMatch(t *testing.T) {
	prefixTree := map[interface{}]interface{}{
		byte('a'): map[interface{}]interface{}{
			byte('b'): map[interface{}]interface{}{
				nil: nil,
			},
		},
	}
	stringToCheck := "ac"
	hasMatch := StringHasPrefixMatch(prefixTree, stringToCheck)
	if hasMatch {
		t.Fail()
	}
}

func TestStringDoesNotHavePrefixMatchTooShort(t *testing.T) {
	prefixTree := map[interface{}]interface{}{
		byte('a'): map[interface{}]interface{}{
			byte('b'): map[interface{}]interface{}{
				nil: nil,
			},
		},
	}
	stringToCheck := "a"
	hasMatch := StringHasPrefixMatch(prefixTree, stringToCheck)
	if hasMatch {
		t.Fail()
	}
}

func stringSliceEq(slice1 []string, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}
