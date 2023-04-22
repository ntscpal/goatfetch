package main

import (
    "bufio"
    "fmt"
    "os"
    "os/exec"
    "strings"
)

func main() {
    logname := os.Getenv("LOGNAME")
    hostnameFetch := exec.Command("hostname")
    hostnameOutput, _ := hostnameFetch.Output()
    underNeed := len(logname) + len(hostnameOutput)
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
    fmt.Print("\033[32m\033[1m                   ",logname)
    fmt.Print("\033[0m@")
    fmt.Print("\033[32m\033[1m",string(hostnameOutput),"\033[0m")
    fmt.Print("                   ")
    for i := 0; i < underNeed; i++ {
        fmt.Print("-")
    }
    fmt.Println("")
    fmt.Println("  \033[1m(_(             ","\033[32mOS\033[0m:",prettyName)        
    fmt.Print("  \033[1m/_/'_____/)      ","\033[32mKernel\033[0m: ",string(kernelOutput))
    fmt.Println("  \033[1m\"  |      |     ","\033[32mShell\033[0m:",shell)
    fmt.Print("  \033[1m   |\"\"\"\"\"\"|      ","\033[32mTerminal\033[0m: ",term, "\n", "\n")
    
    fmt.Print("                   ")
    fmt.Print("\033[0;30m███")
    fmt.Print("\033[0;31m███")
    fmt.Print("\033[0;32m███")
    fmt.Print("\033[0;33m███")
    fmt.Print("\033[0;34m███")
    fmt.Print("\033[0;35m███")
    fmt.Print("\033[0;36m███")
    fmt.Print("\033[0;37m███\n")
    
    fmt.Print("                   ")
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
