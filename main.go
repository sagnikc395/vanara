package main

import (
	"bufio"
	"fmt"
	"log"
	"monkey/lsp/rpc"
	"monkey/repl"
	"os"
	"os/user"
)

const VERSION = "0.0.1"

func handleMessage(logger *log.Logger, msg any) {
	logger.Println(msg)
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("didn't give a good file")
	}
	return log.New(logfile, "[educationalsp]", log.Ldate|log.Ltime|log.Lshortfile)
}

func runLSP() {
	logger := getLogger("/home/sagnikc/Personal/monkey/lsp/log.txt")
	logger.Println("Logging started!")
	fmt.Println("hi")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Text()
		handleMessage(logger, msg)
	}
}

func main() {
	runLSP()
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("👋👋 %s , Try out 🦧-lang , v-%s \n", user.Username, VERSION)
	fmt.Printf("type in some commands 🍌\n")
	repl.Start(os.Stdin, os.Stdout)
}
