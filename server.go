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
	s.mux.Handle("/", routes.Index)
	s.mux.Handle("/ogrenci-ekle", routes.OgrenciEkle{conn})
	s.mux.Handle("/staj-ekle", routes.StajEkle{conn})
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
