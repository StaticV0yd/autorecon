package utils

import (
	"os"
	"os/exec"
)

func Scan(args string) {
	shell, _ := os.LookupEnv("SHELL")
	cmd := &exec.Cmd{
		Path:   shell,
		Args:   []string{shell, "-c", "sudo nmap " + args + " -oA scan"},
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Stdin:  os.Stdin,
	}

	cmd.Run()

}
