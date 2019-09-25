package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

func main() {
	listner, err := net.Listen("tcp", "localhost:8000")

	if err != nil {
		fmt.Println("Something goes bad")
		panic(err)
	}

	defer listner.Close()

	for {
		conn, err := listner.Accept()

		if err != nil {
			fmt.Println("Not accept connection")
			continue
		}

		go repeat(conn)
	}
}

func repeat(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		sendResponse(conn, scanner.Text())
	}
	return
}

func sendResponse(conn net.Conn, text string) {
	writer := bufio.NewWriter(conn)
	writer.WriteString("(((" + strings.ToUpper(text) + ")))\n")
	writer.Flush()
	time.Sleep(time.Second)

	writer.WriteString(" ((" + text + ")) \n")
	writer.Flush()
	time.Sleep(time.Second)

	writer.WriteString("  (" + strings.ToLower(text) + ")\n")
	writer.Flush()
	time.Sleep(time.Second)

	return
}
