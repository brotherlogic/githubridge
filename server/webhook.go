package server

import (
	"io"
	"log"
	"net/http"
)

func (s *Server) githubwebhook(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Unable to parse body: %v", err)
		return
	}

	log.Printf("Received webhook: %v", string(body))
}
