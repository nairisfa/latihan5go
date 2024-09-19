package todo

import (
	"gorm.io/gorm"
)

type Todo struct {
	ID      uint
	Name    string
	Checked bool
}

// Migrate function to create the Todo table
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Todo{})
}

// Function to create a new todo
func CreateTodo(db *gorm.DB, name string, checked bool) Todo {
	todo := Todo{
		Name:    name,
		Checked: checked,
	}
	db.Create(&todo)
	return todo
}

// Function to get all todos
func FindAllTodos(db *gorm.DB) []Todo {
	var todos []Todo
	db.Find(&todos)
	return todos
}

// Function to find a todo by name
func FindTodo(db *gorm.DB, name string) (Todo, bool) {
	var todo Todo
	result := db.Where("name = ?", name).First(&todo)

	// Check if a record was found
	if result.RowsAffected > 0 {
		return todo, true
	}

	return todo, false
}

// Function to delete a todo by name
func DeleteTodos(db *gorm.DB, name string) error {
	// First, find the todo to make sure it exists
	todo, exists := FindTodo(db, name)
	if !exists {
		return gorm.ErrRecordNotFound // Return an error if the record is not found
	}

	// Delete the todo from the database
	db.Delete(&todo)

	return nil // Return nil if successful
}
