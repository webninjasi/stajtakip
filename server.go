package main

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

type StajServer struct {
	addr   string
	server *http.Server
	mux    *http.ServeMux
}

func (s *StajServer) SetHandlers() {
	s.mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("web"))))
}

func (s *StajServer) Run() error {
	logrus.WithFields(logrus.Fields{
		"addr": s.addr,
	}).Info("Sunucu başlatılıyor...")
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
