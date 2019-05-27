package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	tcpConn, ok := conn.(*net.TCPConn)
	if !ok {
		log.Fatal("conn is not *net.TCPConn.")
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, tcpConn)
		log.Println("done")
		tcpConn.CloseRead()
		done <- struct{}{}
	}()
	mustCopy(tcpConn, os.Stdin)
	tcpConn.CloseWrite()
	<-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
