package main

import "strings"

func border(char string, repeatno int) string {
	// repeat chars and/or strings
	chars := strings.Repeat(char, repeatno) + "\n"
	return chars
}
