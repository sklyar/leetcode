package fizz_buzz

import "strconv"

func fizzBuzz(n int) []string {
	res := make([]string, n)
	for i := 0; i < n; i++ {
		res[i] = n2str(i + 1)
	}

	return res
}

func n2str(v int) string {
	if v%3 == 0 && v%5 == 0 {
		return "FizzBuzz"
	}

	if v%3 == 0 {
		return "Fizz"
	}

	if v%5 == 0 {
		return "Buzz"
	}

	return strconv.Itoa(v)
}
