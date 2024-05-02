package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/sagnikc395/monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s ! welcome to the monkey programming language\n", user.Username)
	fmt.Printf("feel free to type in commands \n")
	repl.Start(os.Stdin, os.Stdout)
}
