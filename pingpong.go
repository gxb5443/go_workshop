package main

import (
	"fmt"
	"time"
)

type Ball struct {
	Hits int
}

func main() {
	table := make(chan *Ball)
	go Player("Ping", table)
	go Player("Pong", table)

	table <- new(Ball)
	time.Sleep(time.Second)
	close(table)
}

func Player(name string, table chan *Ball) {
	for {
		ball := <-table
		ball.Hits++
		fmt.Println(name, ball.Hits)
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}
