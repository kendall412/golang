package main

import (
	"log"
	"os/exec"
)

func execProc(cmd string, opt []string, debug *bool) {
	cmd_ := cmd + " "
	for _, v := range opt {
		cmd_ = cmd_ + " " + v
	}
	log.Println("cmd:\n" + cmd_)

	if !*debug {
		_, err := exec.Command(cmd, opt...).Output()
		if err != nil {
			log.Println(err)
		}
	}
}
