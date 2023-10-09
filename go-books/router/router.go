package router

import (
	"github.com/Scramjet911/learning-go/go-books/book"
	"github.com/Scramjet911/learning-go/go-books/server"
	"github.com/gorilla/mux"
)

func CreateRouter(s *server.Server) *mux.Router {
	r := mux.NewRouter()

	book.RegisterBookRoutes(s, r)

	return r
}
