package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
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
	// Create a temporary directory for our code
	tmpDir, err := os.MkdirTemp("", "gorepl")
	if err != nil {
		fmt.Printf("Error creating temp dir: %v\n", err)
		return
	}
	defer os.RemoveAll(tmpDir)

	// Wrap the input in a proper Go program
	program := wrapCode(input)
	
	// Write to a temporary file
	tmpFile := filepath.Join(tmpDir, "main.go")
	if err := os.WriteFile(tmpFile, []byte(program), 0644); err != nil {
		fmt.Printf("Error writing temp file: %v\n", err)
		return
	}

	// Run the program
	cmd := exec.Command("go", "run", tmpFile)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Execution error: %s\n", output)
		return
	}

	fmt.Print(string(output))
}

func wrapCode(input string) string {
	// If the input is a simple expression, wrap it in a print statement
	if !strings.Contains(input, ";") && !strings.Contains(input, "func") && !strings.Contains(input, "for") && !strings.Contains(input, "if") {
		input = fmt.Sprintf("fmt.Println(%s)", input)
	}

	return fmt.Sprintf(`package main

import (
	"fmt"
	"time"
	"strings"
	"math"
)

func main() {
	%s
}
`, input)
}