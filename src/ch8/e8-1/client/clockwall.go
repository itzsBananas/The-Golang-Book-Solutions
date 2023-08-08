package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type clock struct {
	name, addr string
}

func (c *clock) watch(w io.Writer, r io.Reader) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		fmt.Fprintf(w, "%s: %s\n", c.name, s.Text())
	}
	fmt.Println(c.name, "done")
	if s.Err() != nil {
		log.Printf("can't read from %s: %s", c.name, s.Err())
	}
}

func parseArgs(args ...string) []clock {
	clockwall := make([]clock, 0, len(args))

	for _, str := range args {
		op := strings.Split(str, "=")
		if len(op) != 2 {
			continue
		}

		clockwall = append(clockwall, clock{op[0], op[1]})
	}

	return clockwall
}

func main() {
	if len(os.Args) == 1 {
		panic("invalid # of arg's")
	}

	clockwall := parseArgs(os.Args[1:]...)
	fmt.Println(clockwall)

	for _, clock := range clockwall {
		conn, err := net.Dial("tcp", clock.addr)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		go clock.watch(os.Stdout, conn)
	}
	// Sleep while other goroutines do the work.
	for {
		time.Sleep(time.Minute)
	}
}
