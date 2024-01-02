package main

import (
	"log"
	"os/exec"
)

func execProc(cmd string, opt []string, debug *bool) {

	log.Println("execPro - cmd: ", cmd)
	log.Print("execPro - opt: ")
	log.Print(opt)

	if !*debug {
		_, err := exec.Command(cmd, opt...).Output()
		if err != nil {
			log.Println(err)
		}
	}

}
