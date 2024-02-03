package main

import (
	"github.com/fatih/color"
)

/*
Adding color to text output
*/

// printf
var redf = color.New(color.FgRed, color.Bold).PrintfFunc()
var greenf = color.New(color.FgGreen, color.Bold).PrintfFunc()
var bluef = color.New(color.FgHiBlue, color.Bold).PrintfFunc()
var yellowf = color.New(color.FgHiYellow, color.Bold).PrintfFunc()
var italicf = color.New(color.Italic, color.FgYellow, color.Bold).PrintfFunc()

// println
var redln = color.New(color.FgRed, color.Bold).PrintlnFunc()
var greenln = color.New(color.FgGreen, color.Bold).PrintlnFunc()
var blueln = color.New(color.FgHiBlue, color.Bold).PrintlnFunc()
var yellowln = color.New(color.FgHiYellow, color.Bold).PrintlnFunc()
var italicln = color.New(color.Italic, color.FgYellow, color.Bold).PrintlnFunc()
