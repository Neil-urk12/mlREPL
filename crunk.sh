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
    echo ""
    read -p "Enter your choice : " choice
    echo ""

    case $choice in
        1)
            echo "Booting up Python... (Ctrl + D to exit or type exit)"
            echo ""
            exec python3
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
        5)
            echo "Exiting..."
            break
            ;;
        *)
            echo "Invalid choice"
            ;;
    esac
done