package storage

const (
	getAllQuery = `select name, price, id_photo from announcements order by price asc, created_at asc limit $1 offset $2`
)

type AnnouncementsResponse struct {
	Name    string `json:"name" db:"name"`
	Price   int64  `json:"price" db:"price"`
	IdPhoto string `json:"id_photo" db:"id_photo"`
}

func (s *Storage) GetAllAnnouncements(count, page int) ([]*AnnouncementsResponse, error) {
	var announcements []*AnnouncementsResponse

	offset := page * count

	rows, err := s.Db.Query(getAllQuery, count, offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var announcement AnnouncementsResponse
		if err := rows.Scan(&announcement.Name, &announcement.Price, &announcement.IdPhoto); err != nil {
			return nil, err
		}
		announcements = append(announcements, &announcement)
	}

	return announcements, nil
}
