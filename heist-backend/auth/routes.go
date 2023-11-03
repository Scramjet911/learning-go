package auth

import (
	"encoding/json"
	"net/http"

	"github.com/Scramjet911/learning-go/go-books/server"
	"github.com/gorilla/mux"
)

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res, _ := json.Marshal("Health Check!")
	w.Write(res)
}

func RegisterAuthRoutes(s *server.Server, r *mux.Router) {
	AuthHandler := NewAuthHandler(s)
	r.HandleFunc("/", health).Methods("GET")
	r.HandleFunc("/name", AuthHandler.GetNameHash).Methods("POST")
}
