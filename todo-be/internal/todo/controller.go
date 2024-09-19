package todo

import (
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

type Controller struct {
	db *gorm.DB
}

func (c *Controller) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	todos := FindAllTodos(c.db)

	data, err := json.Marshal(todos)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func InitController(db *gorm.DB) *Controller {
	return &Controller{
		db: db,
	}
}

