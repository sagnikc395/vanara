package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/sagnikc395/vanara/pkg/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! Welcome to the Vanara programming environment!\n",
		user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
