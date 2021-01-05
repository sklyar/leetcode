package longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
	m := make(map[rune]int)
	max := 0

	for i, j := 0, 0; j < len(s); j++ {
		r := rune(s[j])
		if v, ok := m[r]; ok && v > i {
			i = v
		}

		m[r] = j + 1
		cur := j - i + 1
		if cur > max {
			max = cur
		}
	}

	return max
}
