package server

import (
	"net/http"

	"github.com/Scramjet911/learning-go/go-movies/config"
	"github.com/gorilla/mux"

	"github.com/jinzhu/gorm"
)

type Server struct {
	DB     *gorm.DB
	Config *config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		Config: cfg,
	}
}

func (server *Server) Start(addr string, r *mux.Router) error {
	return http.ListenAndServe(":"+addr, r)
}
