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
    fmt.Println("OS:",prettyName)        
    fmt.Print("Kernel: ",string(kernelOutput))
    fmt.Println("Shell:",shell)
    fmt.Print("Terminal: ",term, "\n")
}

