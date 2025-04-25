package controllers

import (
	"encoding/json"
	"godock/backend/config"
	"godock/backend/models"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	db := config.InitDB()
	defer db.Close()

	rows, err := db.Query("SELECT id, name, birthdate, gender, job, photo FROM users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.BirthDate, &user.Gender, &user.Job, &user.Photo); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
