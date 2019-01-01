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
	Uyeler []database.Komisyon
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

func (sh KomisyonGuncelle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Post metodu kullanılmalı!", http.StatusNotFound)
		return
	}

	data := templates.NewMain("StajTakip - Komisyon Sil")

	if err := r.ParseForm(); err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Warn("Komisyon güncelleme formu okunamadı!")
		w.WriteHeader(http.StatusBadRequest)
		sablonHatasi(w, tpl_mesaj.ExecuteTemplate(w, "main", data.Error("Komisyon güncelleme formu okunamadı!")))
		return
	}

	var adSoyad string
	var dahilMi string
	var err error

	adSoyad, err = formStr(r.PostFormValue("adSoyad"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		sablonHatasi(w, tpl_mesaj.ExecuteTemplate(w, "main", data.Warning("AdSoyad eksik!")))
		return
	}
	dahilMi = r.PostFormValue("dahil")

	if dahilMi != "" {
		kom := database.Komisyon{adSoyad, true}
		if err := kom.Update(sh.Conn); err != nil {
			logrus.WithFields(logrus.Fields{
				"err": err,
			}).Error("Komisyon üyesi güncellenirken veritabanında bir hata oluştu!")
			w.WriteHeader(http.StatusInternalServerError)
			sablonHatasi(w, tpl_mesaj.ExecuteTemplate(w, "main", data.Error("Komisyon güncellenirken silinirken veritabanında bir hata oluştu!")))
			return
		}
	} else {
		kom := database.Komisyon{adSoyad, false}
		if err := kom.Update(sh.Conn); err != nil {
			logrus.WithFields(logrus.Fields{
				"err": err,
			}).Error("Komisyon üyesi güncellenirken veritabanında bir hata oluştu!")
			w.WriteHeader(http.StatusInternalServerError)
			sablonHatasi(w, tpl_mesaj.ExecuteTemplate(w, "main", data.Error("Komisyon üyesi güncellenirken veritabanında bir hata oluştu!")))
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	sablonHatasi(w, tpl_mesaj.ExecuteTemplate(w, "main", data.Info("Komisyon üyesi veritabanında başarıyla güncellendi.")))
}

func (sh KomisyonSil) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Post metodu kullanılmalı!", http.StatusNotFound)
		return
	}

	data := templates.NewMain("StajTakip - Komisyon Guncelle")

	if err := r.ParseForm(); err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Warn("Komisyon silme formu okunamadı!")
		w.WriteHeader(http.StatusBadRequest)
		sablonHatasi(w, tpl_mesaj.ExecuteTemplate(w, "main", data.Error("Komisyon silme formu okunamadı!")))
		return
	}

	var adSoyad string

	adSoyad = r.PostFormValue("adSoyad")

	kom := database.Komisyon{adSoyad, true}
	if err := kom.Delete(sh.Conn); err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("Komisyon üyesi silinirken veritabanında bir hata oluştu!")
		w.WriteHeader(http.StatusInternalServerError)
		sablonHatasi(w, tpl_mesaj.ExecuteTemplate(w, "main", data.Error("Komisyon üyesi silinirken veritabanında bir hata oluştu!")))
		return
	}

	w.WriteHeader(http.StatusOK)
	sablonHatasi(w, tpl_mesaj.ExecuteTemplate(w, "main", data.Info("Komisyon üyesi veritabanından başarıyla silindi.")))
}
