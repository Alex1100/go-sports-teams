package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
)

type DataBase struct {
	db *sql.DB
}

type JSONableSlice []uint8

type UserResponse struct {
	Id            int64
	Username      string
	FavoriteTeams JSONableSlice
}

func ExtendDB(db *sql.DB) *DataBase {
	return &DataBase{db: db}
}

func (u JSONableSlice) MarshalJSON() ([]byte, error) {
	var result string
	if u == nil {
		result = "null"
	} else {
		result = strings.Join(strings.Fields(fmt.Sprintf("%d", u)), ",")
	}
	return []byte(result), nil
}

func (database *DataBase) LoginUser(username, password string, w http.ResponseWriter) error {
	var (
		id             int64
		usrhandle      string
		password_in_db string
		favorite_teams []uint8
	)

	err := database.db.QueryRow("SELECT id, username, password, favorite_teams FROM users WHERE username = $1", username).Scan(&id, &usrhandle, &password_in_db, &favorite_teams)
	if err != nil {
		return err
	} else {
		if err = bcrypt.CompareHashAndPassword([]byte(password_in_db), []byte(password)); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return err
		} else {
			user := &UserResponse{Id: id, Username: username, FavoriteTeams: favorite_teams}

			m, err := json.Marshal(user)
			if err != nil {
				return err
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(string(m))
			return nil
		}
	}
}
