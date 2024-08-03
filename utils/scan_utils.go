package utils

import (
	"os"
	"os/exec"

	"github.com/StaticV0yd/autorecon/data"
)

func Scan(args string) []data.Host {
	shell, _ := os.LookupEnv("SHELL")
	cmd := &exec.Cmd{
		Path:   shell,
		Args:   []string{shell, "-c", "nmap " + args + " -oN scan.nmap"},
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Stdin:  os.Stdin,
	}

	cmd.Run()
	return ReadNmap()
}

func SuScan(args string) []data.Host {
	shell, _ := os.LookupEnv("SHELL")
	cmd := &exec.Cmd{
		Path:   shell,
		Args:   []string{shell, "-c", "sudo nmap " + args + " -oN scan.nmap"},
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Stdin:  os.Stdin,
	}

	cmd.Run()
	return ReadNmap()
}
