// Exercise 8.1: Modify clock2 to accept a port number, and write a program, clockwall, that
// acts as a client of several clock servers at once, reading the times from each one and displaying
// the results in a table, akin to the wall of clocks seen in some business offices. If you have
// access to geographically distributed computers, run instances remotely ; otherwise run local
// instances on different ports with fake time zones.
// $ TZ=US/Eastern
// ./clock2 -port 8010 &
// $ TZ=Asia/Tokyo
// ./clock2 -port 8020 &
// $ TZ=Europe/London ./clock2 -port 8030 &
// $ clockwall NewYork=localhost:8010 London=localhost:8020 Tokyo=localhost:8030
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

var port = flag.Int("port", 8000, "computer kind of port")

const defaultTimeZone = "America/Los_Angeles"

var timeZone = flag.String("tz", defaultTimeZone, "time zone :)")

func handleConn(c net.Conn) {
	loc, _ := time.LoadLocation(*timeZone)

	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().In(loc).Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	flag.Parse()
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatal(err)
	}
	//!+
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle connections concurrently
	}
	//!-
}
