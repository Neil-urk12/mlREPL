package main

import (
	"bufio"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

type REPL struct {
	scanner *bufio.Scanner
}

func NewREPL() *REPL {
	return &REPL {
		scanner: bufio.NewScanner(os.Stdin),
	}
}

func (r *REPL) Run() {
	for {
		fmt.Print("go>")
		r.scanner.Scan()
		input := r.scanner.Text()
		if input == "exit" {
			break
		}
		r.eval(input)
	}
}

func (r *REPL) eval(input string) {
	fset := token.NewFileSet()

	file, err := parser.ParseExpr(fset, "", input)
	if err != nil {
		fmt.Println(err)
		return
	}

	ast.Print(fset, file)

	
}