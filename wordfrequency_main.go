package main

import (
	"fmt"
	"test/A2SV-PROJECT-PHASE-GO/wordfrequency"
)

func main() {
	s := "hello world he1llo ? .Wor1lD"
	ans := wordfrequency.WordFreq(s)
	fmt.Println(ans)
}
