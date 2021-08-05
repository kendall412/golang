package main

import (
	"fmt"
	"log"
	"strings"
	"os/exec"
	"os"
)

func doProc(cmd string, option string) {
	out, err := exec.Command(cmd, option).Output()
	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("the ls is %s",out)
}


func doProc2(cmd string, option string) {
	proc := exec.Command(cmd, option)
	output, err := proc.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("the %s %s output is %s", cmd, option, output)
}


func doProc3(cmd string) {
	args := strings.Split(cmd, " ")
	proc := exec.Command(args[0], args[1:]...)
	// runs the command
	out, err :=proc.CombinedOutput()
	// Output or CombinedOutput captures output of the cmd
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Output: %s", out)
}

func doProc4(cmd string, option string) {
	// This example displays the results to os.Stdout
	// but is not captured and is not manipulatable
	proc :=exec.Command(cmd, option)
	proc.Stdout = os.Stdout
	proc.Stderr = os.Stderr
	err := proc.Run()
	if err != nil {
		log.Fatal(err)
	}

}


func main() {
	cmd := "ls"
	option := "/Users/kendall"
	doProc4(cmd, option)

	// cmd2 := "ls -lh /Users/kendall"
	// doProc3(cmd2)
}
