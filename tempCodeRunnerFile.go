func palindromeChecker(x string) bool {
	r := len(x) - 1
	for i := range int(len(x) / 2) {
		if x[i] != x[r-i] {
			return false
		}
	}
	return true
}