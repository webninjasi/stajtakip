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

func (s *StajServer) SetHandlers(db *database.StajDatabase) {
	s.mux.Handle("/", routes.Index)
	s.mux.Handle("/ogrenci-ekle", routes.OgrenciEkle{db})
	s.mux.Handle("/staj-ekle", routes.StajEkle{db})
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
