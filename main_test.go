package prefix

import (
	"reflect"
	"testing"
)

func TestStringIsPrefixOfItself(t *testing.T) {
	prefixes := []string{"abba"}
	toCheck := []string{"abba"}
	result := Match(prefixes, toCheck)
	if !stringSliceEq(result, toCheck) {
		t.Fail()
	}
}

func TestStringIsPrefixOfLargerString(t *testing.T) {
	prefixes := []string{"ab"}
	toCheck := []string{"abba"}
	result := Match(prefixes, toCheck)
	if !stringSliceEq(result, toCheck) {
		t.Fail()
	}
}

func TestStringIsNotPrefix(t *testing.T) {
	prefixes := []string{"ab"}
	toCheck := []string{"abba"}
	result := Match(prefixes, toCheck)
	if !stringSliceEq(result, []string{}) {
		t.Fail()
	}
}

func TestBuildPrefixTreeSimpleStrings(t *testing.T) {
	prefixes := []string{"ab", "po"}
	actualPrefixTree := buildPrefixTree(prefixes)
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

func TestStringHasPrefixMatch(t *testing.T) {
	prefixTree := map[interface{}]interface{}{
		byte('a'): map[interface{}]interface{}{
			byte('b'): map[interface{}]interface{}{
				nil: nil,
			},
		},
	}
	stringToCheck := "abc"
	hasMatch := stringHasPrefixMatch(prefixTree, stringToCheck)
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
	hasMatch := stringHasPrefixMatch(prefixTree, stringToCheck)
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
	hasMatch := stringHasPrefixMatch(prefixTree, stringToCheck)
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
