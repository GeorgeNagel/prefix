package prefix

func Match(prefixes []string, toCheck []string) []string {
	// prefixTree := make(map[byte]byte)
	return nil
}

func buildPrefixTree(prefixes []string) map[interface{}]interface{} {
	tree := make(map[interface{}]interface{})
	for _, str := range prefixes {
		currentNode := tree
		for i := 0; i < len(str); i++ {
			b := str[i]
			currentNode[b] = make(map[interface{}]interface{})
			nextNode := currentNode[b].(map[interface{}]interface{})
			currentNode = nextNode
		}
		currentNode[nil] = nil
	}
	return tree
}

func stringHasPrefixMatch(prefixTree map[interface{}]interface{}, stringToCheck string) bool {
	return false
}
