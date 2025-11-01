package main

import (
	"to-do/internal/todo"
)

func main() {
	todos := todo.Todos{}
	todos.Add("kick someone")
	todos.Add("like someone")
	todos.Toggle(0)
	todos.Print()
}
