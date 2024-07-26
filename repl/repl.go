package repl

import (
	"fmt"
	"io"
	"ran/evaluator"
	"ran/lexer"
	"ran/object"
	"ran/parser"

	"github.com/chzyer/readline"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	rl, err := readline.New(PROMPT)
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	env := object.NewEnvironment()
	macroEnv := object.NewEnvironment()

	for {
		line, err := rl.Readline()
		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}

		if line == "clear" {
			fmt.Fprintf(out, "\033[H\033[2J")
			continue
		}

		if line == "exit" {
			break
		}

		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			PrintParserErrors(out, p.Errors())
			continue
		}

		evaluator.DefineMacros(program, macroEnv)
		expanded := evaluator.ExpandMacros(program, macroEnv)
		evaluated := evaluator.Eval(expanded, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func PrintParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Woops! We ran into some errors here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
