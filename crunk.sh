#!/bin/bash
clear
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
            echo "I'm gonna JavaScript on your back! (Ctrl + D to exit or type exit)"
            echo ""
            exec node
            ;;
        4)
            echo "Rust is the best! (Ctrl + D to exit or type exit)"
            echo ""
            exec evcxr
            ;;
        5)

            echo "Exiting..."
            echo "Ciao!"
            break
            ;;
        *)
            echo "Invalid choice"
            ;;
    esac
done