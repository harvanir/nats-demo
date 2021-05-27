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
	"time"

	"github.com/nats-io/nats.go"
)

func printMsg(m *nats.Msg, i int) {
	logrus.Infof("[#%d] Received on [%s]: '%s'", i, m.Subject, string(m.Data))
}

//func subscriber(in chan string) <-chan *nats.Conn {
func subscriber() (chan bool, func()) {
	ncChan := make(chan bool, 1)
	var closeFunc func()

	go func() {
		//defer close(c)
		var urls = nats.DefaultURL
		opts := []nats.Option{nats.Name("NATS Sample Subscriber")}
		opts = setupConnOptions(opts)

		// Connect to NATS
		nc, err := nats.Connect(urls, opts...)
		if err != nil {
			logrus.Fatal(err)
			return
		}

		i := 0

		_, err = nc.Subscribe(subj, func(msg *nats.Msg) {
			i += 1
			go func(ii int) {
				time.Sleep(time.Second * 1)
				printMsg(msg, ii)
			}(i)
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
			logrus.Fatal(err)
			return
		}

		logrus.Infof("Listening on [%s]", subj)
		closeFunc = func() {
			logrus.Info("close function start...")
			nc.Close()
			logrus.Info("close function end...")
		}
		ncChan <- true
	}()
	return ncChan, func() {
		closeFunc()
		close(ncChan)
	}
}

func setupConnOptions(opts []nats.Option) []nats.Option {
	totalWait := 10 * time.Minute
	reconnectDelay := time.Second

	opts = append(opts, nats.ReconnectWait(reconnectDelay))
	opts = append(opts, nats.MaxReconnects(int(totalWait/reconnectDelay)))
	opts = append(opts, nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
		logrus.Infof("Disconnected due to: error(%v), will attempt reconnects for %.0fm", err, totalWait.Minutes())
	}))
	opts = append(opts, nats.ReconnectHandler(func(nc *nats.Conn) {
		logrus.Infof("Reconnected [%s]", nc.ConnectedUrl())
	}))
	opts = append(opts, nats.ClosedHandler(func(nc *nats.Conn) {
		logrus.Fatalf("Exiting with last error: %v", nc.LastError())
	}))
	return opts
}
