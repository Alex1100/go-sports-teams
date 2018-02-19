package main

import (
	"database/sql"
	"fmt"
	gmux "github.com/gorilla/mux"
	"go-sports-teams/routers"
	"log"
	"net/http"
)

const (
	host     = "baasu.db.elephantsql.com"
	port     = 5432
	user     = "kznthyeb"
	password = "raY8OEFmBfOjmMCOK-i2CzdsTJQ6TOqj"
	dbname   = "kznthyeb"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	fmt.Println(psqlInfo)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	mux := gmux.NewRouter()
	routes := routers.ExposeDB(db)

	mux.HandleFunc("/api/login", routes.Login).Methods("POST")
	// mux.Handle("/api/signup", routes.Signup).methods("POST")
	// mux.Handle("/api/sports", routes.GetSports).methods("GET")
	// mux.Handle("/api/teams", routes.GetTeams).methods("GET")
	// mux.Handle("/api/toggle-favorite-team", routes.ToggleTeams).methods("PUT")
	mux.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	http.Handle("/", mux)

	fmt.Println(http.ListenAndServe(":3010", nil))
}
