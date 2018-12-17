package routes

import (
	"net/http"
	"stajtakip/database"
	"stajtakip/templates"

	"github.com/sirupsen/logrus"
)

var tpl_ogrenci_ekle = templates.Load("templates/ogrenci-ekle.html")

type OgrenciEkle struct {
	Conn *database.Connection
}

// Verilen parametrelere göre veritabanına bir öğrenci eklemeye çalışır
func (sh OgrenciEkle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := templates.NewMain("StajTakip - Öğrenci Ekle")

	if r.Method == http.MethodGet {
		sablonHatasi(w, tpl_ogrenci_ekle.ExecuteTemplate(w, "main", data))
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Post metodu kullanılmalı!", http.StatusNotFound)
		return
	}

	if err := r.ParseForm(); err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Warn("Öğrenci ekleme formu okunamadı!")
		w.WriteHeader(http.StatusBadRequest)
		sablonHatasi(w, tpl_ogrenci_ekle.ExecuteTemplate(w, "main", data.Error("Öğrenci bilgileri formunda bir hata var!")))
		return
	}

	var no, ogretim int
	var ad, soyad string
	var err error

	ad, err = formStr(r.PostFormValue("ad"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		sablonHatasi(w, tpl_ogrenci_ekle.ExecuteTemplate(w, "main", data.Warning("Öğrenci adı eksik veya yanlış!")))
		return
	}

	soyad, err = formStr(r.PostFormValue("soyad"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		sablonHatasi(w, tpl_ogrenci_ekle.ExecuteTemplate(w, "main", data.Warning("Öğrenci soyadı eksik veya yanlış!")))
		return
	}

	no, err = formSayi(r.PostFormValue("no"))
	if err != nil || (no < 0) {
		w.WriteHeader(http.StatusBadRequest)
		sablonHatasi(w, tpl_ogrenci_ekle.ExecuteTemplate(w, "main", data.Warning("Öğrenci no eksik veya yanlış!")))
		return
	}

	ogretim, err = formSayi(r.PostFormValue("ogretim"))
	if err != nil || (ogretim != 0 && ogretim != 1) {
		w.WriteHeader(http.StatusBadRequest)
		sablonHatasi(w, tpl_ogrenci_ekle.ExecuteTemplate(w, "main", data.Warning("Öğretim eksik veya yanlış!")))
		return
	}

	ogr := database.Ogrenci{no, ad, soyad, ogretim}
	if err := ogr.Insert(sh.Conn); err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("Öğrenci eklenirken veritabanında bir hata oluştu!")
		w.WriteHeader(http.StatusInternalServerError)
		sablonHatasi(w, tpl_ogrenci_ekle.ExecuteTemplate(w, "main", data.Error("Veritabanında bir hata oluştu!")))
		return
	}

	w.WriteHeader(http.StatusOK)
	sablonHatasi(w, tpl_ogrenci_ekle.ExecuteTemplate(w, "main", data.Info("Öğrenci veritabanına başarıyla eklendi!")))
	// TODO eklenen öğrencinin detay bilgisine giden link ekle
}
