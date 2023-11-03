package router

import (
	"github.com/Scramjet911/learning-go/go-books/auth"
	"github.com/Scramjet911/learning-go/go-books/book"
	"github.com/Scramjet911/learning-go/go-books/server"
	"github.com/Scramjet911/learning-go/go-books/stage"
	"github.com/gorilla/mux"
)

func CreateRouter(s *server.Server) *mux.Router {
	r := mux.NewRouter()

	book.RegisterBookRoutes(s, r)

	auth.RegisterAuthRoutes(s, r)

	stage.RegisterStageRoutes(s, r)

	return r
}
