package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/emilioziniades/interpreterbook/repl"
)

func main() {
	if len(os.Args) > 1 {
		repl.EvaluateFile(os.Args[1], os.Stdout)
		return
	}

	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
