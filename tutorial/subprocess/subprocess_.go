package main

import (
	"log"
	"flag"
	"os/exec"
)

func execProc(cmd string, opt []string, url string){
	log.Println("cmd: ",cmd)
	log.Print("opt: ")
	log.Print(opt)

	_,err:=exec.Command(cmd,opt...).Output()
	if err != nil{
		log.Println(err)
	}
}


func main(){
	debug := flag.Bool("de",false,"debug mode")
	// audio := flag.Bool("au",false,"download file as audio file")
	flag.Parse()
	// opt := []string
	url := "https://www.youtube.com/watch?v=MXIKzekRAj8&pp=ygURaXZlIGJhZGRpZSBhdWRpbyA%3D"
	cmd := "yt"
	// cmd := "yt-dlp"
	opt := []string{"-au","-x",url}
	// opt := []string{"-x","--audio-format","mp3","--path",dest}


	if *debug != true{
		execProc(cmd, opt, url)
		log.Println("debug: ",*debug)
	} 
	log.Println("debug: ",*debug)
	// log.Println(opt)
	
}