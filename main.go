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
}

func NewREPL() *REPL {
	return &REPL{
		scanner:   bufio.NewScanner(os.Stdin),
		buffer:    make([]string, 0),
		functions: make([]string, 0),
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
	// Detect if input is a function declaration
	if strings.HasPrefix(strings.TrimSpace(input), "func ") {
		r.functions = append(r.functions, input)
		return fmt.Sprintf(`package main

import "fmt"

%s

func main() {
	// Function declared
	fmt.Println("Function defined successfully")
}
`, input)
	}

	// Handle print statements
	if strings.HasPrefix(input, "print(") {
		input = fmt.Sprintf("fmt.%s", strings.Replace(input, "print", "Println", 1))
	}

	// Only wrap in fmt.Println if it's a pure expression, not a statement
	trimmedInput := strings.TrimSpace(input)
	isStatement := strings.HasSuffix(trimmedInput, ")") || // function call
		strings.Contains(trimmedInput, "=") || // assignment
		strings.HasPrefix(trimmedInput, "for") || // control structures
		strings.HasPrefix(trimmedInput, "if") ||
		strings.HasPrefix(trimmedInput, "fmt.") // already has print

	if !isStatement {
		input = fmt.Sprintf("fmt.Println(%s)", input)
	}

	// Collect imports
	imports := []string{"fmt"}
	if strings.Contains(input, "strings.") {
		imports = append(imports, "strings")
	}
	if strings.Contains(input, "math.") {
		imports = append(imports, "math")
	}

	importStmt := ""
	if len(imports) == 1 {
		importStmt = fmt.Sprintf("import \"%s\"", imports[0])
	} else if len(imports) > 1 {
		importStmt = "import (\n"
		for _, imp := range imports {
			importStmt += fmt.Sprintf("\t\"%s\"\n", imp)
		}
		importStmt += ")"
	}

	// Include all previously defined functions
	funcs := strings.Join(r.functions, "\n\n")

	return fmt.Sprintf(`package main

%s

%s

func main() {
	%s
}
`, importStmt, funcs, input)
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
