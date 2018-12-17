package routes

import (
	"errors"
	"strconv"
	"strings"
	"net/http"

	"github.com/sirupsen/logrus"
)

var MaxReqSize int64 = 32000000

func Ayarla(mrs int64) {
	if mrs > 0 {
		MaxReqSize = mrs
	}
}

func formSayi(deger string) (int, error) {
	var sayi int

	str, err := formStr(deger)
	if err != nil {
		return 0, err
	}

	sayi, err = strconv.Atoi(str)
	if err != nil {
		return 0, err
	}

	return sayi, nil
}

func formStr(deger string) (string, error) {
	str := strings.TrimSpace(deger)

	if len(str) < 1 {
		return "", errors.New("Eksik değer!")
	}

	return str, nil
}

func sablonHatasi(w http.ResponseWriter, err error) {
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Warn("Şablon çalıştırılamadı!")
		http.Error(w, "Sayfa yüklenemedi!", http.StatusInternalServerError)
	}
}
