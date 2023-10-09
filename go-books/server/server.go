package server

import (
	"net/http"

	"github.com/Scramjet911/learning-go/go-books/config"
	"github.com/Scramjet911/learning-go/go-books/db"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Config *config.Config
}

func NewServer(cfg *config.Config) *Server {
	db.Connect(cfg)
	db.MigrateTables()
	return &Server{
		DB:     db.GetDB(),
		Config: cfg,
	}
}

func (server *Server) Start(addr string, r *mux.Router) error {
	return http.ListenAndServe(":"+addr, r)
}
