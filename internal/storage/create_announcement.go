package storage

import (
	"testProject/internal/validators"
	"time"
)

const (
	createQuery = `insert into announcements (name, description, price, id_photo, created_at)
						values ($1, $2, $3, $4, $5) 
						returning id, name, description, price, id_photo, created_at `
)

func (s *Storage) CreateAnnouncement(announcement *Announcement) (*Announcement, error) {

	_, err := validators.ValidatePrice(announcement.Price)
	if err != nil {
		return nil, err
	}

	announcement.CreatedAt = time.Now()

	rows, err := s.Db.Query(createQuery, announcement.Name, announcement.Description, announcement.Price, announcement.IdPhoto, announcement.CreatedAt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&announcement.Id,
			&announcement.Name,
			&announcement.Description,
			&announcement.Price,
			&announcement.IdPhoto,
			&announcement.CreatedAt,
		); err != nil {
			return nil, err
		}
	}

	return announcement, nil
}
