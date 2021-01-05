package reverse_integer

import "math"

func reverse(a int) int {
	var b int
	for a != 0 {
		pop := a % 10
		a /= 10

		if (b > math.MaxInt32/10 || b == math.MaxInt32/10 && pop > 7) ||
			(b < math.MinInt32/10 || b == -math.MinInt32/10 && pop > 8) {
			return 0
		}

		b = b*10 + pop
	}

	return b
}
