package routes

import (
	"net/http"
	"stajtakip/database"
	"stajtakip/templates"

	"time"

	"github.com/sirupsen/logrus"
)

var tpl_staj_ekle = templates.Load("templates/staj-ekle.html")

const zamanFormati = "2006-01-02"

type StajEkleVars struct {
	Konular []string
	Kurumlar []string
	DenkStaj bool
}

type StajEkle struct {
	Conn *database.Connection
}

// Verilen parametrelere göre veritabanına bir Staj eklemeye çalışır
func (sh StajEkle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := templates.NewMain("StajTakip - Öğrenci Ekle")

	if konular, err := database.KonuListesi(sh.Conn); err == nil {
		if kurumlar, err := database.KurumListesi(sh.Conn); err == nil {
			data.Vars = StajEkleVars{konular, kurumlar, false}
		} else {
			logrus.WithFields(logrus.Fields{
				"err": err,
			}).Error("Kurum listesi yüklenemedi!")
			w.WriteHeader(http.StatusInternalServerError)
			sablonHatasi(w, tpl_staj_ekle.ExecuteTemplate(w, "main", data.Error("Kurum listesi yüklenemedi!")))
			return
		}
	} else {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("Konu listesi yüklenemedi!")
		w.WriteHeader(http.StatusInternalServerError)
		sablonHatasi(w, tpl_staj_ekle.ExecuteTemplate(w, "main", data.Error("Konu listesi yüklenemedi!")))
		return
	}

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

	denk, err := formSayi(r.PostFormValue("denk"))
	if err != nil {
		http.Error(w, "Eksik bilgi!", http.StatusBadRequest)
		return
	}

	if denk == 0 {
		sh.NormalStajEkle(data, w, r)
	} else {
		sh.DenkStajEkle(data, w, r)
	}
}

func (sh StajEkle) NormalStajEkle(data templates.Main, w http.ResponseWriter, r *http.Request) {
	var ogrno, sinif, toplamgun int
	var kurum, sehir, konu, baslangic, bitis string
	var err error

	ogrno, err = formSayi(r.PostFormValue("no"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		sablonHatasi(w, tpl_staj_ekle.ExecuteTemplate(w, "main", data.Warning("Öğrenci no eksik veya yanlış!")))
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

	var t1, t2 time.Time

	baslangic, err = formStr(r.PostFormValue("baslangic"))
	if err == nil {
		t1, err = time.Parse(zamanFormati, baslangic)
	}
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		sablonHatasi(w, tpl_staj_ekle.ExecuteTemplate(w, "main", data.Warning("Başlangıç tarihi eksik veya yanlış!")))
		return
	}

	bitis, err = formStr(r.PostFormValue("bitis"))
	if err == nil {
		t2, err = time.Parse(zamanFormati, bitis)
	}
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		sablonHatasi(w, tpl_staj_ekle.ExecuteTemplate(w, "main", data.Warning("Bitiş tarihi eksik veya yanlış!")))
		return
	}

	if t2.Sub(t1).Hours() < 1 {
			w.WriteHeader(http.StatusBadRequest)
			sablonHatasi(w, tpl_staj_ekle.ExecuteTemplate(w, "main", data.Warning("Başlangıç tarihi, bitiş tarihinden önce olmalıdır!")))
			return
	}

	sinif, err = formSayi(r.PostFormValue("sinif"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		sablonHatasi(w, tpl_staj_ekle.ExecuteTemplate(w, "main", data.Warning("Sınıf eksik veya yanlış!")))
		return
	}

	toplamgun, err = formSayi(r.PostFormValue("toplamgun"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		sablonHatasi(w, tpl_staj_ekle.ExecuteTemplate(w, "main", data.Warning("Toplam gün eksik veya yanlış!")))
		return
	}

	ogr := database.Staj{ogrno, kurum, sehir, konu, baslangic, bitis, sinif, toplamgun, 0, false}
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

func (sh StajEkle) DenkStajEkle(data templates.Main, w http.ResponseWriter, r *http.Request) {
	var ogrno, toplamgun int
	var kurum, okul string
	var err error

	var dat StajEkleVars = data.Vars.(StajEkleVars)
	dat.DenkStaj = true
	data.Vars = dat

	ogrno, err = formSayi(r.PostFormValue("no"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		sablonHatasi(w, tpl_staj_ekle.ExecuteTemplate(w, "main", data.Warning("Öğrenci no eksik veya yanlış!")))
		return
	}

	kurum, err = formStr(r.PostFormValue("kurum"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		sablonHatasi(w, tpl_staj_ekle.ExecuteTemplate(w, "main", data.Warning("Kurum eksik veya yanlış!")))
		return
	}

	okul, err = formStr(r.PostFormValue("okul"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		sablonHatasi(w, tpl_staj_ekle.ExecuteTemplate(w, "main", data.Warning("Önceki okul eksik!")))
		return
	}

	toplamgun, err = formSayi(r.PostFormValue("toplamgun"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		sablonHatasi(w, tpl_staj_ekle.ExecuteTemplate(w, "main", data.Warning("Toplam gün eksik veya yanlış!")))
		return
	}

	ogr := database.DenkStaj{ogrno, kurum, okul, toplamgun, toplamgun/2}
	if err := ogr.Insert(sh.Conn); err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("DenkStaj eklenirken veritabanında bir hata oluştu!")
		w.WriteHeader(http.StatusInternalServerError)
		sablonHatasi(w, tpl_staj_ekle.ExecuteTemplate(w, "main", data.Error("Veritabanında bir hata oluştu!")))
		return
	}

	w.WriteHeader(http.StatusOK)
	sablonHatasi(w, tpl_staj_ekle.ExecuteTemplate(w, "main", data.Info("DGS/Yatay Geçiş Staj bilgisi veritabanına başarıyla eklendi!")))

	// TODO Eski değerleri inputlara ata
}
