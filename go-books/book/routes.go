package book

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

func RegisterBookRoutes(s *server.Server, r *mux.Router) {
	bookHandler := NewBookHandler(s)
	r.HandleFunc("/", health).Methods("GET")
	r.HandleFunc("/books", bookHandler.GetBooks).Methods("GET")
	r.HandleFunc("/books/{id}", bookHandler.GetBook).Methods("GET")
	r.HandleFunc("/books", bookHandler.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", bookHandler.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", bookHandler.DeleteBook).Methods("DELETE")
}
