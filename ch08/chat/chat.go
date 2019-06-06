package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
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
				select {
				/* ex 8.15 */
				case cli.ch <- msg:
				default:
					// skip message
				}

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
	ping := make(chan struct{})
	cli.ch = make(chan string)
	go clientWriter(conn, cli.ch)

	/* ex 8.13 */
	// client killer
	go func() {
		for {
			select {
			case <-ping:
				// survive
			case <-time.After(5 * time.Minute):
				conn.Close()
			}
		}
	}()

	input := bufio.NewScanner(conn)

	/* ex 8.14 */
	cli.ch <- "What's your name?"
	input.Scan()
	cli.name = input.Text()

	messages <- cli.name + " has arrived"
	entering <- cli

	for input.Scan() {
		ping <- struct{}{}
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
