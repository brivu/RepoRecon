package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/brivu/reporecon/backend/src/models"

)

func URLHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var u models.Url
	if r.Method == http.MethodPost {
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = db.Exec("INSERT INTO repositories (repo_url) VALUES ($1)", u.Url)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprint(w, "URL successfully stored")
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func ExtractOwnerAndRepo (githubUrl string){
	
}