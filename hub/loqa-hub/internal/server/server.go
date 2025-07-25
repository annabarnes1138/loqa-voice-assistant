package server

import (
	"fmt"
	"log"
	"net/http"
)

type Config struct {
	Port      string
	ASRURL    string
	IntentURL string
	TTSURL    string
}

type Server struct {
	cfg Config
	mux *http.ServeMux
}

func New(cfg Config) *Server {
	mux := http.NewServeMux()

	s := &Server{cfg: cfg, mux: mux}
	s.routes()
	return s
}

func (s *Server) Start() error {
	return http.ListenAndServe(":"+s.cfg.Port, s.mux)
}

func (s *Server) routes() {
	s.mux.HandleFunc("/health", s.handleHealth)
	// future: /wake, /stream, /session, etc.
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	log.Println("Health check received")
	fmt.Fprintln(w, "ok")
}
