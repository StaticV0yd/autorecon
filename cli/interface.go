package cli

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/StaticV0yd/autorecon/data"
	"github.com/StaticV0yd/autorecon/utils"
)

func CommandLine() int {
	var hosts *[]data.Host = new([]data.Host)

	// in a loop:
	for {
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
		var leave bool = parse(input, hosts)
		if leave {
			fmt.Println("Exiting autorecon...")
			break
		}

		// Execute commands based on parsed input
	}

	return 0
}

func parse(input string, hosts *[]data.Host) bool { //TODO: Implement 'show' command which shows hosts, ports, data, etc.
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
	} else if len(input) >= 4 && input[0:4] == "scan" {
		var args string
		if len(input) > 4 {
			args = input[5:]
		}
		*hosts = utils.Scan(args)
		//hosts = append(hosts, utils.Scan(args)...)
	} else if len(input) >= 6 && input[0:6] == "suscan" { // Check for scan (will be using nmap for the scanning, so 'scan' is essentially an alias for nmap)
		var args string
		if len(input) > 6 {
			args = input[7:]
		}
		*hosts = utils.SuScan(args)
		//*hosts = append(*hosts, utils.SuScan(args)...)
	} else if len(input) >= 4 && input[0:4] == "show" { // Check for commands relating to the database
		if !(len(strings.Split(input, " ")) < 2) {
			if strings.Split(input, " ")[1] == "hosts" {
				for _, host := range *hosts {
					fmt.Print("'" + host.Address + "'" + " ")
				}
			}
		}
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

	} else if len(input) >= 4 && input[0:4] == "test" {
		*hosts = utils.ReadNmap()
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
