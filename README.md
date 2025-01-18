# Neik's Interactive Language Shell

This project provides an interactive shell that allows you to experiment with different programming languages.

## Overview

The primary way to interact with this project is through the `crunk.sh` script. This script presents a menu allowing you to choose between different language environments:

*   Python
*   Go
*   JavaScript
*   Rust

Selecting a language will launch the respective interpreter or REPL.

The `main.go` file contains an experimental implementation of a Go REPL. This can be accessed by selecting the 'Go' option within the `crunk.sh` menu.

## Usage

1. **Run the shell script:**

    ```bash
    ./crunk.sh
    ```

2. **Select a language:** Follow the on-screen prompts to choose the desired language environment.

    For example, selecting '2' will launch the Go REPL.

### Go REPL (Experimental)

The Go environment provided is an experimental REPL. You can enter Go code snippets, and they will be executed.

**Examples:**

- Simple expressions:

  ```
  go> 1 + 2
  3
  go> "Hello" + " World"
  Hello World
  ```

- Multiple statements:

  ```
  go> for i := 0; i < 3; i++ { fmt.Println(i) }
  0
  1
  2
  ```

- Function definitions and calls:

  ```
  go> func add(a, b int) int { return a + b }; fmt.Println(add(5, 3))
  8
  ```

- Conditionals:

  ```
  go> if x := 42; x > 40 { fmt.Println("Greater than 40!") }
  Greater than 40!
  ```

To exit the Go REPL, type `exit`.

## Available Features

*   Interactive shells for Python, Go, JavaScript, and Rust (via `crunk.sh`).
*   Experimental Go REPL implementation (`main.go`).
*   Basic Go syntax support in the experimental REPL.
