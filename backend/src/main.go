package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/brivu/reporecon/backend/src/models"
	"github.com/brivu/reporecon/backend/src/routes"
	"github.com/brivu/reporecon/backend/src/handlers"


	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("postgres", "user=brianvu dbname=reporecon sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	routes.SetupRoutes(db)

	http.ListenAndServe(":8080", nil)
}
