package main

import (
	"fmt"
	"strings"
	"reflect"
)

func main() {
	s := "ffmpeg -ss 00:08:35 -i YDXJ0006.MP4 -t 00:00:30 -c:v copy -c:a copy -y YDXJ0006-short.MP4"

	args := strings.Split(s, " ")
	fmt.Println(args)
	fmt.Println(args[4])

	args2 := strings.Split(s, ":")
	fmt.Println(args2)
	fmt.Println(reflect.TypeOf(args2))
	fmt.Printf("length of %s is %d", args, len(args))
}