package main

import (
	"log"
	"net/http"
	"testProject/internal/config"
	"testProject/internal/handlers"
	"testProject/internal/storage"
	"testProject/pkg/repository"
)

func main() {
	dbCong := config.GetDbConfig()

	db, err := repository.NewPostgresDB(dbCong)
	if err != nil {
		log.Fatalln(err)
	}

	storageDB := storage.NewStorage(db)

	if err := storage.CreateTable(storageDB); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/create-announcement", handlers.CreateHandler(storageDB))
	http.HandleFunc("/get-announcements/", handlers.GetAllAnnouncement(storageDB))
	http.HandleFunc("/get-announcement/", handlers.GetAnnouncementById(storageDB))

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}

}
