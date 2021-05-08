// Copyright 2012-2021 The NATS Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/sirupsen/logrus"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func printMsg(m *nats.Msg, i int) {
	log.Printf("[#%d] Received on [%s]: '%s'", i, m.Subject, string(m.Data))
}

func subscriber() <-chan *nats.Conn {
	ncChan := make(chan *nats.Conn, 1)

	go func(c chan *nats.Conn) {
		defer close(c)
		var urls = nats.DefaultURL
		opts := []nats.Option{nats.Name("NATS Sample Subscriber")}
		opts = setupConnOptions(opts)

		// Connect to NATS
		nc, err := nats.Connect(urls, opts...)
		if err != nil {
			log.Fatal(err)
			return
		}

		i := 0

		_, err = nc.Subscribe(subj, func(msg *nats.Msg) {
			i += 1
			printMsg(msg, i)
		})
		if err != nil {
			logrus.Error("error subscribe: ", err)
			return
		}
		err = nc.Flush()
		if err != nil {
			logrus.Error("error flushing: ", err)
			return
		}

		if err := nc.LastError(); err != nil {
			log.Fatal(err)
			return
		}

		log.Printf("Listening on [%s]", subj)
		c <- nc
	}(ncChan)
	return ncChan
}

func setupConnOptions(opts []nats.Option) []nats.Option {
	totalWait := 10 * time.Minute
	reconnectDelay := time.Second

	opts = append(opts, nats.ReconnectWait(reconnectDelay))
	opts = append(opts, nats.MaxReconnects(int(totalWait/reconnectDelay)))
	opts = append(opts, nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
		log.Printf("Disconnected due to: error(%s), will attempt reconnects for %.0fm", err, totalWait.Minutes())
	}))
	opts = append(opts, nats.ReconnectHandler(func(nc *nats.Conn) {
		log.Printf("Reconnected [%s]", nc.ConnectedUrl())
	}))
	opts = append(opts, nats.ClosedHandler(func(nc *nats.Conn) {
		log.Fatalf("Exiting with last error: %v", nc.LastError())
	}))
	return opts
}
