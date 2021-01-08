package defanging_an_ip_address

func defangIPaddr(address string) string {
	res := make([]byte, 0, len(address)+6)

	for _, b := range []byte(address) {
		if b == '.' {
			res = append(res, '[', '.', ']')
		} else {
			res = append(res, b)
		}
	}

	return string(res)
}
