package storage

import "time"

type Announcement struct {
	Id          int64     `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	Price       int64     `json:"price" db:"price"`
	IdPhoto     string    `json:"id_photo" db:"id_photo"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}
