package routes

import (
	"net/http"
	"stajtakip/database"
	"stajtakip/templates"

	"time"

	"github.com/sirupsen/logrus"
)

var tpl_mulakat_listesi = templates.Load("templates/mulakat-listesi.html")

type MulakatListesiVars struct {
	Mulakatlar []database.MulakatOgrenci
		Komisyon []database.Komisyon
}

type MulakatListesi struct {
	Conn *database.Connection
}

func (sh MulakatListesi) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost && r.Method != http.MethodGet {
		http.Error(w, "Geçersiz metod!", http.StatusNotFound)
		return
	}

	durum := http.StatusOK
	data := templates.NewMain("StajTakip - Mülakat Listesi")

	if r.Method == http.MethodPost {
		// Döngü amacında kullanılmayan for döngüsü
		// Tüm çıkışları break veya return ile kapalı
		for {
			if err := r.ParseForm(); err != nil {
				logrus.WithFields(logrus.Fields{
					"err": err,
				}).Warn("Mülakat formu okunamadı!")
				durum = http.StatusBadRequest
				data = data.Error("Hatalı istek!")
				break
			}

			var gorev string
			var err error

			gorev, err = formStr(r.PostFormValue("gorev"))
			if err != nil {
				http.Error(w, "Hatalı istek!", http.StatusBadRequest)
				return
			}

			if gorev == "yenile" {
				if err := database.MulakatListesiOlustur(sh.Conn); err != nil {
					logrus.WithFields(logrus.Fields{
						"err": err,
					}).Error("Mülakat listesi yenilenemedi!")
					durum = http.StatusInternalServerError
					data = data.Error("Mülakat listesi yenilenemedi!")
					break
				}

				data = data.Info("Mülakat listesi yenilendi!")
				break
			} else if gorev == "guncelle" {
				var no int
				var baslangic, tarih, saat, komisyon1, komisyon2 string

				no, err = formSayi(r.PostFormValue("no"))
				if err != nil {
					http.Error(w, "Hatalı istek!", http.StatusBadRequest)
					return
				}

				baslangic, err = formStr(r.PostFormValue("baslangic"))
				if err != nil {
					http.Error(w, "Hatalı istek!", http.StatusBadRequest)
					return
				}

				tarih, err = formStr(r.PostFormValue("tarih"))
				if err != nil {
					durum = http.StatusBadRequest
					data = data.Warning("Tarih eksik veya yanlış!")
					break
				}

				saat, err = formStr(r.PostFormValue("saat"))
				if err != nil {
					durum = http.StatusBadRequest
					data = data.Warning("Saat eksik veya yanlış!")
					break
				}

				komisyon1, err = formStr(r.PostFormValue("komisyon1"))
				if err != nil {
					durum = http.StatusBadRequest
					data = data.Warning("Komisyon üyesi (1) eksik veya yanlış!")
					break
				}

				komisyon2, err = formStr(r.PostFormValue("komisyon2"))
				if err != nil {
					durum = http.StatusBadRequest
					data = data.Warning("Komisyon üyesi (2) eksik veya yanlış!")
					break
				}

				var tarihSaat time.Time

				tarihSaat, err = time.Parse(database.TarihSaatFormati, tarih + " " + saat)
				if err != nil {
					durum = http.StatusBadRequest
					data = data.Warning("Tarih/saat formatı uygun değil!")
					break
				}

				mul := database.Mulakat{no, baslangic, tarihSaat.Format(database.TarihSaatFormati), komisyon1, komisyon2}
				if err := mul.Update(sh.Conn); err != nil {
					logrus.WithFields(logrus.Fields{
						"err": err,
					}).Error("Mülakat güncellenemedi!")
					durum = http.StatusInternalServerError
					data = data.Error("Mülakat güncellenemedi!")
					break
				}

				data = data.Info("Mülakat listesi güncellendi!")
				break
			}
			break
		}
	}

	vars := MulakatListesiVars{nil, nil}

	if mulakatlar, err := database.MulakatListesi(sh.Conn); err == nil {
		vars.Mulakatlar = mulakatlar
	} else {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("Mülakat listesi yüklenemedi!")
		w.WriteHeader(http.StatusInternalServerError)
		sablonHatasi(w, tpl_mulakat_listesi.ExecuteTemplate(w, "main", data.Error("Mülakat listesi yüklenemedi!")))
		return
	}

	if uyeler, err := database.KomisyonListesi(sh.Conn); err == nil {
		vars.Komisyon = uyeler
	} else {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("Komisyon listesi yüklenemedi!")
		w.WriteHeader(http.StatusInternalServerError)
		sablonHatasi(w, tpl_mulakat_listesi.ExecuteTemplate(w, "main", data.Error("Komisyon listesi yüklenemedi!")))
		return
	}

	data.Vars = vars

	w.WriteHeader(durum)
	sablonHatasi(w, tpl_mulakat_listesi.ExecuteTemplate(w, "main", data))
}
