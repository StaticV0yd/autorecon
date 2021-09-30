package cli

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func CommandLine() int {
	// in a loop:
	for true {
		// Create the prompt and get ready for user input
		USER, _ := os.LookupEnv("USER")
		PWD, _ := os.LookupEnv("PWD")

		//fmt.Print("\n\033[38;5;27m" + PWD + "\n\033[38;5;15m" + USER + "\033[38;5;1m@autorecon>\033[38;5;15m ")
		fmt.Print("\n" + USER + "\033[38;5;1m@autorecon\033[38;5;15m " + "% \033[38;5;27m" + PWD + "\033[38;5;15m\n->> ")

		stdin := bufio.NewReader(os.Stdin)
		input, err := stdin.ReadString('\n')
		if err != nil {
			fmt.Println("ERROR: Could not read from stdin.")
		}

		// Parse the input
		var leave bool
		leave = parse(input)
		if leave {
			fmt.Println("Exiting autorecon...")
			break
		}

		// Execute commands based on parsed input
	}

	return 0
}

func parse(input string) bool {
	// Get rid of any leading whitespace
	for string(input[0]) == " " {
		input = input[1:]
	}
	if string(input[len(input)-1]) == "\n" {
		input = input[0 : len(input)-1]
	}

	// Check to see if the user wants to exit
	if input == "exit" {
		return true
	} else if len(input) >= 4 && input[0:4] == "nmap" { // Check for nmap

	} else if len(input) >= 4 && input[0:4] == "show" { // Check for commands relating to the database

	} else { // Otherwise, pipe output to shell
		shell, _ := os.LookupEnv("SHELL")
		cmd := exec.Command(shell, "-c", input)
		err := cmd.Run()
		if err != nil {
			os.Stderr.WriteString(err.Error() + "\n")
		}
	}

	return false
}
