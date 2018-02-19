package controllers

import (
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	id       int
	Username string
	Password string
}

type DataBase struct {
	db *sql.DB
}

type UserResponse struct {
	Username string
}

func ExtendDB(db *sql.DB) *DataBase {
	return &DataBase{db: db}
}

func (database *DataBase) LoginUser(username, password string) (UserResponse, error) {
	var (
		id             string
		password_in_db string
	)

	err := database.db.QueryRow("SELECT * FROM users WHERE username = $1").Scan(&id, &username, &password_in_db)

	if err != nil {
		fmt.Printf("ERROR LOGGING IN: %s", err)
		return UserResponse{Username: ""}, err
	} else {
		if err = bcrypt.CompareHashAndPassword([]byte(password_in_db), []byte(password)); err != nil {
			return UserResponse{Username: ""}, err
		} else {
			user := UserResponse{Username: username}

			// js, err := json.Marshal(user)
			// if err != nil {
			// 	return UserResponse{Username: ""}, err
			// }

			return user, nil
		}
	}

}
