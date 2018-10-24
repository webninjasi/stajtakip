package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	srv := NewStajServer(os.Getenv("APP_SERVER_ADDR"))
	srv.SetHandlers()

	err := srv.Run()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("Sunucuyu çalıştırırken bir hata meydana geldi!")
	}
}
