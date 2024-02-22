package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("🫡 %s! Try out 🦍-lang\n", user.Username)
	fmt.Printf("type is some commands 🍌\n")
	repl.Start(os.Stdin, os.Stdout)
}
