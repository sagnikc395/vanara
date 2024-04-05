package main

import (
	"bufio"
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

const VERSION = "0.0.1"

func handleMessage(_ any) {

}

func runLSP() {
	fmt.Println("hi")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split()

	for scanner.Scan() {
		msg := scanner.Text()
		handleMessage(msg)
	}
}

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("👋👋 %s , Try out 🦧-lang , v-%s \n", user.Username, VERSION)
	fmt.Printf("type in some commands 🍌\n")
	repl.Start(os.Stdin, os.Stdout)
}
