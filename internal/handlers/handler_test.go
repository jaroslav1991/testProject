package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testProject/internal/config"
	"testProject/internal/storage"
	"testProject/internal/validators"
	"testProject/pkg/repository"
	"testing"
	"time"
)

func TestCreateHandler(t *testing.T) {
	dbCong := config.GetDbConfig()

	db, err := repository.NewPostgresDB(dbCong)
	if err != nil {
		t.Error(err)
	}

	strg := storage.NewStorage(db)

	req := httptest.NewRequest("POST", "/create-announcement", bytes.NewBuffer([]byte(
		`{"name": "test name3", "description": "test description", "price": 1000, "id_photo": "qwerty123"}`)))
	res := httptest.NewRecorder()

	handler := CreateHandler(strg)
	handler(res, req)

	var newObj storage.Announcement

	if err := json.Unmarshal(res.Body.Bytes(), &newObj); err != nil {
		t.Error(err)
	}

	now := time.Now()
	newObj.Id = 0
	newObj.CreatedAt = now

	assert.Equal(t, storage.Announcement{
		Name:        "test name3",
		Description: "test description",
		Price:       1000,
		IdPhoto:     "qwerty123",
		CreatedAt:   now,
	}, newObj)
}

func TestGetAllAnnouncement(t *testing.T) {
	dbCong := config.GetDbConfig()

	db, err := repository.NewPostgresDB(dbCong)
	if err != nil {
		t.Error(err)
	}

	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		t.Error(err)
	}

	defer tx.Rollback()

	if _, err := tx.Exec("delete from announcements"); err != nil {
		t.Error(err)
	}

	strg := storage.NewStorage(tx)

	now := time.Now()

	_, err = strg.CreateAnnouncement(&storage.Announcement{
		Id:          1,
		Name:        "name1",
		Description: "description1",
		Price:       150,
		IdPhoto:     "1",
		CreatedAt:   now,
	})

	_, err = strg.CreateAnnouncement(&storage.Announcement{
		Id:          2,
		Name:        "name2",
		Description: "description2",
		Price:       120,
		IdPhoto:     "2",
		CreatedAt:   now,
	})

	req := httptest.NewRequest("GET", "/get-announcements/?count=10&page=0", nil)
	res := httptest.NewRecorder()

	handler := GetAllAnnouncement(strg)
	handler(res, req)

	var newObj []storage.AnnouncementsResponse

	if err := json.Unmarshal(res.Body.Bytes(), &newObj); err != nil {
		t.Error(err)
	}

	assert.Equal(t, []storage.AnnouncementsResponse{
		{
			Name:    "name2",
			Price:   120,
			IdPhoto: "2",
		},
		{
			Name:    "name1",
			Price:   150,
			IdPhoto: "1",
		},
	}, newObj)
}

func TestGetAnnouncementById(t *testing.T) {
	dbCong := config.GetDbConfig()

	db, err := repository.NewPostgresDB(dbCong)
	if err != nil {
		t.Error(err)
	}

	strg := storage.NewStorage(db)

	req := httptest.NewRequest("GET", "/get-announcement/?id=160", nil)
	res := httptest.NewRecorder()

	handler := GetAnnouncementById(strg)
	handler(res, req)
}

func TestValidatePage(t *testing.T) {
	type args struct {
		s string
		w http.ResponseWriter
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "notInteger",
			args: args{
				s: "qwerty",
				w: httptest.NewRecorder(),
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "lessZero",
			args: args{
				s: "-10",
				w: httptest.NewRecorder(),
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "successResponse",
			args: args{
				s: "2",
				w: httptest.NewRecorder(),
			},
			want:    2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := validators.ValidatePage(tt.args.s, tt.args.w)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidatePage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ValidatePage() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateCount(t *testing.T) {
	type args struct {
		s string
		w http.ResponseWriter
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "notInteger",
			args: args{
				s: "qwerty",
				w: httptest.NewRecorder(),
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "barRequest",
			args: args{
				s: "9",
				w: httptest.NewRecorder(),
			},
			want:    0,
			wantErr: false,
		},

		{
			name: "successResponse",
			args: args{
				s: "10",
				w: httptest.NewRecorder(),
			},
			want:    10,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := validators.ValidateCount(tt.args.s, tt.args.w)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateCount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ValidateCount() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidatePrice(t *testing.T) {
	type args struct {
		price int64
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name:    "lessInteger",
			args:    args{price: -10},
			want:    0,
			wantErr: true,
		},
		{
			name:    "successResponse",
			args:    args{price: 100},
			want:    100,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := validators.ValidatePrice(tt.args.price)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidatePrice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ValidatePrice() got = %v, want %v", got, tt.want)
			}
		})
	}
}
