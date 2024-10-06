package replc

import (
	"fmt"
	"io"
	"ran/compiler"
	"ran/lexer"
	"ran/parser"
	"ran/vm"

	"github.com/chzyer/readline"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	rl, err := readline.New(PROMPT)
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	for {
		line, rl_err := rl.Readline()
		if rl_err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if rl_err == io.EOF {
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

		comp := compiler.New()
		err := comp.Compile(program)
		if err != nil {
			fmt.Fprintf(out, "Woops! Compilation failed:\n %s\n", err)
			continue
		}
		machine := vm.New(comp.Bytecode())
		err = machine.Run()
		if err != nil {
			fmt.Fprintf(out, "Woops! Executing bytecode failed:\n %s\n", err)
			continue
		}
		lastPopped := machine.LastPoppedStackElem()
		io.WriteString(out, lastPopped.Inspect())
		io.WriteString(out, "\n")

	}
}

func PrintParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Woops! We ran into some errors here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
