# Go REPL

This program implements a simple Read-Eval-Print Loop (REPL) for the Go programming language. It allows users to interactively enter Go code, have it executed, and see the results printed to the console.

## Usage

To run the REPL, execute the `main.go` file using the Go toolchain:

```bash
go run main.go
```

Once the REPL is running, you can enter Go code at the prompt (`go> `). Press `Shift+Enter` to enter multi-line input, and press `Enter` on a blank line or enter a complete code block to execute it. Type `exit` or press `Ctrl+D` to quit the REPL.

## Code Structure

### Package Declaration

```go
package main
```

The code belongs to the `main` package, indicating that it is an executable program.

### Imports

```go
import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)
```

The program imports necessary standard library packages:

-   `bufio`: For buffered I/O, used to read user input.
-   `fmt`: For formatted I/O, used to print output to the console.
-   `os`: For operating system functionality, used to create temporary directories and interact with the file system.
-   `os/exec`: For executing external commands, used to run the Go code.
-   `path/filepath`: For manipulating file paths, used to construct paths to temporary files.
-   `strings`: For string manipulation, used to process user input and code.

### REPL Struct

```go
type REPL struct {
	scanner   *bufio.Scanner
	buffer    []string
	functions []string
	types     []string
	vars      []string
}
```

The `REPL` struct represents the REPL environment and holds its state:

-   `scanner`: A `bufio.Scanner` to read user input from standard input.
-   `buffer`: A string slice to store multi-line input before execution.
-   `functions`: A string slice to store function declarations.
-   `types`: A string slice to store type declarations.
-   `vars`: A string slice to store variable declarations.

### `NewREPL` Function

```go
func NewREPL() *REPL {
	return &REPL{
		scanner:   bufio.NewScanner(os.Stdin),
		buffer:    make([]string, 0),
		functions: make([]string, 0),
		types:     make([]string, 0),
		vars:      make([]string, 0),
	}
}
```

`NewREPL` is a constructor function that creates and initializes a new `REPL` instance. It sets up the `bufio.Scanner` to read from standard input and initializes the `buffer`, `functions`, `types`, and `vars` slices.

### `Run` Function

```go
func (r *REPL) Run() {
	// ...
}
```

The `Run` method starts the REPL, continuously reading user input, evaluating it, and printing the output. It handles multi-line input and checks for the `exit` command or `Ctrl+D` to terminate the REPL.

### `isCompleteInput` Function

```go
func isCompleteInput(input string) bool {
	// ...
}
```

`isCompleteInput` checks if the given input string represents a complete Go code block. It determines this by counting the number of opening and closing braces and handling cases with simple expressions or statements without braces.

### `eval` Function

```go
func (r *REPL) eval(input string) {
	// ...
}
```

The `eval` method evaluates the given Go code input. It creates a temporary directory, wraps the input into a valid Go program using `wrapCode`, writes the code to a temporary file, and then executes the file using the `go run` command. The output of the execution is printed to the console.

### `wrapCode` Function

```go
func (r *REPL) wrapCode(input string) string {
	// ...
}
```

`wrapCode` takes the user input and wraps it into a valid Go program. It handles type and variable declarations, separating package-level and function-level variables. The resulting program is returned as a string.

### `main` Function

```go
func main() {
	// ...
}
```

The `main` function is the entry point of the program. It prints a welcome message, creates a new `REPL` instance using `NewREPL`, and starts the REPL by calling the `Run` method.

## Additional Notes

-   The REPL currently has limited error handling and may not handle all edge cases gracefully.
-   The code could be further improved by adding support for more Go language features and providing better error messages.
