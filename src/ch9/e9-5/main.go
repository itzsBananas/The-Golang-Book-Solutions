// Exercise 9.5: Write a program with two goroutines that send messages back and forth over
// two unbuffered channels in ping-pong fashion. How many communications per second can
// the program sustain?

// about 200,000 send/recieve operations

package main

import (
	"fmt"
	"os"
	"time"
)

func PingPong(ch1 chan int, ch2 chan int, done <-chan struct{}) {
	count := 0
	for {
		select {
		case <-done:
			return
		case ch2 <- count:
		case ch1 <- count:
		case c := <-ch1:
			count = c + 1
		case c := <-ch2:
			count = c + 1
		}
	}
}

func main() {
	done := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		close(done)
	}()
	for {
		var (
			ping  = make(chan int)
			pong  = make(chan int)
			reset = make(chan struct{})
		)
		go PingPong(ping, pong, reset)
		go PingPong(pong, ping, reset)
		select {
		case <-done:
			return
		case <-time.After(1 * time.Second):
			select {
			case res := <-ping:
				fmt.Println(res)
			case res := <-pong:
				fmt.Println(res)
			}
		}
		close(reset)
	}
}
