package routes

import (
	"net/http"
	"stajtakip/database"
	"stajtakip/templates"

	"github.com/sirupsen/logrus"
)

var tpl_komisyon = templates.Load("templates/komisyon.html")

type KomisyonListesi struct {
	Conn *database.Connection
}

type KomisyonEkle struct {
	Conn *database.Connection
}

type KomisyonSil struct {
	Conn *database.Connection
}

type KomisyonGuncelle struct {
	Conn *database.Connection
}

type KomisyonVars struct {
	Uyeler []string
}

func (sh KomisyonListesi) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := templates.NewMain("StajTakip - Komisyon Listesi")

	if r.Method != http.MethodGet {
		http.Error(w, "Get metodu kullanılmalı!", http.StatusNotFound)
		return
	}

	if uyeler, err := database.KomisyonListesi(sh.Conn); err == nil {
		data.Vars = KomisyonVars{uyeler}
	} else {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("Komisyon listesi yüklenemedi!")
		w.WriteHeader(http.StatusInternalServerError)
		sablonHatasi(w, tpl_komisyon.ExecuteTemplate(w, "main", data.Error("Komisyon listesi yüklenemedi!")))
		return
	}

	w.WriteHeader(http.StatusOK)
	sablonHatasi(w, tpl_komisyon.ExecuteTemplate(w, "main", data))
}

func (sh KomisyonEkle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Post metodu kullanılmalı!", http.StatusNotFound)
		return
	}

	data := templates.NewMain("StajTakip - Komisyon Ekle")

	if err := r.ParseForm(); err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Warn("Komisyon ekleme formu okunamadı!")
		w.WriteHeader(http.StatusBadRequest)
		sablonHatasi(w, tpl_mesaj.ExecuteTemplate(w, "main", data.Error("Komisyon ekleme formu okunamadı!")))
		return
	}

	var adSoyad string
	var err error

	adSoyad, err = formStr(r.PostFormValue("adSoyad"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		sablonHatasi(w, tpl_mesaj.ExecuteTemplate(w, "main", data.Warning("AdSoyad eksik!")))
		return
	}

	kom := database.Komisyon{adSoyad, true}
	if err := kom.Insert(sh.Conn); err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("Komisyon üyesi eklenirken veritabanında bir hata oluştu!")
		w.WriteHeader(http.StatusInternalServerError)
		sablonHatasi(w, tpl_mesaj.ExecuteTemplate(w, "main", data.Error("Komisyon üyesi eklenirken veritabanında bir hata oluştu!")))
		return
	}

	w.WriteHeader(http.StatusOK)
	sablonHatasi(w, tpl_mesaj.ExecuteTemplate(w, "main", data.Info("Komisyon üyesi veritabanına başarıyla eklendi!")))
}
