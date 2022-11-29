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
		log.Fatalf("failed to connect database; %v", err)
	}

	storage := storage.NewStoragePg(psqlConn)

	note, err := storage.Note().Create(&repo.Note{
		UserID:      1,
		Title:       "hello",
		Description: "how are you",
	})
	if err != nil {
		log.Fatalf("failed to create note: %v", err)
	}

	user, err := storage.User().Create(&repo.User{
		FirstName: "tobias",
		LastName:  "amigo",
		Email:     "amigoguys@gmail.com",
	})

	fmt.Println(user)
	fmt.Println(note)
}
