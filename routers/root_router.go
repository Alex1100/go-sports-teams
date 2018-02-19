package routers

import (
	"database/sql"
	"encoding/json"
	_ "github.com/lib/pq"
	"go-sports-teams/controllers"
	"io/ioutil"
	"net/http"
)

type DataBase struct {
	db *sql.DB
}

func ExposeDB(db *sql.DB) *DataBase {
	return &DataBase{db: db}
}

func (database *DataBase) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
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

	ctrl := controllers.ExtendDB(database.db)

	ctrl.LoginUser(req_keys[0], req_keys[1], w)
}
