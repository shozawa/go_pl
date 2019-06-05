package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

type client struct {
	name string
	ch   chan string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli.ch <- msg
			}
		case cli := <-entering:
			var names []string
			for cli := range clients {
				names = append(names, cli.name)
			}
			if len(clients) < 1 {
				cli.ch <- "nobody in here"
			} else {
				cli.ch <- strings.Join(names, ", ") + " in here"
			}
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.ch)
		}
	}
}

func handleConn(conn net.Conn) {
	var cli client
	cli.ch = make(chan string)
	go clientWriter(conn, cli.ch)

	cli.name = conn.RemoteAddr().String()
	cli.ch <- "You are " + cli.name
	messages <- cli.name + " has arrived"
	entering <- cli

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- cli.name + ": " + input.Text()
	}

	leaving <- cli
	messages <- cli.name + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
