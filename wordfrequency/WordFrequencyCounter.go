package wordfrequency

import (
	"strings"
	"unicode"
)

func WordFreq(s string) map[string]int {
	// var s string
	// scanner := bufio.NewReader(os.Stdin)
	// s, err := scanner.ReadString('\n')
	// if err != nil {
	// 	return make(map[string]int)
	// }
	freq := make(map[string]int)
	s = strings.ToLower(s)
	for _, i := range strings.Fields(s) {
		word := ""
		for _, j := range strings.Trim(i, "\n") {
			val := string(rune(j))
			if unicode.IsLetter(j) {
				word += val
			}
		}
		if len(word) > 0 {

			freq[word] += 1
		}
	}
	return freq
}
