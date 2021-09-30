package cli

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func CommandLine() int {
	// in a loop:
	for true {
		// Create the prompt and get ready for user input
		USER, _ := os.LookupEnv("USER")
		//PWD, _ := os.LookupEnv("PWD")
		cmd := exec.Command("pwd")
		stdout, _ := cmd.StdoutPipe()
		cmd.Start()
		data, _ := io.ReadAll(stdout)
		cmd.Wait()
		PWD := string(data)
		PWD = PWD[:len(PWD)-1]

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

	} else if (len(input) == 2 && input[0:2] == "cd") || (len(input) > 2 && input[0:3] == "cd ") {
		var path string
		if len(input) > 2 {
			path = input[3:]
		} else {
			path, _ = os.UserHomeDir()
		}
		err := os.Chdir(path)
		if err != nil {
			os.Stderr.WriteString(err.Error())
		}

	} else { // Otherwise, pipe output to shell
		shell, _ := os.LookupEnv("SHELL")
		cmd := &exec.Cmd{
			Path:   shell,
			Args:   []string{shell, "-c", input},
			Stdout: os.Stdout,
			Stderr: os.Stderr,
			Stdin:  os.Stdin,
		}

		cmd.Run()
	}

	return false
}
