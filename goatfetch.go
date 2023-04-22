package main

import (
    "bufio"
    "fmt"
    "os"
    "os/exec"
    "strings"
)

func main() {
    shellPath := os.Getenv("SHELL")
    shell := strings.TrimPrefix(shellPath, "/bin/")
    term := os.Getenv("TERM")
    prettyName := ""
    kernelFetch := exec.Command("uname", "-r")
    kernelOutput, _ := kernelFetch.Output()
    file, _ := os.Open("/etc/os-release")
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if strings.HasPrefix(line, "PRETTY_NAME=") {
            prettyName = strings.Trim(strings.TrimPrefix(line, "PRETTY_NAME="), `"`)
        }
    }
    fmt.Println("")
    fmt.Println("  \033[1m(_(             ","\033[32mOS\033[0m:",prettyName)        
    fmt.Print("  \033[1m/_/'_____/)      ","\033[32mKernel\033[0m: ",string(kernelOutput))
    fmt.Println("  \033[1m\"  |      |     ","\033[32mShell\033[0m:",shell)
    fmt.Print("  \033[1m   |\"\"\"\"\"\"|      ","\033[32mTerminal\033[0m: ",term, "\n", "\n")
}

