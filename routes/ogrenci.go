package routes

import (
	"fmt"
	"net/http"
	"stajtakip/database"
	"strconv"
	"strings"

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

	var hangiOgretim int

	ogrno := strings.TrimSpace(r.PostFormValue("no"))
	ad := strings.TrimSpace(r.PostFormValue("ad"))
	soyad := strings.TrimSpace(r.PostFormValue("soyad"))
	ogretim := strings.TrimSpace(r.PostFormValue("ogretim"))

	if len(ad) < 1 || len(soyad) < 1 {
		http.Error(w, "Öğrenci adı veya soyadı eksik olmamalı!", http.StatusBadRequest)
		return
	}

	no, err := strconv.Atoi(ogrno)
	if err != nil {
		http.Error(w, "Öğrenci no bir sayı olmalı!", http.StatusBadRequest)
		return
	}

	hangiOgretim, err = strconv.Atoi(ogretim)
	if err != nil {
		http.Error(w, "Öğrenci no bir sayı olmalı!", http.StatusBadRequest)
		return
	}

	ogr := database.Ogrenci{no, ad, soyad, hangiOgretim}
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
