package routes

import (
	"net/http"
	"stajtakip/database"
	"stajtakip/templates"

	"github.com/sirupsen/logrus"
)

var tpl_raporlar = templates.Load("templates/raporlar.html")

type Raporlar struct {
	Conn *database.Connection
}

// Verilen parametrelere göre veritabanına bir öğrenci eklemeye çalışır
func (sh Raporlar) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Get metodu kullanılmalı!", http.StatusNotFound)
		return
	}

	err := tpl_raporlar.ExecuteTemplate(w, "main", templates.Main{"StajTakip - Raporlar"})
	if err != nil {
		http.Error(w, "Sayfa yüklenemedi!", http.StatusInternalServerError)
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Warn("Şablon çalıştırılamadı!")
	}
}
