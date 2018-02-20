package routers

import (
	"database/sql"
	_ "github.com/lib/pq"
	"go-sports-teams/controllers"
	"net/http"
)

type DataBase struct {
	db *sql.DB
}

func ExposeDB(db *sql.DB) *DataBase {
	return &DataBase{db: db}
}

func (database *DataBase) Login(w http.ResponseWriter, r *http.Request) {
	controllers.ExtendDBUsers(database.db).LoginUser(w, r)
}

func (database *DataBase) Signup(w http.ResponseWriter, r *http.Request) {
	controllers.ExtendDBUsers(database.db).SignupUser(w, r)
}

func (database *DataBase) ToggleTeam(w http.ResponseWriter, r *http.Request) {
	controllers.ExtendDBUsers(database.db).ToggleFavTeam(w, r)
}

func (database *DataBase) GetTeams(w http.ResponseWriter, r *http.Request) {
	controllers.ExtendDBTeams(database.db).GetTeams(w, r)
}
