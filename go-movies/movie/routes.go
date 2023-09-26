package movie

import (
	"fmt"

	s "github.com/Scramjet911/learning-go/go-movies/server"
	"github.com/gorilla/mux"
)

func ConfigureRoutes(server *s.Server) *mux.Router {
	r := mux.NewRouter()
	movieHandler := NewMovieHandler(server)

	r.HandleFunc("/movies", movieHandler.GetMovies).Methods("GET")
	r.HandleFunc("/movies/[id]", movieHandler.GetMovie).Methods("GET")
	r.HandleFunc("/movies", movieHandler.CreateMovie).Methods("POST")
	r.HandleFunc("/movies/[id]", movieHandler.UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/[id]", movieHandler.DeleteMovie).Methods("DELETE")

	fmt.Printf("Router setup complete\n")

	return r
}
