package main

import (
	"errors"
	"unicode"
)

const slashRune = 92

var ErrInvalidString = errors.New("invalid string")

func UnpackString(input string) (string, error) {
	var result string

	if len(input) == 0 {
		return "", nil
	}

	runes := []rune(input)
	if !validateString(runes) {
		return "", ErrInvalidString
	}

	slash := false
	for i := 0; i < len(runes); i++ {
		if runes[i] == slashRune {
			result += string(runes[i])
			slash = true
			continue
		}

		if unicode.IsDigit(runes[i]) {
			if runes[i] == '0' {
				result = result[:len(result)-1]
				continue
			}
			stringRepeat := ""
			for j := 0; j < int(runes[i]-'0')-1; j++ {
				if slash {
					stringRepeat += "\\" + string(runes[i-1])
					continue
				}
				stringRepeat += string(runes[i-1])
			}
			result += stringRepeat
			continue
		}

		result += string(runes[i])
	}

	return result, nil
}

func validateString(input []rune) bool {
	if len(input) == 1 && unicode.IsDigit(input[0]) {
		return false
	}

	for i := 0; i < len(input)-1; i++ {

		if unicode.IsDigit(input[i]) && unicode.IsDigit(input[i+1]) {
			return false
		}
		if unicode.IsDigit(input[0]) {
			return false
		}
	}

	return true
}
