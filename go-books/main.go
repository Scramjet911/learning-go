package main

import (
	"fmt"
	"log"

	"github.com/Scramjet911/learning-go/go-books/config"
	"github.com/Scramjet911/learning-go/go-books/router"
	"github.com/Scramjet911/learning-go/go-books/server"
)

func main() {
	cfg := config.NewConfig()
	s := server.NewServer(cfg)

	r := router.CreateRouter(s)

	fmt.Printf("Starting server at port 9090\n")

	log.Fatal(s.Start("9090", r))
}
