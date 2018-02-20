package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
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

func ExtendDBUsers(db *sql.DB) *DataBase {
	return &DataBase{db: db}
}

func (u JSONableSlice) MarshalJSON() ([]byte, error) {
	var result string
	if u == nil {
		result = "null"
	} else {
		result = strings.Join(strings.Fields(fmt.Sprintf("\n%d", u)), ",")
	}
	return []byte(result), nil
}

func (database *DataBase) LoginUser(w http.ResponseWriter, r *http.Request) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	req_params := make(map[string]string)
	e := json.Unmarshal(body, &req_params)

	if e != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return e
	}
	req_keys := make([]string, len(req_params))
	i := 0
	for _, v := range req_params {
		req_keys[i] = v
		i++
	}

	username := req_keys[0]
	password := req_keys[1]

	var (
		id             int64
		usrhandle      string
		password_in_db string
		favorite_teams []uint8
	)

	err = database.db.QueryRow("SELECT id, username, password, favorite_teams FROM users WHERE username = $1", username).Scan(&id, &usrhandle, &password_in_db, &favorite_teams)
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

func (database *DataBase) SignupUser(w http.ResponseWriter, r *http.Request) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	req_params := make(map[string]string)
	e := json.Unmarshal(body, &req_params)
	if e != nil {
		panic(e)
	}
	req_keys := make([]string, len(req_params))
	i := 0
	for _, v := range req_params {
		req_keys[i] = v
		i++
	}

	username := req_keys[0]
	password := req_keys[1]

	secret, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	_, err = database.db.Exec(`INSERT INTO users (username, password, favorite_teams) values($1, $2, $3)`, username, secret, `{}`)
	if err != nil {
		fmt.Println("ERR IS: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	var (
		id             int64
		usrhandle      string
		password_in_db string
		favorite_teams []uint8
	)

	err = database.db.QueryRow("SELECT id, username, password, favorite_teams FROM users WHERE username = $1", username).Scan(&id, &usrhandle, &password_in_db, &favorite_teams)
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

func (database *DataBase) ToggleFavTeam(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("WOOOOOOO")
	return nil
}
