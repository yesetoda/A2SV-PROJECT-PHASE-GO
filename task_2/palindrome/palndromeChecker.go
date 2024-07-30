package palindrome

import (
	"strings"
	"unicode"
)

func PalindromeChecker(x string) bool {

	s := strings.Trim(x, "\n")
	s = strings.ToLower(s)
	y := ""
	for _, i := range strings.Fields(s) {
		for _, j := range i {
			if unicode.IsLetter(j) || unicode.IsNumber(j) {
				y += string(j)
			}
		}
	}
	r := len(y) - 1
	for i := range int(len(y) / 2) {
		if y[i] != y[r-i] {
			return false
		}
	}
	return true
}
