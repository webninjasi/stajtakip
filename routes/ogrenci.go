package routes

import (
	"fmt"
	"net/http"
	"stajtakip/database"

	"github.com/sirupsen/logrus"
)

type OgrenciEkle struct {
	Conn *database.Connection
}

// Verilen parametrelere göre veritabanına bir öğrenci eklemeye çalışır
func (sh OgrenciEkle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Post metodu kullanılmalı!", http.StatusNotFound)
		return
	}

	if err := r.ParseForm(); err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Warn("Öğrenci ekleme formu okunamadı!")
		http.Error(w, "Öğrenci bilgileri formunda bir hata var!", http.StatusBadRequest)
		return
	}

	var no, ogretim int
	var ad, soyad string
	var err error

	ad, err = formStr(r.PostFormValue("ad"))
	if err != nil {
		http.Error(w, "Öğrenci adı eksik veya yanlış!", http.StatusBadRequest)
		return
	}

	soyad, err = formStr(r.PostFormValue("soyad"))
	if err != nil {
		http.Error(w, "Öğrenci soyadı eksik veya yanlış!", http.StatusBadRequest)
		return
	}

	no, err = formSayi(r.PostFormValue("no"))
	if err != nil {
		http.Error(w, "Öğrenci no eksik veya yanlış!", http.StatusBadRequest)
		return
	}

	ogretim, err = formSayi(r.PostFormValue("ogretim"))
	if err != nil {
		http.Error(w, "Öğretim eksik veya yanlış!", http.StatusBadRequest)
		return
	}

	ogr := database.Ogrenci{no, ad, soyad, ogretim}
	if err := ogr.Insert(sh.Conn); err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("Öğrenci eklenirken veritabanında bir hata oluştu!")
		http.Error(w, "Veritabanında bir hata oluştu!", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Öğrenci veritabanına başarıyla eklendi!")

	// TODO daha fazla field (isteğe bağlı olanlar vb.)
	// TODO fieldların max değerlerini vb. kontrol et
}

// TODO öğrenci bilgileri düzenleme
