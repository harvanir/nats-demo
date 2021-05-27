package main

import (
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

func publisher() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		logrus.Fatal(err)
		return
	}
	defer nc.Close()
	// publish
	count := 0
	loop := 10
	msg := []byte("Hello")
	for i := 0; i < loop; i++ {
		count += 1
		publish(count, nc, subj, msg)
	}
}

func publish(i int, nc *nats.Conn, subj string, msg []byte) {
	err := nc.Publish(subj, msg)
	if err != nil {
		logrus.Error("error publishing: ", err)
		return
	}
	logrus.Infof("[#%d] Publishing [%s] : '%s'\n", i, subj, msg)
	err = nc.Flush()
	if err != nil {
		logrus.Error("error flushing: ", err)
		return
	}

	if err := nc.LastError(); err != nil {
		logrus.Fatal(err)
	} else {
		logrus.Infof("[#%d] Published [%s] : '%s'\n", i, subj, msg)
	}
}
