// Package main implements a simple REPL (Read-Eval-Print Loop) for Go.
// It allows users to enter Go code snippets which are then executed
// and the output is printed to the console.
package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// REPL represents the REPL environment.
type REPL struct {
	scanner   *bufio.Scanner
	buffer    []string
	functions []string
	types     []string
	pkgVars   []string // Package-level variables
	localVars []string // Function-level variables
	lastExpr  string
}

// NewREPL creates and initializes a new REPL instance.
func NewREPL() *REPL {
	return &REPL{
		scanner:   bufio.NewScanner(os.Stdin),
		buffer:    make([]string, 0),
		functions: make([]string, 0),
		types:     make([]string, 0),
		pkgVars:   make([]string, 0),
		localVars: make([]string, 0),
	}
}

// Run starts the REPL continuously reading evaluating and printing input.
func (r *REPL) Run() {
	fmt.Println("Go REPL (Press Shift+Enter for new line Ctrl+D or 'exit' to quit)")
	fmt.Println("")

	for {
		if len(r.buffer) == 0 {
			fmt.Print("go> ")
		} else {
			fmt.Print("... ")
		}

		scanned := r.scanner.Scan()
		if !scanned {
			if r.scanner.Err() == nil {
				fmt.Println("\nGoodbye!")
				break
			}
			fmt.Printf("Error reading input: %v\n", r.scanner.Err())
			break
		}

		input := r.scanner.Text()
		if input == "exit" && len(r.buffer) == 0 {
			fmt.Println("\nGoodbye!")
			break
		}

		r.buffer = append(r.buffer, input)

		// Check if the input ends with a blank line or is complete
		if input == "" || isCompleteInput(strings.Join(r.buffer, "\n")) {
			if len(r.buffer) > 0 {
				r.eval(strings.Join(r.buffer, "\n"))
				r.buffer = r.buffer[:0] // Clear the buffer
			}
		}
	}
}

// isCompleteInput checks if the input code block is complete.
// It does this by checking if the number of opening braces '{' matches
// the number of closing braces '}'. It also handles cases where the input
// is a simple expression or statement without braces.
func isCompleteInput(input string) bool {
	input = strings.TrimSpace(input)
	if input == "" {
		return false
	}

	// Count braces and parentheses
	openBraces := strings.Count(input, "{")
	closeBraces := strings.Count(input, "}")
	openParens := strings.Count(input, "(")
	closeParens := strings.Count(input, ")")

	// Handle control structures
	if strings.HasPrefix(input, "if ") || strings.HasPrefix(input, "for ") {
		return openBraces == closeBraces
	}

	// Check for struct initialization
	if strings.Contains(input, "struct{") || strings.Contains(input, "struct {") {
		return openBraces == closeBraces
	}

	// Handle multi-line slice/array literals
	if strings.Contains(input, "[") {
		lines := strings.Split(input, "\n")
		lastLine := strings.TrimSpace(lines[len(lines)-1])
		if !strings.HasSuffix(lastLine, ",") && !strings.HasSuffix(lastLine, "]") {
			return false
		}
	}

	// If it's a simple expression or statement
	if openBraces == 0 && openParens == closeParens &&
		!strings.HasSuffix(input, "{") && !strings.HasSuffix(input, ",") {
		return true
	}

	return openBraces == closeBraces && openParens == closeParens
}

// eval evaluates the given Go code input.
// It creates a temporary directory writes the code to a temporary file
// and then executes the file using the 'go run' command.
// The output of the execution is then printed to the console.
func (r *REPL) eval(input string) {
	// Create a temporary directory for our code
	tmpDir, err := os.MkdirTemp("", "gorepl")
	if err != nil {
		fmt.Printf("Error creating temp dir: %v\n", err)
		return
	}
	defer os.RemoveAll(tmpDir)

	// Wrap the input in a proper Go program
	program := r.wrapCode(input)

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

// wrapCode wraps the user input into a valid Go program.
// It handles type and variable declarations and it separates
// package-level and function-level variables.
// The resulting program is then returned as a string.
func (r *REPL) wrapCode(input string) string {
	trimmedInput := strings.TrimSpace(input)

	// Handle type declarations
	if strings.HasPrefix(trimmedInput, "type ") {
		r.types = append(r.types, input)
		return fmt.Sprintf(`package main

import "fmt"

%s

func main() {
    fmt.Println("Type defined successfully")
}`, strings.Join(r.types, "\n"))
	}

	// Handle package-level variable declarations
	if strings.HasPrefix(trimmedInput, "var ") {
		r.pkgVars = append(r.pkgVars, input)
		varName := strings.Split(strings.Split(input, " ")[1], "=")[0]
		return fmt.Sprintf(`package main

import "fmt"

%s

%s

func main() {
    fmt.Printf("Variable declared: %%v\n", %s)
}`, strings.Join(r.types, "\n"),
			strings.Join(r.pkgVars, "\n"),
			varName)
	}

	// Handle function declarations
	if strings.HasPrefix(trimmedInput, "func ") && !strings.HasPrefix(trimmedInput, "func main") {
		r.functions = append(r.functions, input)
		return fmt.Sprintf(`package main

import "fmt"

%s

%s

%s

func main() {
    fmt.Println("Function defined successfully")
}`, strings.Join(r.types, "\n"),
			strings.Join(r.pkgVars, "\n"),
			strings.Join(r.functions, "\n"))
	}

	// Handle := declarations (local variables)
	if strings.Contains(trimmedInput, ":=") {
		r.localVars = append(r.localVars, input)
		varName := strings.Split(input, ":=")[0]
		return fmt.Sprintf(`package main

import "fmt"

%s

%s

%s

func main() {
    %s
    fmt.Printf("Variable declared: %%v\n", %s)
}`, strings.Join(r.types, "\n"),
			strings.Join(r.pkgVars, "\n"),
			strings.Join(r.functions, "\n"),
			strings.Join(r.localVars, "\n    "),
			strings.TrimSpace(varName))
	}

	// Handle function calls and expressions
	return fmt.Sprintf(`package main

import "fmt"

%s

%s

%s

func main() {
    %s
    %s
}`, strings.Join(r.types, "\n"),
		strings.Join(r.pkgVars, "\n"),
		strings.Join(r.functions, "\n"),
		strings.Join(r.localVars, "\n    "),
		input)
}

// Helper method to get local variable declarations
func (r *REPL) getLocalVars() []string {
	localVars := []string{}
	for _, v := range r.localVars {
		if !strings.HasPrefix(strings.TrimSpace(v), "var ") {
			localVars = append(localVars, v)
		}
	}
	return localVars
}

// main is the entry point of the program.
func main() {
	fmt.Println("")
	fmt.Println(" ---------------------------------------")
	fmt.Println("|==Golang Interactive Code Environment==|")
	fmt.Println(" ---------------------------------------")
	fmt.Println("")
	repl := NewREPL()
	repl.Run()
}
