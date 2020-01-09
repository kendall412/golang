package main

import (
	"fmt"
	// "log"
	"os/exec")

func main(){
	//ps := "C:\\Windows\\System32\\WindowsPowerShell\\v1.0\\powershell.exe"

	cmd1 := exec.Command("powershell.exe date")
	cmd2 := exec.Command("powershell.exe ls")
	out1,err1 := cmd1.Output()
	out2,err2 := cmd2.Output()
	// err := cmd.Run()		
	if err1 != nil {
		panic(err1)
	}
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(out1)
	fmt.Println(out2)
}
