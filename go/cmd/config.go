package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"runtime"
	"strings"
)

func configure() {
	// Prepare logging formatter
	customFormatter := &logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return "", fmt.Sprintf(" %s:%d", formatFilePath(f.File), f.Line)
		},
	}
	logrus.SetFormatter(customFormatter)
	logrus.SetReportCaller(true)
}

func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	if len(arr) >= 2 {
		return fmt.Sprintf("%s/%s", arr[len(arr)-2], arr[len(arr)-1])
	}

	return arr[len(arr)-1]
}
