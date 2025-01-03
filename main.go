// Package main implements a simple REPL (Read-Eval-Print Loop) for Go.
// It allows users to enter Go code snippets, which are then executed,
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
	scanner   *bufio.Scanner // Scanner for reading user input.
	buffer    []string      // Buffer to store multi-line input.
	functions []string      // Store function declarations.
	types     []string      // Store type declarations.
	vars      []string      // Store variable declarations.
}

// NewREPL creates and initializes a new REPL instance.
func NewREPL() *REPL {
	return &REPL{
		scanner:   bufio.NewScanner(os.Stdin),
		buffer:    make([]string, 0),
		functions: make([]string, 0),
		types:     make([]string, 0),
		vars:      make([]string, 0),
	}
}

// Run starts the REPL, continuously reading, evaluating, and printing input.
func (r *REPL) Run() {
	fmt.Println("Go REPL (Press Shift+Enter for new line, Ctrl+D or 'exit' to quit)")
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

	// Check for complete blocks
	openBraces := strings.Count(input, "{")
	closeBraces := strings.Count(input, "}")

	// If it's a simple expression or statement without braces
	if openBraces == 0 && !strings.HasSuffix(input, "{") && !strings.HasSuffix(input, ",") {
		return true
	}

	// Check if all blocks are closed
	return openBraces > 0 && openBraces == closeBraces
}

// eval evaluates the given Go code input.
// It creates a temporary directory, writes the code to a temporary file,
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
// It handles type and variable declarations, and it separates
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
}
`, input)
	}

	// Handle variable declarations
	if strings.Contains(input, ":=") || strings.HasPrefix(trimmedInput, "var ") {
		// Determine if it's a package-level or function-level declaration
		isPackageLevel := strings.HasPrefix(trimmedInput, "var ")
		r.vars = append(r.vars, input)

		if isPackageLevel {
			// For package-level declarations (var x type)
			return fmt.Sprintf(`package main

import "fmt"

%s

%s

func main() {
	fmt.Printf("Variable declared: %%v\n", %s)
}
`, strings.Join(r.types, "\n\n"),
				strings.Join(r.vars, "\n"),
				strings.Split(strings.Split(input, " ")[1], "=")[0]) // Extract var name
		} else {
			// For function-level declarations (:=)
			return fmt.Sprintf(`package main

import "fmt"

%s

%s

func main() {
	%s
	fmt.Printf("Variable declared: %%v\n", %s)
}
`, strings.Join(r.types, "\n\n"),
				strings.Join(r.vars[:len(r.vars)-1], "\n"),
				input,
				strings.Split(input, ":=")[0])
		}
	}

	// If not a type or variable declaration, treat it as regular code
	declarations := strings.Join(r.types, "\n\n")
	packageVars := []string{}
	localVars := []string{}

	// Separate package-level and function-level variables
	for _, v := range r.vars {
		if strings.HasPrefix(strings.TrimSpace(v), "var ") {
			packageVars = append(packageVars, v)
		} else {
			localVars = append(localVars, v)
		}
	}

	return fmt.Sprintf(`package main

import "fmt"

%s

%s

func main() {
	%s
	%s
}
`, declarations,
		strings.Join(packageVars, "\n"),
		strings.Join(localVars, "\n\t"),
		input)
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
