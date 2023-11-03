package stage

import (
	"github.com/Scramjet911/learning-go/go-books/server"
	"github.com/gorilla/mux"
)

func RegisterStageRoutes(s *server.Server, r *mux.Router) {
	StageHandler := NewStageHandler(s)
	r.HandleFunc("/alpha", StageHandler.FirstStage).Methods("POST")
}
