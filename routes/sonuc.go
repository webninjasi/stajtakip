package routes

import (
  "net/http"
  "time"
  "stajtakip/database"
  "stajtakip/templates"

  "github.com/sirupsen/logrus"
)

var tpl_sonuclar = templates.Load("templates/mulakat-sonuc.html")

type SonucListeleVars struct {
  Mul []database.MulakatOgrenciStaj
}

type SonucListele struct {
	Conn *database.Connection
}

func (sh SonucListele) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodGet {
  	http.Error(w, "Geçersiz metod!", http.StatusNotFound)
  	return
  }

  data := templates.NewMain("StajTakip - Mülakat Sonuçları")
  vars := SonucListeleVars{
    Mul: nil,
  }
  data.Vars = vars

  if err := r.ParseForm(); err != nil {
    logrus.WithFields(logrus.Fields{
      "err": err,
    }).Warn("Sayfa parametreleri okunamadı!")
    w.WriteHeader(http.StatusBadRequest)
    sablonHatasi(w, tpl_sonuclar.ExecuteTemplate(w, "main", data.Error("Parametre hatası!")))
    return
  }

  var baslangic, bitis string
  var err error

  bitisstr := r.FormValue("bitis")
  baslangicstr := r.FormValue("baslangic")

  if bitisstr == "" || baslangicstr == "" {
    data.Vars = vars
  	w.WriteHeader(http.StatusOK)
  	sablonHatasi(w, tpl_sonuclar.ExecuteTemplate(w, "main", data))
    return
  }

  baslangic, err = formStr(baslangicstr)
  if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		sablonHatasi(w, tpl_sonuclar.ExecuteTemplate(w, "main", data.Warning("Başlangıç tarihi eksik veya yanlış!")))
    return
  }

  bitis, err = formStr(bitisstr)
  if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		sablonHatasi(w, tpl_sonuclar.ExecuteTemplate(w, "main", data.Warning("Başlangıç tarihi eksik veya yanlış!")))
    return
  }

  bas, err := time.Parse(database.TarihFormati, baslangic)
  if err != nil {
  	http.Error(w, "Hatalı tarih formatı!", http.StatusBadRequest)
    return
  }

  son, err := time.Parse(database.TarihFormati, bitis)
  if err != nil {
  	http.Error(w, "Hatalı tarih formatı!", http.StatusBadRequest)
    return
  }

  if bas.After(son) {
		w.WriteHeader(http.StatusBadRequest)
		sablonHatasi(w, tpl_sonuclar.ExecuteTemplate(w, "main", data.Warning("Başlangıç tarihi, bitiş tarihinden sonra olamaz!")))
    return
  }

  vars.Mul, err = database.MulakatSonucListesi(sh.Conn, baslangic, bitis)
  if err == database.ErrVeriBulunamadi {
    w.WriteHeader(http.StatusBadRequest)
    sablonHatasi(w, tpl_sonuclar.ExecuteTemplate(w, "main", data.Warning("Mülakat sonucu bulunamadı!")))
    return
  } else if err != nil {
    logrus.WithFields(logrus.Fields{
      "err": err,
    }).Error("Mülakat aranırken veritabanında bir hata oluştu!")
    w.WriteHeader(http.StatusInternalServerError)
    sablonHatasi(w, tpl_sonuclar.ExecuteTemplate(w, "main", data.Error("Veritabanında bir hata oluştu!")))
    return
  }

	w.WriteHeader(http.StatusOK)
	sablonHatasi(w, tpl_sonuclar.ExecuteTemplate(w, "main", data))
}
