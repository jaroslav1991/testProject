package storage

const (
	createTableQuery = `create table if not exists announcements (
    id bigserial primary key,
    name varchar(200) not null,
    description varchar(1000) not null,
    price bigint not null,
    id_photo varchar(50),
    created_at timestamp 
);`
)

func CreateTable(s *Storage) error {
	_, err := s.Db.Exec(createTableQuery)
	return err
}
