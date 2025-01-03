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
	scanner *bufio.Scanner
}

func NewREPL() *REPL {
	return &REPL{
		scanner: bufio.NewScanner(os.Stdin),
	}
}

func (r *REPL) Run() {
	for {
		fmt.Print("go>")
		scanned := r.scanner.Scan()
		if !scanned {
			if r.scanner.Err() == nil {
				// EOF encountered (Ctrl+D)
				fmt.Println("\nGoodbye!")
				break
			}
			fmt.Printf("Error reading input: %v\n", r.scanner.Err())
			break
		}
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
	// Handle print statements
	if strings.HasPrefix(input, "print(") {
		input = fmt.Sprintf("fmt.%s", strings.Replace(input, "print", "Println", 1))
	}

	// If the input is a simple expression, wrap it in a print statement
	if !strings.Contains(input, ";") && 
	   !strings.Contains(input, "func") && 
	   !strings.Contains(input, "for") && 
	   !strings.Contains(input, "if") &&
	   !strings.Contains(input, "fmt.") {
		input = fmt.Sprintf("fmt.Println(%s)", input)
	}

	// Only include required imports
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

	return fmt.Sprintf(`package main

%s

func main() {
	%s
}
`, importStmt, input)
}

func main() {
	repl := NewREPL()
	repl.Run()
}
