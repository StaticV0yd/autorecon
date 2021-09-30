package main

import (
	"os"

	"github.com/StaticV0yd/autorecon/cli"
)

func main() {
	// Ideally should not be run as root (root should only be used when running nmap commands and such that need it),
	//		so at the beginning, this should check to see if user is root and warn the user if they are root.
	user, _ := os.LookupEnv("USER")
	if user == "root" {
		os.Stderr.WriteString("ERROR: autorecon should not be run as root!\n")
		return
	}
	cli.CommandLine()

}
