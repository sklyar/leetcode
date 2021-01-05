package longest_common_prefix

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prop := strs[0]

	var prefix string
	for i := 0; i < len(prop); i++ {
		newPrefix := prefix + string(prop[i])
		newPrefixLength := len(newPrefix)

		match := true
		for j := 1; j < len(strs); j++ {
			currentStr := strs[j]
			if newPrefixLength > len(currentStr) || newPrefix != currentStr[0:newPrefixLength] {
				match = false
				break
			}
		}

		if !match {
			break
		}

		prefix = newPrefix
	}

	return prefix
}
