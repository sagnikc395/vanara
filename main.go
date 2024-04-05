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
const logfile = "/Users/sagnik3/Desktop/sagnikc395/Research/Projects/monkey/log.txt"

func handleMessage(logger *log.Logger, msg any) {
	logger.Println(msg)
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("didn't give a good file")
	}
	return log.New(logfile, "[monkeylsp]", log.Ldate|log.Ltime|log.Lshortfile)
}

func runLSP() {
	logger := getLogger(logfile)
	logger.Println("Logging started!")
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
