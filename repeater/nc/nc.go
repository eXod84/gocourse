package main

import (
	"io"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	go redirect(conn, os.Stdout)
	redirect(os.Stdin, conn)
}

func redirect(src io.Reader, dst io.Writer) {
	if _, err := io.Copy(dst, src); err != nil {
		panic(err)
	}
}
