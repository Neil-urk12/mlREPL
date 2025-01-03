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
            echo "Booting up Go...HERE WE GO! (Ctrl + D to exit or type exit)"
            echo ""
            exec gore
            ;;
        3)
            read -p "Enter Rust code: " code
            temp_file="temp_$$.rs"
            echo "$code" > "$temp_file"
            rustc "$temp_file" && ./temp 2>/dev/null || echo "Error: Invalid Rust code"
            rm "$temp_file" temp 2>/dev/null
            ;;
        4)
            echo "I'm gonna JavaScript on your back! (Ctrl + D to exit or type exit)"
            echo ""
            exec node
            ;;
        5)
            echo "Exiting...\nCiao!"
            break
            ;;
        *)
            echo "Invalid choice"
            ;;
    esac
done