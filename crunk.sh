#!/bin/bash
clear
./checker.sh

# Read installed languages from temp file
temp_file="/tmp/installed_langs.txt"
declare -A langs
while IFS='=' read -r lang status; do
    langs["$lang"]=$status
done < "$temp_file"

while true; do
    # Count installed languages to help with menu formatting
    installed_count=0
    for lang in "${!langs[@]}"; do
        if [ "${langs[$lang]}" == "1" ]; then
            ((installed_count++))
        fi
    done

    # Calculate padding for centered text
    padding=$((25))
    echo
    echo " $(printf '%*s' $padding "")"
    echo " -------------------------------------------"
    echo "|           Neik's Crunk Shell              |"
    echo "|                                           |"

    # Initialize menu counter
    counter=1
    declare -A menu_items

    # Display installed languages
    if [ "${langs[python]}" == "1" ]; then
        echo "|    $counter. Python                              |"
        menu_items[$counter]="python"
        ((counter++))
    fi
    if [ "${langs[go]}" == "1" ]; then
        echo "|    $counter. Go                                  |"
        menu_items[$counter]="go"
        ((counter++))
    fi
    if [ "${langs[node]}" == "1" ]; then
        echo "|    $counter. JavaScript                          |"
        menu_items[$counter]="node"
        ((counter++))
    fi
    if [ "${langs[rust]}" == "1" ]; then
        echo "|    $counter. Rust                                |"
        menu_items[$counter]="rust"
        ((counter++))
    fi
    if [ "${langs[java]}" == "1" ]; then
        echo "|    $counter. Java                                |"
        menu_items[$counter]="java"
        ((counter++))
    fi

    # Display Exit option
    echo "|                                           |"
    echo "|    $counter. Exit                                |"
    echo "|                                           |"
    echo " -------------------------------------------"
    
    # Display missing languages message
    if [ $installed_count -lt 4 ]; then
        echo -e "\nMissing languages:"
        if [ "${langs[python]}" == "0" ]; then
            echo "- Python"
        fi
        if [ "${langs[go]}" == "0" ]; then
            echo "- Go"
        fi
        if [ "${langs[node]}" == "0" ]; then
            echo "- JavaScript (Node.js)"
        fi
        if [ "${langs[rust]}" == "0" ]; then
            echo "- Rust"
        fi
        if [ "${langs[java]}" == "0" ]; then
            echo "- Java"
        fi
        echo -e "\nRun the script again to install missing languages."
    fi

    echo -e "\nSelect a language: "
    echo ""
    read -p "Enter your choice : " choice
    echo ""

    if [ "$choice" == "$counter" ]; then
        echo "Exiting..."
        echo "Ciao!"
        break
    elif [[ "$choice" =~ ^[0-9]+$ ]] && [ "$choice" -gt 0 ] && [ "$choice" -lt "$counter" ]; then
        selected=${menu_items[$choice]}
        case $selected in
            "python")
                echo "Booting up Python... (Ctrl + D to exit or type exit)"
                echo ""
                exec python3
                ;;
            "go")
                echo "Booting up Go...HERE WE GO! (Ctrl + D to exit or type exit)"
                echo ""
                exec gore
                ;;
            "node")
                echo "I'm gonna JavaScript on your back! (Ctrl + D to exit or type exit)"
                echo ""
                exec node
                ;;
            "rust")
                echo "Rust is the best! (Ctrl + D to exit or type exit)"
                echo ""
                exec evcxr
                ;;
            "java")
                echo "Wooaahhh. Are we writing some legacy code? (Ctrl + D to exit or type exit)"
                echo ""
                exec evcxr
                ;;
        esac
    else
        echo "Invalid choice"
    fi
done
