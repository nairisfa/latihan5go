package todo

import (
	"net/http"

	"gorm.io/gorm"
)

func GetRouteHandler(db *gorm.DB) *http.ServeMux {
	mux := http.NewServeMux()

	c := InitController(db)

	mux.HandleFunc("GET /", c.GetAllTodos)
	// mux.HandleFunc("POST /", c.CreateTodo)

	return mux
}

