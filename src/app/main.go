package app

import (
	"net/http"
	"time"
)

type server struct {
	startedAt time.Time
	router    *http.ServeMux
}

func (s *server) addHandler(path string, handler func(http.ResponseWriter, *http.Request)) {
	s.router.HandleFunc(path, handler)
	if len(path) < 2 {
		return
	}
	if path[len(path)-1] == '/' {
		s.router.HandleFunc(path[:len(path)-1], handler)
	} else {
		s.router.HandleFunc(path+"/", handler)
	}
}

func (s *server) setRoutes() {
	s.addHandler("/ping", s.handlePing())
}

func (s *server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.router.ServeHTTP(w, req)
}

// NewServer ...
func NewServer() (http.Handler, error) {
	s := &server{
		startedAt: time.Now(),
		router:    http.NewServeMux(),
	}
	s.setRoutes()

	return s, nil
}
