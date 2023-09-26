package movie

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	s "github.com/Scramjet911/learning-go/go-movies/server"
	"github.com/gorilla/mux"
)

type MovieHandler struct {
	server    *s.Server
	movieList []Movie
}

func NewMovieHandler(server *s.Server) *MovieHandler {
	return &MovieHandler{server: server}
}

func (m *MovieHandler) GetMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%v", m.movieList)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(m.movieList)
}

func (m *MovieHandler) GetMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	w.Header().Set("Content-Type", "application/json")

	for _, movie := range m.movieList {
		if movie.ID == params["id"] {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func (m *MovieHandler) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	w.Header().Set("Content-Type", "application/json")

	for index, movie := range m.movieList {
		if movie.ID == params["id"] {
			m.movieList = append(m.movieList[:index], m.movieList[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(m.movieList)
}

func (m *MovieHandler) CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))

	m.movieList = append(m.movieList, movie)

	fmt.Printf("%v", m.movieList)

	json.NewEncoder(w).Encode(m.movieList)
}

func (m *MovieHandler) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var updatedMovie Movie
	_ = json.NewDecoder(r.Body).Decode(&updatedMovie)
	w.Header().Set("Content-Type", "application/json")

	for index, movie := range m.movieList {
		if movie.ID == params["id"] {
			m.movieList = append(m.movieList[:index], m.movieList[index+1:]...)
			m.movieList = append(m.movieList, updatedMovie)
			break
		}
	}
	json.NewEncoder(w).Encode(m.movieList)
}
