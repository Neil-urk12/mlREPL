package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type REPL struct {
	scanner   *bufio.Scanner
	buffer    []string
	functions []string // Store function declarations
	types     []string // Store type declarations
	vars      []string // Store variable declarations
}

func NewREPL() *REPL {
	return &REPL{
		scanner:   bufio.NewScanner(os.Stdin),
		buffer:    make([]string, 0),
		functions: make([]string, 0),
		types:     make([]string, 0),
		vars:      make([]string, 0),
	}
}

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

// isCompleteInput checks if the input code block is complete
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

	// Rest of the function remains the same
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

func main() {
	fmt.Println("")
	fmt.Println(" ---------------------------------------")
	fmt.Println("|==Golang Interactive Code Environment==|")
	fmt.Println(" ---------------------------------------")
	fmt.Println("")
	repl := NewREPL()
	repl.Run()
}
