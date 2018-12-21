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

type KonularVars struct {
	Konular []string
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
