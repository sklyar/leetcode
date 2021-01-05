package palindrome_number

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x < 2 {
		return true
	}

	original := x

	var reversed int
	for x != 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}

	return reversed == original
}
