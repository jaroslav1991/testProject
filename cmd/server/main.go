package main

import (
	"log"
	"net/http"
	"testProject/internal/config"
	"testProject/internal/handlers"
	"testProject/pkg/repository"
)

func main() {

	dbCong, err := config.GetDbConfig()
	if err != nil {
		log.Fatalln(err)
	}

	db, err := repository.NewPostgresDB(dbCong)
	if err != nil {
		log.Fatalln(err)
	}

	storage := handlers.NewStorage(db)

	http.HandleFunc("/create-announcement", handlers.CreateHandler(storage))
	http.HandleFunc("/get-announcements/", handlers.GetAllAnnouncement(storage))
	http.HandleFunc("/get-announcement/", handlers.GetAnnouncementById(storage))

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatalln(err)
	}

}
