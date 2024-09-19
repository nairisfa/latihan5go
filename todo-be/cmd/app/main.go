package main

import (
	"lala/asal/internal/todo"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Initialize the database connection
func main() {
	// Buka database menggunakan driver baru
	db, err := gorm.Open(postgres.Open("postgres://postgres:123456@localhost:5432/postgres"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	todo.Migrate(db)

	mux := todo.GetRouteHandler(db)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	srv.ListenAndServe()
}
