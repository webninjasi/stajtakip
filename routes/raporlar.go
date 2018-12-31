package routes

import (
	"net/http"
	"stajtakip/database"
	"stajtakip/templates"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

var tpl_404 = templates.Load("templates/404.html")
var tpl_raporlar = templates.Load("templates/raporlar.html")

type Raporlar struct {
	Conn *database.Connection
}

type RaporlarVars struct {
	KonuRapor  []database.RaporKonu
	SehirRapor []database.RaporSehir
}

func (sh Raporlar) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)

		err := tpl_404.ExecuteTemplate(w, "main", templates.NewMain("StajTakip - Sayfa Bulunamadı"))
		if err != nil {
			http.Error(w, "Sayfa yüklenemedi!", http.StatusInternalServerError)
			logrus.WithFields(logrus.Fields{
				"err": err,
			}).Warn("Şablon çalıştırılamadı!")
		}
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Get metodu kullanılmalı!", http.StatusNotFound)
		return
	}

	data := templates.NewMain("StajTakip - Raporlar")

	var year int
	var err error

	yearstr := r.FormValue("year")
	if yearstr == "" {
		year = time.Now().Year()
	} else {
		year, err = strconv.Atoi(yearstr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			sablonHatasi(w, tpl_raporlar.ExecuteTemplate(w, "main", data.Warning("Yıl eksik!")))
			return
		}
	}

	if sehirRapor, err := database.RaporSehirler(sh.Conn, year); err == nil {
		if konuRapor, err := database.RaporKonular(sh.Conn, year); err == nil {
			data.Vars = RaporlarVars{konuRapor, sehirRapor}
			sablonHatasi(w, tpl_raporlar.ExecuteTemplate(w, "main", data))
			return
		} else {
			logrus.WithFields(logrus.Fields{
				"err": err,
			}).Error("Konu raporu yüklenemedi!")
			w.WriteHeader(http.StatusInternalServerError)
			sablonHatasi(w, tpl_mesaj.ExecuteTemplate(w, "main", data.Error("Konu raporu listesi yüklenemedi!")))
			return
		}
	} else {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("Sehir raporu yüklenemedi!")
		w.WriteHeader(http.StatusInternalServerError)
		sablonHatasi(w, tpl_mesaj.ExecuteTemplate(w, "main", data.Error("Sehir raporu listesi yüklenemedi!")))
		return
	}

}
