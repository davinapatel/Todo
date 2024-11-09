package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type ActionItem struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time // Using a pointer because this field can be null
}

type ToDos []ActionItem

func (todos *ToDos) add(title string) {
	// Adding a todo item to our list of todos

	todo_item := ActionItem{
		Title:       title,
		Completed:   false,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
	}

	*todos = append(*todos, todo_item)
}

func (todos *ToDos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("invalid index")
		fmt.Println(err)
		return err
	}
	return nil
}

func (todos *ToDos) delete(index int) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	//fmt.Println(*todos, &todos)
	*todos = append(t[:index], t[index+1:]...) // removes element at index from slice
	// ... is the unpack operator, so it adds the elements of the second slice as sep args

	return nil
}

func (todos *ToDos) toggle(index int) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	isCompleted := t[index].Completed

	if !isCompleted {
		completionTime := time.Now()
		t[index].CompletedAt = &completionTime
	}

	t[index].Completed = !isCompleted

	return nil
}

func (todos *ToDos) edit(index int, title string) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	t[index].Title = title
	return nil
}

func (todos *ToDos) print() {
	// Prints ToDo table out

	table := table.New(os.Stdout)
	table.SetRowLines(true)
	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")

	for index, t := range *todos {
		completed := "\u274c"
		completedAt := ""

		if t.Completed {
			completed = "\u2611\ufe0f"
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(time.RFC1123)
			}
		}

		table.AddRow(
			strconv.Itoa(index),
			t.Title,
			completed,
			t.CreatedAt.Format(time.RFC1123),
			completedAt,
		)
	}

	table.Render()
}
