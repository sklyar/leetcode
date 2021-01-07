package validate_ip_address

import (
	"strconv"
	"strings"
)

func validIPAddress(ip string) string {
	if isIPv4(ip) {
		return "IPv4"
	}

	if isIPv6(ip) {
		return "IPv6"
	}

	return "Neither"
}

func isIPv4(ip string) bool {
	octs := strings.Split(ip, ".")
	if len(octs) != 4 {
		return false
	}

	for _, oct := range octs {
		if len(oct) > 1 && oct[0] == '0' {
			return false
		}

		n, err := strconv.Atoi(oct)
		if err != nil {
			return false
		}

		if n < 0 || n > 255 {
			return false
		}
	}

	return true
}

func isIPv6(ip string) bool {
	octs := strings.Split(ip, ":")
	if len(octs) != 8 {
		return false
	}

	for _, oct := range octs {
		if len(oct) < 1 || len(oct) > 4 {
			return false
		}

		for _, ch := range oct {
			if !isDigit(ch) && !isValidLetter(ch) {
				return false
			}
		}
	}

	return true
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func isValidLetter(r rune) bool {
	return (r >= 'a' && r <= 'f') || (r >= 'A' && r <= 'F')
}
