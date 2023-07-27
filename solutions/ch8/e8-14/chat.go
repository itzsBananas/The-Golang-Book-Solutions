// Exercise 8.14: Change the chat server’s network protocol so that each client provides its name
// on entering. Use that name instead of the network address when prefixing each message with
// its sender’s identity.
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
)

// !+broadcaster
type client struct {
	addr string
	ch   chan<- string // an outgoing message channel
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli.ch <- msg
			}

		case cli := <-entering:
			clients[cli] = true
			var buf bytes.Buffer
			buf.WriteString("Current List of Users:\n")
			for cli := range clients {
				buf.WriteString(fmt.Sprintf("%s\n", cli.addr))
			}
			go func() { messages <- buf.String() }()

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.ch)
		}
	}
}

//!-broadcaster

// !+handleConn
func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	input := bufio.NewScanner(conn)
	ch <- "Enter name: "
	input.Scan()
	who := input.Text()

	cli := client{who, ch}

	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- cli

	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- cli
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

//!-handleConn

// !+main
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

//!-main
