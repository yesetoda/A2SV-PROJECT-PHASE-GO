package Controllers

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func ReadAndTrimSpaces(message string, reader *bufio.Reader) (string, error) {
	fmt.Print(message)
	s, err := reader.ReadString('\n')
	if err != nil {
		return ReadAndTrimSpaces(message, reader)
	}
	s = strings.Trim(s, "\n")
	return strings.Trim(s, " "), nil
}

func ReadInteger(message string, reader *bufio.Reader, MaxVal int, MinVal int) (int, error) {
	fmt.Print(message)
	year, err := reader.ReadString('\n')
	if err != nil {
		return ReadInteger(message, reader, MaxVal, MinVal)
	}

	year = strings.Trim(year, "\n")
	year = strings.Trim(year, " ")
	if !IsNumber(year) {
		fmt.Println("the input contains non numeric value")
		return ReadInteger(message, reader, MaxVal, MinVal)
	}
	intYear, _ := strconv.Atoi(year)
	if intYear < MinVal || intYear > MaxVal {
		fmt.Println("value should be between", MinVal, "---", MaxVal)
		return ReadInteger(message, reader, MaxVal, MinVal)

	}
	return intYear, nil
}

func IsNumber(s string) bool {
	for _, i := range s {
		if !unicode.IsDigit(i) {
			return false
		}
	}
	return true
}

func IsAlpha(s string) bool {
	for _, i := range s {

		if !unicode.IsLetter(i) {
			return false
		}
	}
	return true
}

func ReadAlphabet(message string, reader *bufio.Reader) (string, error) {
	fmt.Print(message)
	letter, err := reader.ReadString('\n')
	if err != nil {
		return ReadAlphabet(message, reader)
	}
	letter = strings.Trim(letter, "\n")
	letter = strings.Trim(letter, " ")
	for !IsAlpha(letter) {
		fmt.Print("the input contains non alphabetic character")
		return ReadAlphabet(message, reader)
	}
	return letter, nil
}
