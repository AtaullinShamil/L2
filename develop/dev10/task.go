package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

func main() {
	var timeoutString string
	flag.StringVar(&timeoutString, "timeout", "10s", "timeout time")
	flag.Parse()

	if len(flag.Args()) != 2 {
		fmt.Println("use : go-telnet --timeout=10s <host> <port>")
		return
	}

	timeoutInt, _ := strconv.Atoi(timeoutString[:len(timeoutString)-1])
	timeoutDuration := time.Duration(timeoutInt) * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), timeoutDuration)
	defer cancel()

	conn, err := net.DialTimeout("tcp", flag.Args()[0]+":"+flag.Args()[1], timeoutDuration)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	defer conn.Close()

	go handleCtrlD(cancel)
	go copyTo(os.Stdout, conn, ctx)
	copyTo(conn, os.Stdin, ctx)
}

func handleCtrlD(cancel context.CancelFunc) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGQUIT)
	<-sigChan
	cancel()
}

func copyTo(dst io.Writer, src io.Reader, ctx context.Context) {
	buf := make([]byte, 1024)
	for {
		select {
		case <-ctx.Done():
			return
		default:
			n, err := src.Read(buf)
			if err != nil {
				if err == io.EOF || err == context.Canceled {
					return
				}
				log.Fatal(err)
			}
			_, err = dst.Write(buf[:n])
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
