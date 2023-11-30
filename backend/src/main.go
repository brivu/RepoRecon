package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/brivu/reporecon/backend/src/routes"
	"github.com/rs/cors"

	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("postgres", "user=brianvu dbname=reporecon sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8080"}, // replace with your frontend URL
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	handler := c.Handler(routes.SetupRoutes(db))

	fmt.Println("Server is running on :3000")
	http.ListenAndServe(":3000", handler)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
