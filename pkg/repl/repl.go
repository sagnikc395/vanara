package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/sagnikc395/vanara/pkg/lexer"
	"github.com/sagnikc395/vanara/pkg/token"
)

const VERSION = `0.1.0`

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf("%s", fmt.Sprintf("vanara %s >>>", VERSION))
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.NewLexer(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
