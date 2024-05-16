package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		cmdArgs := strings.Split(text, " ")

		if len(cmdArgs) == 1 && cmdArgs[0] == "" {
			continue
		}

		switch cmdArgs[0] {
		case "cd":
			if len(cmdArgs) > 1 {
				changeDirectory(cmdArgs[1])
			} else {
				changeDirectory("")
			}

		case "pwd":
			printWorkingDirectory()

		case "echo":
			if len(cmdArgs) > 1 {
				echoFunc(cmdArgs[1:])
			} else {
				fmt.Println()
			}

		case "kill":
			if len(cmdArgs) == 1 {
				fmt.Println("kill: not enough arguments")
			} else {
				killProcess(cmdArgs[1:])
			}

		case "ps":
			listProcesses()

		case "nc":
			netcat(cmdArgs[1:])

		case "exit":
			os.Exit(0)

		default:
			var cmd *exec.Cmd
			if len(cmdArgs) > 1 {
				cmd = exec.Command(cmdArgs[0], cmdArgs[1:]...)
			} else {
				cmd = exec.Command(cmdArgs[0])
			}
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		}
	}
}

func changeDirectory(dir string) {
	if dir == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		err = os.Chdir(homeDir)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	} else {
		err := os.Chdir(dir)
		if err != nil {

		}
	}
}

func printWorkingDirectory() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Println(dir)
	}
}

func echoFunc(args []string) {
	for ind, arg := range args {
		if ind == len(args)-1 {
			fmt.Print(arg + "\n")
		} else {
			fmt.Print(arg + " ")
		}
	}
}

func killProcess(pids []string) {
	for _, pid := range pids {
		p, err := strconv.Atoi(pid)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		process, err := os.FindProcess(p)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		err = process.Signal(syscall.SIGTERM)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
	}
}

func listProcesses() {
	cmd := exec.Command("ps")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Println(out.String())
	}
}

func netcat(cmdArgs []string) {
	if len(cmdArgs) < 3 {
		fmt.Println("use : nc <tcp/udp> <host> <port>")
		return
	}
	conn, err := net.Dial(cmdArgs[0], cmdArgs[1]+":"+cmdArgs[2])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	defer conn.Close()

	fmt.Println("Connected to ", cmdArgs[1]+":"+cmdArgs[2])

	go func() {
		_, err := io.Copy(conn, os.Stdin)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
	}()

	_, err = io.Copy(os.Stdout, conn)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	conn.Close()
}
