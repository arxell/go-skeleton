package app

import (
	"fmt"
	"net/http"
)

func (s *server) handlePing() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "pong")
	}
}
