package routers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"go-sports-teams/controllers"
	"io/ioutil"
	"net/http"
)

type DataBase struct {
	db *sql.DB
}

type UserResponse struct {
	Username string
}

func ExposeDB(db *sql.DB) *DataBase {
	return &DataBase{db: db}
}

func (database *DataBase) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("FUNC SHOT OFF IN THE LOGIN FUNC!!")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	fmt.Println("REQ IS: ", body)

	ctrl := controllers.ExtendDB(database.db)

	res, err := ctrl.LoginUser("alex1100.software@gmail.com", "alex1100")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(&UserResponse{Username: "Alex"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Println("SHOULD RETURN: ", data)
}
