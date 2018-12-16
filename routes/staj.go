package routes

import (
	"fmt"
	"net/http"
	"stajtakip/database"

	"github.com/sirupsen/logrus"
)

type StajEkle struct {
	Conn *database.Connection
}

// Verilen parametrelere göre veritabanına bir Staj eklemeye çalışır
func (sh StajEkle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Post metodu kullanılmalı!", http.StatusNotFound)
		return
	}

	if err := r.ParseForm(); err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Warn("Staj ekleme formu okunamadı!")
		http.Error(w, "Formda bir hata var!", http.StatusBadRequest)
		return
	}

	var ogrno, sinif int
	var kurum, sehir, konu, baslangic, bitis string
	var err error

	ogrno, err = formSayi(r.PostFormValue("no"))
	if err != nil {
		http.Error(w, "Öğrenci no eksik veya yanlış!", http.StatusBadRequest)
		return
	}

	sinif, err = formSayi(r.PostFormValue("sinif"))
	if err != nil {
		http.Error(w, "Sınıf eksik veya yanlış!", http.StatusBadRequest)
		return
	}

	kurum, err = formStr(r.PostFormValue("kurum"))
	if err != nil {
		http.Error(w, "Kurum eksik veya yanlış!", http.StatusBadRequest)
		return
	}

	sehir, err = formStr(r.PostFormValue("sehir"))
	if err != nil {
		http.Error(w, "Şehir eksik veya yanlış!", http.StatusBadRequest)
		return
	}

	konu, err = formStr(r.PostFormValue("konu"))
	if err != nil {
		http.Error(w, "Konu eksik veya yanlış!", http.StatusBadRequest)
		return
	}

	baslangic, err = formStr(r.PostFormValue("baslangic"))
	if err != nil {
		http.Error(w, "Başlangıç tarihi eksik veya yanlış!", http.StatusBadRequest)
		return
	}

	bitis, err = formStr(r.PostFormValue("bitis"))
	if err != nil {
		http.Error(w, "Bitiş tarihi eksik veya yanlış!", http.StatusBadRequest)
		return
	}

	ogr := database.Staj{ogrno, sinif, kurum, sehir, konu, baslangic, bitis}
	if err := ogr.Insert(sh.Conn); err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("Staj eklenirken veritabanında bir hata oluştu!")
		http.Error(w, "Veritabanında bir hata oluştu!", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Staj veritabanına başarıyla eklendi!")

	// TODO daha fazla field (isteğe bağlı olanlar vb.)
	// TODO fieldların max değerlerini vb. kontrol et
}

// TODO öğrenci bilgileri düzenleme

// TODO dgs öğrencileri için staj tablosunda sınıfı 0 olarak seç
