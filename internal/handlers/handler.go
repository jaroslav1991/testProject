package handlers

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"testProject/internal/validators"
	"time"
)

type Announcement struct {
	Id          int64     `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	Price       int64     `json:"price" db:"price"`
	IdPhoto     string    `json:"id_photo" db:"id_photo"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

func (a *Announcement) TimeNow() time.Time {
	return time.Now()
}

type Storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{db: db}
}

func CreateHandler(storage *Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var announcement Announcement

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

			res, err := storage.CreateAnnouncement(&announcement)
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

func GetAllAnnouncement(storage *Storage) http.HandlerFunc {
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

			res, err := storage.GetAllAnnouncements(count, page)
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

func GetAnnouncementById(storage *Storage) http.HandlerFunc {
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

			res, err := storage.GetAnnouncementsById(id)
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
