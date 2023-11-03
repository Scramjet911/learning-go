package book

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	bm "github.com/Scramjet911/learning-go/go-books/book/models"
	s "github.com/Scramjet911/learning-go/go-books/server"
	"github.com/Scramjet911/learning-go/go-books/util"
	"github.com/gorilla/mux"
	"gorm.io/gorm/clause"
)

type BookHandler struct {
	server *s.Server
}

func NewBookHandler(server *s.Server) *BookHandler {
	return &BookHandler{server: server}
}

func (b *BookHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("GetBooksssss")
	w.Header().Set("Content-Type", "application/json")
	var books []bm.Book
	if res := b.server.DB.Find(&books); res.Error == nil {
		fmt.Printf("books: %v", books)
		res, _ := json.Marshal(books)
		w.WriteHeader(http.StatusOK)
		fmt.Printf("result %v, books: %v", res, books)
		w.Write(res)
	} else {
		fmt.Printf("Get books error: %v", res.Error)
		log.Fatal(res.Error)
	}
}

func (b *BookHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	w.Header().Set("Content-Type", "application/json")

	if params["id"] != "" {
		book := bm.Book{ID: params["id"]}
		if res := b.server.DB.First(&book); res.Error == nil {
			res, _ := json.Marshal(book)
			w.WriteHeader(http.StatusOK)
			w.Write(res)
		} else {
			http.Error(w, "Entity not found", 404)
		}
	} else {
		http.Error(w, "Bad request", 400)
	}
}

func (b *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	w.Header().Set("Content-Type", "application/json")

	if params["id"] != "" {
		book := bm.Book{ID: params["id"]}
		if res := b.server.DB.Clauses(clause.Returning{}).Delete(&book); res.Error == nil {
			res, _ := json.Marshal(book)
			w.WriteHeader(http.StatusOK)
			w.Write(res)
			fmt.Printf("book : %v\n marshall: %v", book, res)
		} else {
			http.Error(w, "Entity not found", 404)
		}
	} else {
		http.Error(w, "Bad request", 400)
	}
}

func (b *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	createBook := &bm.Book{}
	w.Header().Set("Content-Type", "application/json")
	util.ParseBody(r, createBook)

	createBook.ID = strconv.Itoa(rand.Intn(10000000))

	fmt.Printf("%v", createBook)

	b.server.DB.Create(&createBook)

	res, _ := json.Marshal(createBook)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (b *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	updateBook := &bm.Book{}
	util.ParseBody(r, updateBook)

	bookId := params["id"]
	if bookId != "" {
		book := bm.Book{ID: bookId}
		if res := b.server.DB.First(&book); res.Error == nil {
			fmt.Printf("New Book: %v, Old Book: %v", updateBook, book)
			if updateBook.Author == "" {
				updateBook.Author = book.Author
			}
			if updateBook.Title == "" {
				updateBook.Title = book.Title
			}
			if updateBook.Isbn == "" {
				updateBook.Isbn = book.Isbn
			}
			if updateBook.ID == "" {
				updateBook.ID = book.ID
			}
			fmt.Printf("Final Book: %v", updateBook)

			b.server.DB.Save(&updateBook)

			res, _ := json.Marshal(updateBook)
			w.WriteHeader(http.StatusOK)
			w.Write(res)
		} else {
			http.Error(w, "Entity not found", 404)
		}
	} else {
		http.Error(w, "Bad request", 400)
	}

}
