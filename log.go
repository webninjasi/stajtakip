package main

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

func LogBaslat(logpath string) {
	// Dosya belirtilmemişse sadece stdout'a yazdır
	if logpath == "" {
		logrus.SetOutput(os.Stdout)
		return
	}

	// Hem dosyaya hem stdout'a yazdır
	logfile, err := os.Create(logpath)
	if err != nil {
		logrus.SetOutput(os.Stdout)
		logrus.WithFields(logrus.Fields{
			"filename": logpath,
			"error":    err,
		}).Error("Log dosyası açılamadı!")
		return
	}

	confile := io.MultiWriter(os.Stdout, logfile)
	logrus.SetOutput(confile)
}
