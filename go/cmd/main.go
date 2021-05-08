package main

import (
	"log"
)

const subj = "foo"

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)

	c := <-subscriber()
	defer c.Close()

	c2 := <-publisher()
	defer c2.Close()
}
