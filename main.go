package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	srvAddr := os.Getenv("APP_SERVER_ADDR")
	srv := NewStajServer(srvAddr)
	srv.SetHandlers()

	datasrc := os.Getenv("APP_DATA_SOURCE")
	if datasrc == "" {
		logrus.Error("Veritabanı adresi belirtilmemiş! (APP_DATA_SOURCE)")
		return
	}

	db := NewStajVeritabani(datasrc)
	dbOk := make(chan bool)
	errChan := make(chan error)

	go func() {
		logrus.WithFields(logrus.Fields{
			"data-source": datasrc,
		}).Info("Veritabanına bağlanılıyor...")

		// Veritabanına bağlantıyı sağla
		if err := db.ConnectAndExecute(dbOk); err != nil {
			logrus.WithFields(logrus.Fields{
				"err": err,
			}).Error("Veritabanı bağlantısında bir hata meydana geldi!")
			errChan <- err
		}
	}()

	go func() {
		<-dbOk // Veritabanı bağlantısını bekle

		logrus.WithFields(logrus.Fields{
			"addr": srvAddr,
		}).Info("Sunucu başlatılıyor...")

		// Sunucuyu başlat
		if err := srv.Run(); err != nil {
			logrus.WithFields(logrus.Fields{
				"err": err,
			}).Error("Sunucuyu çalıştırırken bir hata meydana geldi!")
			errChan <- err
		}
	}()

	<-errChan // Bir hata oluşmasını bekle
}
