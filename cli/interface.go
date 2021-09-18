package cli

import (
	"bufio"
	"fmt"
	"os"
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
	//fmt.Println(input)
	if input == "exit\n" {
		return true
	}

	return false
}
