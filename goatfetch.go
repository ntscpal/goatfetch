package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Get username
	logname := os.Getenv("LOGNAME")
	// Get hostname
	hostnameFetch := exec.Command("hostname")
	hostnameOutput, _ := hostnameFetch.Output()
	// Determine the number of dashes to seperate the username and hostname from the other statistics
	underNeed := len(logname) + len(hostnameOutput)
	// Determine the user's shell
	shellPath := os.Getenv("SHELL")
	// Only display the shell without its directory, e.g. 'bash' instead of '/bin/bash'
	shell := strings.TrimPrefix(shellPath, "/bin/")
	// Determine user's terminal
	term := os.Getenv("TERM")
	prettyName := ""
	// Fetch the kernel
	kernelFetch := exec.Command("uname", "-r")
	kernelOutput, _ := kernelFetch.Output()
	// Determine the 'pretty name' defined in /etc/os-release
	osReleaseFile, _ := os.Open("/etc/os-release")
	defer osReleaseFile.Close()
	scanner := bufio.NewScanner(osReleaseFile)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "PRETTY_NAME=") {
			prettyName = strings.Trim(strings.TrimPrefix(line, "PRETTY_NAME="), `"`)
		}
	}
	// Print the username
	fmt.Print("\033[32m\033[1m                   ", logname)
	// Print the atmark in 'user@hostname'
	fmt.Print("\033[0m@")
	// Print the computer hostname
	fmt.Print("\033[32m\033[1m", string(hostnameOutput), "\033[0m")
	// Blank spaces for formatting
	fmt.Print("                   ")
	// Print dashes to seperate the 'user@hostname' from the system statistics
	for i := 0; i < underNeed; i++ {
		fmt.Print("-")
	}
	// Start new line
	fmt.Println()
	// Print the os name, or specifically the 'pretty name'
	fmt.Println("  \033[1m(_(             ", "\033[32mOS\033[0m:", prettyName)
	// Print the kernel name
	fmt.Print("  \033[1m/_/'______/)     ", "\033[32mKernel\033[0m: ", string(kernelOutput))
	// Print the user's shell
	fmt.Println("  \033[1m\"  |      |     ", "\033[32mShell\033[0m:", shell)
	// Print the user's terminal
	fmt.Print("  \033[1m   |\"\"\"\"\"\"|      ", "\033[32mTerminal\033[0m: ", term, "\n", "\n")

	// Blank spaces for formatting
	fmt.Print("                   ")
	// Print ANSI colours
	fmt.Print("\033[0;30m███")
	fmt.Print("\033[0;31m███")
	fmt.Print("\033[0;32m███")
	fmt.Print("\033[0;33m███")
	fmt.Print("\033[0;34m███")
	fmt.Print("\033[0;35m███")
	fmt.Print("\033[0;36m███")
	fmt.Print("\033[0;37m███\n")

	// Blank spaces for formatting
	fmt.Print("                   ")
	// Print brighter ANSI colours
	fmt.Print("\033[0;90m███")
	fmt.Print("\033[0;91m███")
	fmt.Print("\033[0;92m███")
	fmt.Print("\033[0;93m███")
	fmt.Print("\033[0;94m███")
	fmt.Print("\033[0;95m███")
	fmt.Print("\033[0;96m███")
	fmt.Print("\033[0;97m███")
	fmt.Print("\033[0m\n\n")
}
