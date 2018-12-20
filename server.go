package main

import (
	"net/http"
	"stajtakip/database"
	"stajtakip/routes"
)

type StajServer struct {
	addr   string
	server *http.Server
	mux    *http.ServeMux
}

func (s *StajServer) SetHandlers(conn *database.Connection) {
	// TODO giriş yapmadan erişilemesin?

	s.mux.Handle("/assets/", routes.Assets)

	// Raporlar
	s.mux.Handle("/", routes.Raporlar{conn})

	// Öğrenci
	s.mux.Handle("/ogrenci-ekle", routes.OgrenciEkle{conn})
	s.mux.Handle("/ogrenci-listele", routes.OgrenciListesi{conn}) // Stajı bitenler
	s.mux.Handle("/ogrenci-listele-pdf", routes.TODO{conn})
	s.mux.Handle("/ogrenci-listele-xls", routes.TODO{conn})
	s.mux.Handle("/ogrenci-belge-ekle", routes.TODO{conn}) // DGS için pdf
	s.mux.Handle("/ogrenci-ara", routes.TODO{conn})        // OgrNo -> Bilgiler, stajlar, staj bitim durumu

	// Staj
	s.mux.Handle("/staj-ekle", routes.StajEkle{conn})
	s.mux.Handle("/konu-listele", routes.TODO{conn})     // eklerken ajax ile çek?
	s.mux.Handle("/kurum-listele", routes.TODO{conn})    // eklerken ajax ile çek?
	s.mux.Handle("/onceki-staj-ekle", routes.TODO{conn}) // DGS/Yatay Geçiş
	s.mux.Handle("/konular", routes.TODO{conn})
	s.mux.Handle("/konu-ekle", routes.TODO{conn})
	s.mux.Handle("/konu-sil", routes.TODO{conn})
	s.mux.Handle("/konu-guncelle", routes.TODO{conn})

	// Mülakat
	s.mux.Handle("/mulakat-listele", routes.TODO{conn}) // Tarih/saat, komisyon -> öğrenci, ...
	s.mux.Handle("/sonuc-listele", routes.TODO{conn})   // Mülakat sonuçları
	s.mux.Handle("/sonuc-listele-pdf", routes.TODO{conn})
	s.mux.Handle("/sonuc-listele-xls", routes.TODO{conn})
	s.mux.Handle("/staj-degerlendir", routes.TODO{conn}) // Mülakat sonucu ekle
	s.mux.Handle("/komisyon", routes.TODO{conn})
	s.mux.Handle("/komisyon-ekle", routes.TODO{conn})
	s.mux.Handle("/komisyon-cikar", routes.TODO{conn})
	s.mux.Handle("/komisyon-ata", routes.TODO{conn})

	// Rapor
	s.mux.Handle("/rapor-il", routes.TODO{conn})   // İl bazında başarı
	s.mux.Handle("/rapor-konu", routes.TODO{conn}) // Konu bazında başarı/dağılım (yıllık)
}

func (s *StajServer) Run() error {
	return s.server.ListenAndServe()
}

func NewStajServer(addr string) *StajServer {
	mux := http.NewServeMux()
	server := &http.Server{Addr: addr, Handler: mux}

	return &StajServer{
		addr:   addr,
		server: server,
		mux:    mux,
	}
}
