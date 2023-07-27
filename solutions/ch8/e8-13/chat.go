// Exercise 8.13: Make the chat server disconnect idle clients, such as those that have sent no
// messages in the last five minutes. Hint: calling conn.Close() in another goroutine unblocks
// active Read calls such as the one done by input.Scan().
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

// !+broadcaster
type client chan<- string // an outgoing message channel

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
				cli <- msg
			}

		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

//!-broadcaster

// !+handleConn
func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	idleReset := make(chan struct{})
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ch

	go func(conn net.Conn) {
		const timeLimit = 10
		ticker := time.NewTicker(1 * time.Second)
		for countdown := timeLimit; countdown > 0; countdown-- {
			select {
			case <-ticker.C:
				//do nothing
			case <-idleReset:
				countdown = timeLimit
			}
			fmt.Printf("%s: %d\n", who, countdown)
		}
		conn.Close()
	}(conn)

	input := bufio.NewScanner(conn)
	for input.Scan() {
		go func() {
			idleReset <- struct{}{}
		}()
		messages <- who + ": " + input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- ch
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
