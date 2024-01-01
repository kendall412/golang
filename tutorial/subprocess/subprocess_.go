package main

import (
	"fmt"
	"os/exec"
)

func execProc(cmd string, opt []string, url string){
	fmt.Println("cmd: ",cmd)
	fmt.Print("opt: ")
	fmt.Print(opt)

	_,err:=exec.Command(cmd,opt...).Output()
	if err != nil{
		fmt.Println(err)
	}
}

// func execProc_(dest string, url string, display bool) {
// 	cmd := "yt-dlp"
// 	opt := []string{
// 		"-x",
// 		"--audio-format",
// 		"mp3",
// 		"--path",
// 	}

// 	// displays the string of commands to be executed
// 	opt = append(opt, dest, url)
// 	if display {
// 		fmt.Println(cmd, opt)
// 	}

// 	// executes the command
// 	_, err := exec.Command(cmd, opt...).Output()

// 	if err != nil {
// 		fmt.Println(err)
// 	}

func main(){
	// opt := []string
	url := "https://www.youtube.com/watch?v=MXIKzekRAj8&pp=ygURaXZlIGJhZGRpZSBhdWRpbyA%3D"
	// cmd := "yt"
	cmd := "yt-dlp"
	// opt := []string{"-au","-x",url}
	opt := []string{"-x","--audio-format","mp3","--path",dest}

	execProc(cmd, opt, url)
}