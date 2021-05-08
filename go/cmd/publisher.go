package main

import (
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
	"log"
)

func publisher() <-chan *nats.Conn {
	ncChan := make(chan *nats.Conn)

	go func(c chan *nats.Conn) {
		defer close(c)

		nc, err := nats.Connect(nats.DefaultURL)
		if err != nil {
			log.Fatal(err)
			return
		}

		i := 0
		msg := []byte("Hello")
		i += 1
		publish(i, nc, subj, msg)
		i += 1
		publish(i, nc, subj, msg)
		i += 1
		publish(i, nc, subj, msg)
		i += 1
		publish(i, nc, subj, msg)
		i += 1
		publish(i, nc, subj, msg)
		i += 1
		publish(i, nc, subj, msg)
		i += 1
		publish(i, nc, subj, msg)
		i += 1
		publish(i, nc, subj, msg)
		i += 1
		publish(i, nc, subj, msg)
		c <- nc
	}(ncChan)
	return ncChan
}

func publish(i int, nc *nats.Conn, subj string, msg []byte) {
	err := nc.Publish(subj, msg)
	if err != nil {
		logrus.Error("error publishing: ", err)
		return
	}
	log.Printf("[#%d] Publishing [%s] : '%s'\n", i, subj, msg)
	err = nc.Flush()
	if err != nil {
		logrus.Error("error flushing: ", err)
		return
	}

	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("[#%d] Published [%s] : '%s'\n", i, subj, msg)
	}
}
