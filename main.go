package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

const VERSION = "0.0.1"

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hi %s , Try out monkey-lang , v-%s \n", user.Username, VERSION)
	fmt.Printf("type in some commands 🍌\n")
	repl.Start(os.Stdin, os.Stdout)
}
