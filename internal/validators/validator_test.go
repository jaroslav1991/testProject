package validators

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

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
			got, err := ValidatePage(tt.args.s, tt.args.w)
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
			got, err := ValidateCount(tt.args.s, tt.args.w)
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
			got, err := ValidatePrice(tt.args.price)
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
