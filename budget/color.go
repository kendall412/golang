package main

import (
	"os"
	"strings"

	"github.com/fatih/color"
)

/*
Adding color to text output
*/
var red = color.New(color.FgRed, color.Bold)
var green = color.New(color.FgGreen, color.Bold)
var blue = color.New(color.FgHiBlue, color.Bold)
var yellow = color.New(color.FgHiYellow, color.Bold)
var italic = color.New(color.Italic, color.FgYellow, color.Bold)
var magenta = color.New(color.FgMagenta, color.Bold)

func printError(msg string) {
	repeatno := 40
	red.Println(strings.Repeat("*", repeatno))
	red.Println("ERROR: " + msg)
	red.Println(strings.Repeat("*", repeatno))
	os.Exit(EXIT_CODE_ERROR)
}

func printSum(sum float64) {
	blue.Printf("TOTAL: ")
	green.Printf("$%.2f\n", sum)
}

func printInfo(msg string) {
	yellow.Println(msg)
}
