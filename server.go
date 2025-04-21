package firewall

import "net/http"

type Server struct {
	Allow   func(r *http.Request) bool
	Handler http.Handler
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !s.Allow(r) {
		http.NotFound(w, r)
		return
	}
	s.Handler.ServeHTTP(w, r)
}
