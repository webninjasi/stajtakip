package main

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

func init() {
	logpath := os.Getenv("APP_LOG_FILE")
	if logpath == "" {
		logrus.SetOutput(os.Stdout)
		return
	}

	logfile, err := os.Open(os.Getenv("APP_LOG_FILE"))
	if err != nil {
		logrus.SetOutput(os.Stdout)
		logrus.WithFields(logrus.Fields{
			"filename": os.Getenv("APP_LOG_FILE"),
			"error":    err,
		}).Error("Log dosyası açılamadı!")
		return
	}

	confile := io.MultiWriter(os.Stdout, logfile)
	logrus.SetOutput(confile)
}
