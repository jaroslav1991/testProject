package validators

import (
	"errors"
	"log"
	"net/http"
	"strconv"
)

const (
	pageDefault  = 0
	countDefault = 10
)

func ValidatePage(s string, w http.ResponseWriter) (int, error) {
	page, err := strconv.Atoi(s)
	if err != nil {
		log.Println("can't convert, page must be integer")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("page must be integer, not string"))
		return 0, err
	}

	if page < 0 {
		log.Println("page must be positive or zero")
		w.Write([]byte("page must be positive or zero"))
		w.WriteHeader(http.StatusBadRequest)
		return 0, err
	}
	w.WriteHeader(http.StatusOK)
	return page, nil
}

func ValidateCount(s string, w http.ResponseWriter) (int, error) {
	count, err := strconv.Atoi(s)
	if err != nil {
		log.Println("can't convert, count must be integer")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("count must be integer, not string"))
		return 0, err
	}
	if count != 10 {
		log.Println("count must be 10")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("count must be 10"))
		return 0, err
	}
	return count, nil
}

func ValidatePrice(price int64) (int64, error) {

	if price < 0 {

		return 0, errors.New("price must be positive")
	}
	return price, nil
}
