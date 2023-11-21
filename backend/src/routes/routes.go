package routes

import (
	"database/sql"
	"net/http"

	"github.com/brivu/reporecon/backend/routes"
)

func SetupRoutes(db *sql.DB) {
	http.HandleFunc("/url", func(w http.ResponseWriter, r *http.Request) {
		handlers.URLHandler(w, r, db)
	})
}
