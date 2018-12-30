package routes

import (
	"net/http"
	
	"stajtakip/database"
	"stajtakip/templates"
	"stajtakip/cfg"

	"github.com/sirupsen/logrus"
)

var tpl_ogrenci_bilgi = templates.Load("templates/ogrenci-bilgi.html")

type OgrenciBilgiVars struct {
	Ogr *database.Ogrenci
  Stajlar []database.Staj
	Basari bool
}

type OgrenciBilgi struct {
	Conn *database.Connection
}

func (sh OgrenciBilgi) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Get metodu kullanılmalı!", http.StatusNotFound)
		return
	}

  data := templates.NewMain("StajTakip - Öğrenci Bilgileri")
  data.Vars = OgrenciBilgiVars{
    Ogr: nil,
  }

	if err := r.ParseForm(); err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Warn("Öğrenci bilgi sayfası parametreleri okunamadı!")
		w.WriteHeader(http.StatusBadRequest)
		sablonHatasi(w, tpl_ogrenci_bilgi.ExecuteTemplate(w, "main", data.Error("Parametre hatası!")))
		return
	}

	var no int
	var err error
  var ogr *database.Ogrenci
  var stjlar []database.Staj

	no, err = formSayi(r.FormValue("no"))
	if err != nil || (no < 0) {
		w.WriteHeader(http.StatusBadRequest)
		sablonHatasi(w, tpl_ogrenci_bilgi.ExecuteTemplate(w, "main", data.Warning("Öğrenci no eksik veya yanlış!")))
		return
	}

  ogr, err = database.OgrenciBul(sh.Conn, no)
  if err == database.Err_ogrenci_yok {
    w.WriteHeader(http.StatusBadRequest)
    sablonHatasi(w, tpl_ogrenci_bilgi.ExecuteTemplate(w, "main", data.Error("Öğrenci bulunamadı!")))
    return
  } else if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("Öğrenci bilgileri aranırken veritabanında bir hata oluştu!")
		w.WriteHeader(http.StatusInternalServerError)
		sablonHatasi(w, tpl_ogrenci_bilgi.ExecuteTemplate(w, "main", data.Error("Veritabanında bir hata oluştu!")))
		return
	}

  stjlar, err = database.OgrenciStajListesi(sh.Conn, no)
  if err != nil {
    logrus.WithFields(logrus.Fields{
      "err": err,
    }).Error("Öğrenci staj bilgileri aranırken veritabanında bir hata oluştu!")
    w.WriteHeader(http.StatusInternalServerError)
    sablonHatasi(w, tpl_ogrenci_bilgi.ExecuteTemplate(w, "main", data.Error("Veritabanında bir hata oluştu!")))
    return
  }

	var kabulGun, toplamGun int = 0, 0
	var basari bool = false

	for _, stj := range stjlar {
		kabulGun += stj.KabulGun
		toplamGun += stj.ToplamGun
	}

	if kabulGun >= cfg.GerekenStajGunu() && toplamGun >= 60 {
		basari = true
	}

  data.Vars = OgrenciBilgiVars{
    Ogr: ogr,
    Stajlar: stjlar,
		Basari: basari,
  }

	w.WriteHeader(http.StatusOK)
	sablonHatasi(w, tpl_ogrenci_bilgi.ExecuteTemplate(w, "main", data))
}
