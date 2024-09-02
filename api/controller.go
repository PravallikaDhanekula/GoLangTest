package api

import (
	"database/sql"
	"net/http"
)

func CreateHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed ", http.StatusMethodNotAllowed)
		}

		if err := CreateBookLogic(db, w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func UpdateHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut && r.Method != http.MethodPatch {
			http.Error(w, "method not allowed ", http.StatusMethodNotAllowed)
		}

		err1 := UpdateBookLogic(db, w, r)
		if err1 != nil {
			http.Error(w, err1.Error(), http.StatusInternalServerError)
		}
	}

}

func CreateCustomerHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed ", http.StatusMethodNotAllowed)
		}
		if err := CreateCustomerLogic(db, w, r); err != nil {
			//fmt.Print(ex)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
