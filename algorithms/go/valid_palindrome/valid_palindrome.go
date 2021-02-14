package valid_palindrome

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		for !isChar(rune(s[i])) && i < j {
			i++
		}

		for !isChar(rune(s[j])) && i < j {
			j--
		}

		if strings.ToLower(string(s[i])) != strings.ToLower(string(s[j])) {
			return false
		}

	}

	return true
}

func isChar(ch rune) bool {
	return unicode.IsLetter(ch) || unicode.IsNumber(ch)
}
