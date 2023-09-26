package main

import (
	"fmt"
	"log"

	"github.com/Scramjet911/learning-go/go-movies/config"
	movie "github.com/Scramjet911/learning-go/go-movies/movie"
	"github.com/Scramjet911/learning-go/go-movies/server"
)

func main() {
	cfg := config.NewConfig()
	s := server.NewServer(cfg)

	r := movie.ConfigureRoutes(s)

	fmt.Printf("Starting server at port 9090\n")

	log.Fatal(s.Start("9090", r))
}
