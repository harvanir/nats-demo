package main

import (
	"github.com/nats-io/nats.go"
	"log"
)

const subj = "foo"

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)

	ncChan := make(chan *nats.Conn, 1)
	go func() {
		ncChan <- subscriber()
	}()
	nc := <-ncChan
	defer nc.Close()

	ncChan2 := make(chan *nats.Conn, 1)
	go func() {
		ncChan2 <- publisher()
	}()
	nc2 := <-ncChan2
	defer nc2.Close()
}
