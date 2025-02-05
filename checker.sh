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

# Create a temporary file to store language status
temp_file="/tmp/installed_langs.txt"
> "$temp_file"

# Check each language and store status
if command_exists python3; then
    echo "python=1" >> "$temp_file"
else
    echo "python=0" >> "$temp_file"
    echo "Python is not installed"
    install_package "Python" "python3"
fi

if command_exists go; then
    echo "go=1" >> "$temp_file"
else
    echo "go=0" >> "$temp_file"
    echo "Go is not installed"
    install_package "Go" "golang"
fi

if command_exists node; then
    echo "node=1" >> "$temp_file"
else
    echo "node=0" >> "$temp_file"
    echo "JavaScript (Node.js) is not installed"
    install_package "Node.js" "nodejs"
fi

if command_exists rustc; then
    echo "rust=1" >> "$temp_file"
else
    echo "rust=0" >> "$temp_file"
    echo "Rust is not installed"
    install_package "Rust" "rust"
fi

echo -e "Language check complete..\nPress any key to continue..."
read -n 1 -s -r
