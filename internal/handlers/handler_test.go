package handlers

import (
	"bytes"
	"net/http/httptest"
	"testProject/internal/config"
	"testProject/pkg/repository"
	"testing"
)

func TestCreateHandler(t *testing.T) {
	dbCong, err := config.GetDbConfig()
	if err != nil {
		t.Error(err)
	}

	db, err := repository.NewPostgresDB(dbCong)
	if err != nil {
		t.Error(err)
	}

	req := httptest.NewRequest("POST", "/create-announcement", bytes.NewBuffer([]byte(
		`{"name": "test1", "description": "test description", "price": 100, "id_photo": "qwerty123"}`)))
	res := httptest.NewRecorder()

	CreateHandler(&Storage{db: db})(res, req)
}

func TestGetAllAnnouncement(t *testing.T) {
	dbCong, err := config.GetDbConfig()
	if err != nil {
		t.Error(err)
	}

	db, err := repository.NewPostgresDB(dbCong)
	if err != nil {
		t.Error(err)
	}

	req := httptest.NewRequest("GET", "/get-announcements/?count=10&page=1", nil)
	res := httptest.NewRecorder()

	GetAllAnnouncement(&Storage{db: db})(res, req)

}

func TestGetAnnouncementById(t *testing.T) {
	dbCong, err := config.GetDbConfig()
	if err != nil {
		t.Error(err)
	}

	db, err := repository.NewPostgresDB(dbCong)
	if err != nil {
		t.Error(err)
	}

	req := httptest.NewRequest("GET", "/get-announcement/?id=12", nil)
	res := httptest.NewRecorder()

	GetAnnouncementById(&Storage{db: db})(res, req)

}
