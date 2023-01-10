package storage

import (
	"log"
)

const (
	getByIdQuery = `select name, price, id_photo from announcements where id=$1`
)

type AnnouncementById struct {
	Name    string                    `json:"name" db:"name"`
	Price   int64                     `json:"price" db:"price"`
	IdPhoto string                    `json:"id_photo" db:"id_photo"`
	Fields  *OptionalAnnouncementById `json:"fields"`
}

type OptionalAnnouncementById struct {
	Description string `json:"description" db:"description"`
}

func (s *Storage) GetAnnouncementsById(id int) (*AnnouncementById, error) {
	var announcement AnnouncementById
	rows, err := s.Db.Query(getByIdQuery, id)
	if err != nil {
		return nil, err
	}

	log.Printf("open announcement with id = %d", id)
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&announcement.Name, &announcement.Price, &announcement.IdPhoto); err != nil {
			return nil, err
		}
	}
	res := &AnnouncementById{
		Name:    announcement.Name,
		Price:   announcement.Price,
		IdPhoto: announcement.IdPhoto,
	}
	return res, nil
}
