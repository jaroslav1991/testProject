package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testProject/internal/config"
	"testProject/internal/validators"
	"testProject/pkg/repository"
	"testing"
)

func TestCreateHandler(t *testing.T) {
	dbCong := config.GetDbConfig()

	db, err := repository.NewPostgresDB(dbCong)
	if err != nil {
		t.Error(err)
	}

	strg := NewStorage(db)

	req := httptest.NewRequest("POST", "/create-announcement", bytes.NewBuffer([]byte(
		`{"name": "test name3", "description": "test description", "price": 1000, "id_photo": "qwerty123"}`)))
	res := httptest.NewRecorder()

	CreateHandler(strg)(res, req)
}

func TestGetAllAnnouncement(t *testing.T) {
	dbCong := config.GetDbConfig()

	db, err := repository.NewPostgresDB(dbCong)
	if err != nil {
		t.Error(err)
	}

	req := httptest.NewRequest("GET", "/get-announcements/?count=10&page=0", nil)
	res := httptest.NewRecorder()

	GetAllAnnouncement(&Storage{db: db})(res, req)

}

func TestGetAnnouncementById(t *testing.T) {
	dbCong := config.GetDbConfig()

	db, err := repository.NewPostgresDB(dbCong)
	if err != nil {
		t.Error(err)
	}

	req := httptest.NewRequest("GET", "/get-announcement/?id=12", nil)
	res := httptest.NewRecorder()

	GetAnnouncementById(&Storage{db: db})(res, req)

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

//func TestGetAllAnnouncement1(t *testing.T) {
//	type args struct {
//		storage *Storage
//	}
//	tests := []struct {
//		name string
//		args args
//		want http.HandlerFunc
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := GetAllAnnouncement(tt.args.storage); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("GetAllAnnouncement() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
