package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func palindromeChecker(x string) bool {
	r := len(x) - 1
	for i := range int(len(x) / 2) {
		if x[i] != x[r-i] {
			return false
		}
	}
	return true
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	s, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	s = strings.Trim(s, "\n")
	s = strings.ToLower(s)
	x := ""
	for _, i := range strings.Fields(s) {
		for _, j := range i {
			if unicode.IsLetter(j) || unicode.IsNumber(j) {
				x += string(j)
			}
		}
	}
	if palindromeChecker(x) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
