package main

import (
	"github.com/sirupsen/logrus"
	"time"
)

const subj = "foo"

func main() {
	configure()
	chanBool, closeFunc := subscriber()

	select {
	case boolVal, ok := <-chanBool:
		logrus.Infof("ok: %v", ok)
		publisher()
		time.Sleep(time.Second * 4)
		if ok {
			logrus.Infof("chan value: %v", boolVal)
			defer closeFunc()
		}
	}
}
