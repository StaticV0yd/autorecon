package utils

import (
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"unicode"

	"github.com/StaticV0yd/autorecon/data"
)

// Reading and parsing through the nmap scan file to populate hosts and host information
func ReadNmap() []data.Host {
	var file *os.File
	var contents []byte
	var filePath string = "scan.nmap"

	file, err := os.Open(filePath)
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}

	contents, err = io.ReadAll(file)
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}

	var host data.Host
	var port data.Port
	var fileWrite string
	var hosts []data.Host

	//print(string(contents))
	shell, _ := os.LookupEnv("SHELL")

	lines := strings.Split(string(contents), "\n")
	for _, line := range lines {
		if len(line) > 21 && line[0:21] == "Nmap scan report for " {
			host.Ports = *new([]data.Port)
			host.Address = *new(string)
			host.Address = line[21:]
			mkdir := &exec.Cmd{
				Path:   shell,
				Args:   []string{shell, "-c", "mkdir " + host.Address},
				Stdout: os.Stdout,
				Stderr: os.Stderr,
				Stdin:  os.Stdin,
			}
			hosts = append(hosts, host)
			mkdir.Run()
		} else if len(line) > 0 && unicode.IsDigit(rune(line[0])) {
			segments := strings.Split(line, " ")
			if strings.ContainsRune(segments[0], '/') {
				port.Number = *new(int)
				port.Data = *new(string)
				port.Protocol = *new(string)
				port.Number, _ = strconv.Atoi(strings.Split(segments[0], "/")[0])
				port.Protocol = strings.Split(segments[0], "/")[1]
				port.Data = line
				fileWrite = "./" + host.Address + "/" + strconv.Itoa(port.Number)
				addToFile := &exec.Cmd{
					Path:   shell,
					Args:   []string{shell, "-c", "echo \"" + line + "\" >> " + fileWrite},
					Stdout: os.Stdout,
					Stderr: os.Stderr,
					Stdin:  os.Stdin,
				}
				addToFile.Run()
			}
			//touch.Args = []string{shell, "-c", "touch " + fileWrite}
			//touch.Run()
		} else if len(line) > 0 && rune(line[0]) == '|' {
			addToFile := &exec.Cmd{
				Path:   shell,
				Args:   []string{shell, "-c", "echo \"" + line + "\" >> " + fileWrite},
				Stdout: os.Stdout,
				Stderr: os.Stderr,
				Stdin:  os.Stdin,
			}
			addToFile.Run()
		}

		if len(host.Address) > 0 {
			addToFile := &exec.Cmd{
				Path:   shell,
				Args:   []string{shell, "-c", "echo \"" + line + "\" >> ./" + host.Address + "/" + host.Address},
				Stdout: os.Stdout,
				Stderr: os.Stderr,
				Stdin:  os.Stdin,
			}
			addToFile.Run()
		}
	}
	return hosts
}
