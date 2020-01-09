package main

import ("fmt"
		"log"
		"os/exec")

func main(){
	ps := "C:\\Windows\\System32\\WindowsPowerShell\\v1.0\\powershell.exe"

	cmd := exec.Command(ps,"c:\\users\\hurd\\Desktop","ls")

	out,err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out) 
}	