package main

import (
	"fmt"
	"io"
	"os"
	"os/user"
	"ran/evaluator"
	"ran/lexer"
	"ran/object"
	"ran/parser"
	"ran/repl"
	"ran/replc"
)

func main() {

	env := object.NewEnvironment()
	macroEnv := object.NewEnvironment()
	out := os.Stdout

	if len(os.Args) > 1 && os.Args[1] != "--vm" {
		file := os.Args[1]
		dat, err := os.ReadFile(
			file,
		)
		if err != nil {
			fmt.Println("Error reading file")
			return
		}
		l := lexer.New(string(dat))
		p := parser.New(l)
		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			repl.PrintParserErrors(out, p.Errors())
			return
		}

		evaluator.DefineMacros(program, macroEnv)
		expanded := evaluator.ExpandMacros(program, macroEnv)
		evaluated := evaluator.Eval(expanded, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
		return
	}

	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the a programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")

	if len(os.Args) > 1 && os.Args[1] == "--vm" {
		replc.Start(os.Stdin, os.Stdout)
		return
	}
	repl.Start(os.Stdin, os.Stdout)
}
