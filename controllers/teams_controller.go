package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

func ExtendDBTeams(db *sql.DB) *DataBase {
	return &DataBase{db: db}
}

func (database *DataBase) GetTeams(w http.ResponseWriter, r *http.Request) error {
	var (
		id              int64
		full_name       string
		display_name    string
		abbr            string
		logo            string
		primary_color   string
		secondary_color string
		sports_id       int64
	)

	fin := make([]map[string]interface{}, 0)

	rows, err := database.db.Query("SELECT id, full_name, display_name, abbr, logo, primary_color, secondary_color, sports_id FROM teams")

	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&id, &full_name, &display_name, &abbr, &logo, &primary_color, &secondary_color, &sports_id)
		if err != nil {
			log.Fatal(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return err
		}

		test_result := make(map[string]interface{})

		test_result["id"] = id
		test_result["full_name"] = full_name
		test_result["display_name"] = display_name
		test_result["abbr"] = abbr
		test_result["logo"] = logo
		test_result["primary_color"] = primary_color
		test_result["secondary_color"] = secondary_color
		test_result["sports_id"] = sports_id

		fin = append(fin, test_result)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	m, err := json.Marshal(fin)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(string(m))
	return nil
}
