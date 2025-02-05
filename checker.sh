#!/bin/bash

# Function to check if a command exists
command_exists() {
    command -v "$1" >/dev/null 2>&1
}

# Function to install a package
install_package() {
    read -p "Do you want to install $1? (y/n): " choice
    case "$choice" in 
        y|Y )
            case "$(uname -s)" in
                Linux*)
                    if command_exists apt; then
                        sudo apt update && sudo apt install -y "$2"
                    elif command_exists dnf; then
                        sudo dnf install -y "$2"
                    elif command_exists yum; then
                        sudo yum install -y "$2"
                    elif command_exists pacman; then
                        sudo pacman -S --noconfirm "$2"
                    else
                        echo "Package manager not supported"
                    fi
                    ;;
                Darwin*)
                    if command_exists brew; then
                        brew install "$2"
                    else
                        echo "Please install Homebrew first"
                    fi
                    ;;
                *)
                    echo "Operating system not supported"
                    ;;
            esac
            ;;
        n|N )
            echo "Skipping installation of $1"
            ;;
        * )
            echo "Invalid choice"
            ;;
    esac
}

# Check all languages
if command_exists rustc && command_exists node && command_exists java && command_exists go && command_exists python3; then
    echo "All programming languages (Rust, Node.js, Java, Go, and Python) are already installed!"
else
    # Check Rust
    if ! command_exists rustc; then
        echo "Rust is not installed"
        install_package "Rust" "rust"
    fi

    # Check JavaScript (Node.js)
    if ! command_exists node; then
        echo "JavaScript (Node.js) is not installed"
        install_package "Node.js" "nodejs"
    fi

    # Check Java
    if ! command_exists java; then
        echo "Java is not installed"
        install_package "Java" "default-jdk"
    fi

    # Check Go
    if ! command_exists go; then
        echo "Go is not installed"
        install_package "Go" "golang"
    fi

    # Check Python
    if ! command_exists python3; then
        echo "Python is not installed"
        install_package "Python" "python3"
    fi
fi

echo -e "\nPress any key to continue..."
read -n 1 -s -r
