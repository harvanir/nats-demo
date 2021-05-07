package main

import (
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
	"log"
)

func publisher() *nats.Conn {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	msg := []byte("Hello")
	publish(nc, subj, msg)
	publish(nc, subj, msg)
	publish(nc, subj, msg)
	publish(nc, subj, msg)
	publish(nc, subj, msg)
	publish(nc, subj, msg)
	publish(nc, subj, msg)
	publish(nc, subj, msg)
	publish(nc, subj, msg)
	return nc
}

func publish(nc *nats.Conn, subj string, msg []byte) {
	err := nc.Publish(subj, msg)
	if err != nil {
		logrus.Error("error publishing: ", err)
		return
	}
	log.Printf("Publishing [%s] : '%s'\n", subj, msg)
	err = nc.Flush()
	if err != nil {
		logrus.Error("error flushing: ", err)
		return
	}

	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Published [%s] : '%s'\n", subj, msg)
	}
}
