// repl/repl.go

package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/TangoEnSkai/interpreter-go/gopher/lexer"
	"github.com/TangoEnSkai/interpreter-go/gopher/token"
)

const PROMPT = "ʕ ◔ϖ◔ʔ >> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		l := lexer.New(line)
		fmt.Println(l)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
