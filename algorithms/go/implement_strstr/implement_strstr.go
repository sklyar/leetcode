package implement_strstr

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	index := -1
	for i := 0; i < len(haystack); i++ {
		left := len(haystack) - i
		if len(needle) > left {
			break
		}

		if haystack[i] != needle[0] {
			continue
		}

		match := true
		for j := 1; j < len(needle); j++ {
			if needle[j] != haystack[i+j] {
				match = false
				break
			}
		}

		if match {
			index = i
			break
		}

	}

	return index
}
