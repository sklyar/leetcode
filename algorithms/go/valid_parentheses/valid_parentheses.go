package valid_parentheses

func isValid(s string) bool {
	stack := make([]rune, 0, len(s)/2)
	stackSize := 0

	invertBrackets := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, r := range s {
		if v, ok := invertBrackets[r]; ok {
			if stackSize == 0 {
				return false
			}

			top := stack[stackSize-1]
			if v != top {
				return false
			}

			stackSize--
			stack = stack[:stackSize]
		} else {
			stackSize++
			stack = append(stack, r)
		}
	}

	return stackSize == 0
}
