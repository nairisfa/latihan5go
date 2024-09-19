package main

import (
	"bufio"
	"fmt"
	"lala/asal/internal/todo"
	"os"

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

	scanner := bufio.NewScanner(os.Stdin)

		fmt.Print("Todo: ")
		scanner.Scan()
		name := scanner.Text()
		
		fmt.Print("Done?: ")
		scanner.Scan()
		check := scanner.Text()
		var checked bool


		if check == "y"{
			checked = true
		} else {
			checked = false

		}

		todo.CreateTodo(db, name, checked)

		
			
		fmt.Print(todo.FindAllTodos(db))
		fmt.Println("delete: ")
		scanner.Scan()
		name1:= scanner.Text()

		todo.DeleteTodos(db, name1)
		fmt.Print(todo.FindAllTodos(db))
}
