package main

import (
	"to-do/internal/config"
	"to-do/internal/repository"
	"to-do/internal/todo"
)

func main() {
	todos := todo.Todos{}
	storage := repository.NewStorage[todo.Todos]("todos.json")
	storage.Load(&todos)
	cmdFlags := config.NewCmdFlags()
	cmdFlags.Execute(&todos)
	storage.Save(todos)
}
