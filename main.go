package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/sagnikc395/vanara/repl"
)

const VERSION = "0.0.1"

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("welcome %s! vanara v%s programming environemnt", user.Username, VERSION)
	fmt.Printf("feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
