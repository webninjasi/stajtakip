package routes

import (
	"net/http"
	"stajtakip/database"
	"stajtakip/templates"

	"github.com/sirupsen/logrus"
)

var tpl_konu_listesi = templates.Load("templates/konular.html")

type KonuListesi struct {
	Conn *database.Connection
}

type KonuEkle struct {
	Conn *database.Connection
}

type KonuSil struct {
	Conn *database.Connection
}

type KonuGuncelle struct {
	Conn *database.Connection
}

type KonularVars struct {
	Konular []database.Konu
}

func (sh KonuListesi) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := templates.NewMain("StajTakip - Konu Listesi")

	if r.Method != http.MethodGet {
		http.Error(w, "Get metodu kullanılmalı!", http.StatusNotFound)
		return
	}

	if konular, err := database.KonuListesi(sh.Conn); err == nil {
		data.Vars = KonularVars{konular}
	} else {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("Konu listesi yüklenemedi!")
		w.WriteHeader(http.StatusInternalServerError)
		sablonHatasi(w, tpl_konu_listesi.ExecuteTemplate(w, "main", data.Error("Konu listesi yüklenemedi!")))
		return
	}

	w.WriteHeader(http.StatusOK)
	sablonHatasi(w, tpl_konu_listesi.ExecuteTemplate(w, "main", data))
}

func (sh KonuEkle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Post metodu kullanılmalı!", http.StatusNotFound)
		return
	}

	data := templates.NewMain("StajTakip - Konu Ekle")

	if err := r.ParseForm(); err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Warn("Konu ekleme formu okunamadı!")
		w.WriteHeader(http.StatusBadRequest)
		sablonHatasi(w, tpl_mesaj.ExecuteTemplate(w, "main", data.Error("Konu ekleme formu okunamadı!")))
		return
	}

	var baslik string
	var err error

	baslik, err = formStr(r.PostFormValue("Baslik"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		sablonHatasi(w, tpl_mesaj.ExecuteTemplate(w, "main", data.Warning("Baslik eksik!")))
		return
	}

	konu := database.Konu{baslik, true}
	if err := konu.Insert(sh.Conn); err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("Konu eklenirken veritabanında bir hata oluştu!")
		w.WriteHeader(http.StatusInternalServerError)
		sablonHatasi(w, tpl_mesaj.ExecuteTemplate(w, "main", data.Error("Konu eklenirken veritabanında bir hata oluştu!")))
		return
	}

	w.WriteHeader(http.StatusOK)
	sablonHatasi(w, tpl_mesaj.ExecuteTemplate(w, "main", data.Info("Konu veritabanına başarıyla eklendi!")))
}

func (sh KonuGuncelle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Post metodu kullanılmalı!", http.StatusNotFound)
		return
	}

	data := templates.NewMain("StajTakip - Konu Sil")

	if err := r.ParseForm(); err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Warn("Konu silme formu okunamadı!")
		w.WriteHeader(http.StatusBadRequest)
		sablonHatasi(w, tpl_mesaj.ExecuteTemplate(w, "main", data.Error("Konu silme formu okunamadı!")))
		return
	}

	var baslik string
	var aktifMi string

	baslik = r.PostFormValue("baslik")

	aktifMi = r.PostFormValue("aktif")
	if aktifMi != "" {
		konu := database.Konu{baslik, true}
		if err := konu.Update(sh.Conn); err != nil {
			logrus.WithFields(logrus.Fields{
				"err": err,
			}).Error("Konu eklenirken veritabanında bir hata oluştu!")
			w.WriteHeader(http.StatusInternalServerError)
			sablonHatasi(w, tpl_mesaj.ExecuteTemplate(w, "main", data.Error("Konu eklenirken veritabanında bir hata oluştu!")))
			return
		}
	} else {
		konu := database.Konu{baslik, false}
		if err := konu.Update(sh.Conn); err != nil {
			logrus.WithFields(logrus.Fields{
				"err": err,
			}).Error("Konu eklenirken veritabanında bir hata oluştu!")
			w.WriteHeader(http.StatusInternalServerError)
			sablonHatasi(w, tpl_mesaj.ExecuteTemplate(w, "main", data.Error("Konu eklenirken veritabanında bir hata oluştu!")))
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	sablonHatasi(w, tpl_mesaj.ExecuteTemplate(w, "main", data.Info("Konu veritabanından başarıyla değiştirildi.")))
}

func (sh KonuSil) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Post metodu kullanılmalı!", http.StatusNotFound)
		return
	}

	data := templates.NewMain("StajTakip - Konu Sil")

	if err := r.ParseForm(); err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Warn("Konu silme formu okunamadı!")
		w.WriteHeader(http.StatusBadRequest)
		sablonHatasi(w, tpl_mesaj.ExecuteTemplate(w, "main", data.Error("Konu silme formu okunamadı!")))
		return
	}

	var baslik string

	baslik = r.PostFormValue("baslik")

	konu := database.Konu{baslik, true}
	if err := konu.Delete(sh.Conn); err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("Konu eklenirken veritabanında bir hata oluştu!")
		w.WriteHeader(http.StatusInternalServerError)
		sablonHatasi(w, tpl_mesaj.ExecuteTemplate(w, "main", data.Error("Konu eklenirken veritabanında bir hata oluştu!")))
		return
	}

	w.WriteHeader(http.StatusOK)
	sablonHatasi(w, tpl_mesaj.ExecuteTemplate(w, "main", data.Info("Konu veritabanından başarıyla değiştirildi.")))
}
