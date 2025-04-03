/*
using "flag" module to create subcommands.
*/

package main

import (
	"flag"
	"fmt"
	"os"
)

func add(x, y int) int {
	return x + y
}

func getFlags() (*flag.FlagSet, *flag.FlagSet, *flag.FlagSet, *string, *string, *int, *int) {
	// subcommand "debug"
	debugCmd := flag.NewFlagSet("debug", flag.ExitOnError)
	debugStr := debugCmd.String("s", "", "enables debug mode")
	// subcommand "main"
	mainCmd := flag.NewFlagSet("main", flag.ExitOnError)
	mainStr := mainCmd.String("s", "", "enables main print")
	// subcommand "add"
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	arg_1 := addCmd.Int("x", 0, "1 of 2 argument for addCmd")
	arg_2 := addCmd.Int("y", 0, "2 of 2 argument for addCmd")

	return debugCmd, mainCmd, addCmd, debugStr, mainStr, arg_1, arg_2
}

func main() {
	debugCmd, mainCmd, addCmd, debugStr, mainStr, arg_1, arg_2 := getFlags()

	if len(os.Args) < 2 {
		fmt.Println("expect arguments")
		os.Exit(1)
	}

	switch os.Args[1] {

	case "debug":
		debugCmd.Parse(os.Args[2:])
		fmt.Println("subcmd: debug")
		fmt.Printf("you have chosen %s", *debugStr)

	case "main":
		mainCmd.Parse(os.Args[2:])
		fmt.Println("subcmd: main")
		fmt.Printf("you have chosen %s", *mainStr)

	case "add":
		addCmd.Parse((os.Args[2:]))
		fmt.Println("subcmd: addCmd")
		fmt.Printf("%d + %d = %d\n", *arg_1, *arg_2, add(*arg_1, *arg_2))

	default:
		fmt.Println("expecting 'debug' or 'main' subcommands")
		os.Exit(1)
	}
}
