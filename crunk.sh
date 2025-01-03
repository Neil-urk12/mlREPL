#!/bin/bash
echo $'\n'

while true; do
    echo " -------------------------------------------"
    echo "|           Neik's Crunk Shell              |"
    echo "|                                           |"
    echo "|    1. Python    2. Go    3. JavaScript    |"
    echo "|                                           |"
    echo "|           4. Rust    5. Exit              |"
    echo "|                                           |"
    echo " -------------------------------------------"
    echo $'\n'"Select a language: "

    read -p "Enter your choice : " choice


    case $choice in
        1)
            read -p "Enter Python code: " code
            python -c "$code" 2>/dev/null || echo "Error: Invalid Python code"
            ;;
        2)
            read -p "Enter Go code: " code
            temp_file="temp_$$.go"
            echo "$code" > "$temp_file"
            go run "$temp_file" 2>/dev/null || echo "Error: Invalid Go code"
            rm "$temp_file"
            ;;
        3)
            read -p "Enter Rust code: " code
            temp_file="temp_$$.rs"
            echo "$code" > "$temp_file"
            rustc "$temp_file" && ./temp 2>/dev/null || echo "Error: Invalid Rust code"
            rm "$temp_file" temp 2>/dev/null
            ;;
        4)
            read -p "Enter JavaScript code: " code
            temp_file="temp_$$.js"
            echo "$code" > "$temp_file"
            node "$temp_file" 2>/dev/null || echo "Error: Invalid JavaScript code"
            rm "$temp_file" temp 2>/dev/null
            ;;
        *)
            echo "Invalid choice"
            ;;
    esac
