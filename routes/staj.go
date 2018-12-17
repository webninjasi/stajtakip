package routes

import (
	"net/http"
	"stajtakip/database"
	"stajtakip/templates"

	"github.com/sirupsen/logrus"
)

var tpl_staj_ekle = templates.Load("templates/staj-ekle.html")

type StajEkle struct {
	Conn *database.Connection
}

// Verilen parametrelere göre veritabanına bir Staj eklemeye çalışır
func (sh StajEkle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := templates.NewMain("StajTakip - Öğrenci Ekle")

	if r.Method == http.MethodGet {
		sablonHatasi(w, tpl_staj_ekle.ExecuteTemplate(w, "main", data))
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Post metodu kullanılmalı!", http.StatusNotFound)
		return
	}

	if err := r.ParseForm(); err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Warn("Staj ekleme formu okunamadı!")
		w.WriteHeader(http.StatusBadRequest)
		sablonHatasi(w, tpl_staj_ekle.ExecuteTemplate(w, "main", data.Error("Staj ekleme formunda bir hata var!")))
		return
	}

	var ogrno, sinif int
	var kurum, sehir, konu, baslangic, bitis string
	var err error

	ogrno, err = formSayi(r.PostFormValue("no"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		sablonHatasi(w, tpl_staj_ekle.ExecuteTemplate(w, "main", data.Warning("Öğrenci no eksik veya yanlış!")))
		return
	}

	sinif, err = formSayi(r.PostFormValue("sinif"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		sablonHatasi(w, tpl_staj_ekle.ExecuteTemplate(w, "main", data.Warning("Sınıf eksik veya yanlış!")))
		return
	}

	kurum, err = formStr(r.PostFormValue("kurum"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		sablonHatasi(w, tpl_staj_ekle.ExecuteTemplate(w, "main", data.Warning("Kurum eksik veya yanlış!")))
		return
	}

	sehir, err = formStr(r.PostFormValue("sehir"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		sablonHatasi(w, tpl_staj_ekle.ExecuteTemplate(w, "main", data.Warning("Şehir eksik veya yanlış!")))
		return
	}

	konu, err = formStr(r.PostFormValue("konu"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		sablonHatasi(w, tpl_staj_ekle.ExecuteTemplate(w, "main", data.Warning("Konu eksik veya yanlış!")))
		return
	}

	baslangic, err = formStr(r.PostFormValue("baslangic"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		sablonHatasi(w, tpl_staj_ekle.ExecuteTemplate(w, "main", data.Warning("Başlangıç tarihi eksik veya yanlış!")))
		return
	}

	bitis, err = formStr(r.PostFormValue("bitis"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		sablonHatasi(w, tpl_staj_ekle.ExecuteTemplate(w, "main", data.Warning("Bitiş tarihi eksik veya yanlış!")))
		return
	}

	ogr := database.Staj{ogrno, sinif, kurum, sehir, konu, baslangic, bitis}
	if err := ogr.Insert(sh.Conn); err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("Staj eklenirken veritabanında bir hata oluştu!")
		w.WriteHeader(http.StatusInternalServerError)
		sablonHatasi(w, tpl_staj_ekle.ExecuteTemplate(w, "main", data.Error("Veritabanında bir hata oluştu!")))
		return
	}

	w.WriteHeader(http.StatusOK)
	sablonHatasi(w, tpl_staj_ekle.ExecuteTemplate(w, "main", data.Info("Staj bilgisi veritabanına başarıyla eklendi!")))

	// TODO Eski değerleri inputlara ata
}

// TODO dgs öğrencileri için staj tablosunda sınıfı 0 olarak seç
