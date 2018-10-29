package main

import (
	"net/http"
)

type StajServer struct {
	addr   string
	server *http.Server
	mux    *http.ServeMux
}

func (s *StajServer) SetHandlers() {
	// Kök dizinde web klasörünü sun
	s.mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("web"))))
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
