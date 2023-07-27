package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/emilioziniades/interpreterbook/evaluator"
	"github.com/emilioziniades/interpreterbook/lexer"
	"github.com/emilioziniades/interpreterbook/object"
	"github.com/emilioziniades/interpreterbook/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		evaluateText(line, env, out)
	}
}

func EvaluateFile(filename string, out io.Writer) {
	file, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	env := object.NewEnvironment()
	evaluateText(string(file), env, out)
}

func evaluateText(text string, env *object.Environment, out io.Writer) {

	l := lexer.New(text)
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		printParseErrors(out, p.Errors())
		return
	}

	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		io.WriteString(out, evaluated.Inspect())
		io.WriteString(out, "\n")
	}
}

const MONKEY_FACE = `            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
`

func printParseErrors(out io.Writer, errors []string) {
	io.WriteString(out, MONKEY_FACE)
	io.WriteString(out, "Whoops! We ran into some monkey business here!\n")
	io.WriteString(out, "parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
