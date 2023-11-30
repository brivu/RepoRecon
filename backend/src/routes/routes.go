package routes

import (
	"database/sql"
	"net/http"

	"github.com/brivu/reporecon/backend/src/handlers"
)

func SetupRoutes(db *sql.DB) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/url", func(w http.ResponseWriter, r *http.Request) {
		handlers.URLHandler(w, r, db)
	})

	return mux
}
