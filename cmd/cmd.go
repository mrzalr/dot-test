package cmd

import (
	"log"
	"os"

	"github.com/dot-test/internal/models"
	"github.com/dot-test/internal/server"
	"github.com/dot-test/pkg/db"
	"github.com/joho/godotenv"
)

func StartApp() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("load env error: %s\n", err.Error())
	}

	db, err := db.NewMysql()
	if err != nil {
		log.Fatalf("connecting database error: %s\n", err.Error())
	}

	if err := db.AutoMigrate(
		&models.Post{},
	); err != nil {
		log.Fatalf("auto migrate error: %s\n", err.Error())
	}

	s := server.New(db)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	log.Printf("Listening on port :%s\n", port)
	if err := s.Run(port); err != nil {
		log.Fatalf("run server error: %s\n", err.Error())
	}
}
