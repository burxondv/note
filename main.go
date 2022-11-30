package main

import (
	"fmt"
	"log"

	"github.com/burxondv/note/config"
	"github.com/burxondv/note/storage"
	"github.com/burxondv/note/storage/repo"
	"github.com/jmoiron/sqlx"
)

func main() {

	cfg := config.Load(".")

	psqlUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Database,
	)

	psqlConn, err := sqlx.Connect("postgres", psqlUrl)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	storage := storage.NewStoragePg(psqlConn)

	// Botta faqat 'Note' ni 'Update'i qilingan, lekin hammasi ishlavotti

	note, err := storage.Note().Update(&repo.Note{
		UserID: 3,
		Title: "I'm singer",
		Description: "My first song is \"Fetish\"",
		ID: 3,
	})
	if err!= nil {
		log.Fatalf("failed to update note: %v", err)
	}
	fmt.Println(note)

}
