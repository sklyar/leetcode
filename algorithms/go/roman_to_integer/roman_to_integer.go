package roman_to_integer

func romanToInt(s string) int {
	r2i := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	size := len(s)
	runes := []rune(s)
	result := 0

	for i := 0; i < size; i++ {
		current := r2i[runes[i]]
		if i != size-1 {
			next := r2i[runes[i+1]]
			if current < next {
				current = next - current
				i++
			}
		}

		result += current
	}

	return result
}
