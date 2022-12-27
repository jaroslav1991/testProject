package handlers

import "testProject/internal/validators"

const (
	createQuery = `insert into announcements (name, description, price, id_photo, created_at) values ($1, $2, $3, $4, $5) returning id`
)

func (s *Storage) CreateAnnouncement(announcement *Announcement) (*Announcement, error) {

	price, err := validators.ValidatePrice(announcement.Price)
	if err != nil {
		return nil, err
	}
	price = announcement.Price

	rows, err := s.db.Query(createQuery, announcement.Name, announcement.Description, price, announcement.IdPhoto, announcement.TimeNow())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&announcement.Id); err != nil {
			return nil, err
		}
	}
	res := &Announcement{
		Id:          announcement.Id,
		Name:        announcement.Name,
		Description: announcement.Description,
		Price:       price,
		IdPhoto:     announcement.IdPhoto,
		CreatedAt:   announcement.TimeNow(),
	}
	return res, nil
}
