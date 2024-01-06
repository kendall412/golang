package main

import "strings"

func border(char string, repeatno int) string {
	chars := strings.Repeat(char, repeatno) + "\n"
	return chars
}
