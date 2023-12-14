/*
Scanner will error with lines longer than 65536 characters. If you know your line length is greater than 64K, use the Buffer() method to increase the scanner's capacity:

https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
...
scanner := bufio.NewScanner(file)

const maxCapacity int = longLineLen  // your required line length
buf := make([]byte, maxCapacity)
scanner.Buffer(buf, maxCapacity)

for scanner.Scan() {
...
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func seeIfHas(f *os.File, text string) {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if scanner.Text() == "i love march." {
			fmt.Println(scanner.Text())
		}
	}
}

func seeIfItContains(f *os.File, text string) {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), text) {
			fmt.Println(scanner.Text())
		}
	}
}

func main() {
	file, _ := os.Open("sample_text.txt")

	sampleTxt := "i love"
	// seeIfHas(file, sampleTxt)
	seeIfItContains(file, sampleTxt)

}
