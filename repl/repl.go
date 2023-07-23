package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/Kmelow/monke-go/lexer"
	"github.com/Kmelow/monke-go/token"
)

const PROMPT = "🐒~> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, "%s", PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
