package db

import (
	"fmt"

	bm "github.com/Scramjet911/learning-go/go-books/book/models"
	"github.com/Scramjet911/learning-go/go-books/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect(cfg *config.Config) {
	Dbdriver := cfg.DB.Driver
	DbHost := cfg.DB.Host
	DbPort := cfg.DB.Port
	DbUser := cfg.DB.User
	DbName := cfg.DB.Name
	DbPassword := cfg.DB.Password

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Cannot connect to %s database\n", Dbdriver)
		panic(err)
	} else {
		fmt.Printf("We are connected to the %s database\n", Dbdriver)
	}
	db = DB
}

func MigrateTables() {
	db.AutoMigrate(&bm.Book{})
}

func GetDB() *gorm.DB {
	return db
}
