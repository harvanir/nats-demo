package main

import (
	"log"
	"time"
)

const subj = "foo"

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	chanConn := subscriber()

	select {
	case conSub, ok := <-chanConn:
		log.Printf("ok: %v", ok)
		publisher()
		time.Sleep(time.Second * 4)
		if ok {
			defer func() {
				log.Printf("closing chanConn...")
				conSub.Close()
				close(chanConn)
			}()
		}
	}
}
