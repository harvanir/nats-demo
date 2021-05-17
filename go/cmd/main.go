package main

import (
	"log"
	"time"
)

const subj = "foo"

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	chanBool, closeFunc := subscriber()

	select {
	case boolVal, ok := <-chanBool:
		log.Printf("ok: %v", ok)
		publisher()
		time.Sleep(time.Second * 4)
		if ok {
			log.Printf("chan value: %v", boolVal)
			defer closeFunc()
		}
	}
}
