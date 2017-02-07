package prefix

func Match(prefixTree map[interface{}]interface{}, toCheck []string) []string {
	matchingStrings := []string{}
	for _, str := range toCheck {
		hasPrefixMatch := StringHasPrefixMatch(prefixTree, str)
		if hasPrefixMatch {
			matchingStrings = append(matchingStrings, str)
		}
	}
	return matchingStrings
}

func BuildPrefixTree(prefixes []string) map[interface{}]interface{} {
	tree := make(map[interface{}]interface{})
	for _, str := range prefixes {
		currentNode := tree
		for i := 0; i < len(str); i++ {
			b := str[i]

			childMap, childMapExists := currentNode[b]
			if !childMapExists {
				childMap = make(map[interface{}]interface{})
				currentNode[b] = childMap
			}
			nextNode := childMap.(map[interface{}]interface{})

			currentNode = nextNode
		}
		currentNode[nil] = nil
	}
	return tree
}

func StringHasPrefixMatch(prefixTree map[interface{}]interface{}, stringToCheck string) bool {
	currentNode := prefixTree
	for i := 0; i < len(stringToCheck); i++ {
		b := stringToCheck[i]
		_, nilAtThisLevel := currentNode[nil]
		if nilAtThisLevel {
			return true
		}

		_, byteAtThisLevel := currentNode[b]
		if !byteAtThisLevel {
			return false
		} else {
			currentNode = currentNode[b].(map[interface{}]interface{})
		}
	}

	_, nilAtThisLevel := currentNode[nil]
	if nilAtThisLevel {
		return true
	} else {
		return false
	}
}
