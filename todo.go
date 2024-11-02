package main

import "time"

type Todo struct {
	Title     string
	Completed bool
	CreatedAt time.Time
	// Pointer as this can be null
	CompletedAt *time.Time
}

type ToDos []Todo

func (todos *Todos) add(title string) {
	// Adding a todo item to our list of todos

	todo_item := Todo{
		Title:       title,
		Completed:   false,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
	}

	*todos = append(*todos, todo_item)

}
