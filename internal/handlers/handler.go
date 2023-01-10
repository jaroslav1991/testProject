package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"testProject/internal/storage"
	"testProject/internal/validators"
)

func CreateHandler(storageDB *storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var announcement storage.Announcement

		if r.Method == http.MethodPost {

			body, err := io.ReadAll(r.Body)
			if err != nil {
				log.Println("can't read body", err)
				return
			}

			defer r.Body.Close()

			if err := json.Unmarshal(body, &announcement); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("price must be number, not string"))
				log.Println("can't unmarshal body", err)
				return
			}

			res, err := storageDB.CreateAnnouncement(&announcement)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				log.Println("can't create announcement", err)
				return
			}
			bodyRes, err := json.Marshal(res)
			if err != nil {
				return
			}
			w.Write(bodyRes)
			//log.Println(announcement)
		}

	}
}

func GetAllAnnouncement(storageDB *storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {

			err := r.ParseForm()
			if err != nil {
				log.Println("can't parse form", err)
			}

			count, err := validators.ValidateCount(r.Form.Get("count"), w)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			page, err := validators.ValidatePage(r.Form.Get("page"), w)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("page must be positive or zero"))
				return
			}

			res, err := storageDB.GetAllAnnouncements(count, page)
			if err != nil {
				log.Println(err)
				return
			}

			bodyRes, err := json.Marshal(res)
			if err != nil {
				log.Println("can't marshal bodyRes", err)
			}
			log.Println(string(bodyRes))

			w.Write(bodyRes)
		}
	}
}

func GetAnnouncementById(storageDB *storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodGet {

			r.ParseForm()
			s := r.Form.Get("id")
			log.Println(s)

			id, err := strconv.Atoi(s)
			if err != nil {
				log.Println(err)
				return
			}

			res, err := storageDB.GetAnnouncementsById(id)
			if err != nil {
				log.Println(err)
				return
			}

			bodyRes, err := json.Marshal(res)
			if err != nil {
				log.Println("can't marshal bodyRes", err)
				return
			}

			w.Write(bodyRes)
		}

	}
}
