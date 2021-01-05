package maximum_69_number

import "strconv"

func maximum69Number(n int) int {
	s := []rune(strconv.Itoa(n))

	for i := range s {
		if s[i] == '6' {
			s[i] = '9'
			break
		}
	}

	v, err := strconv.Atoi(string(s))
	if err != nil {
		panic(err)
	}
	return v
}
