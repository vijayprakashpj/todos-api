package domain

import "github.com/google/uuid"

type Todo struct {
	ID    string
	Title string
	Done  bool
	Body  string
}

func (todo *Todo) Create(title string, body string) *Todo {
	newId, _ := uuid.NewUUID()
	todo = &Todo{
		ID:    newId.String(),
		Title: title,
		Body:  body,
		Done:  false,
	}

	return todo
}

func (todo *Todo) Update(updatedTodo Todo) *Todo {
	todo.Title = updatedTodo.Title
	todo.Body = updatedTodo.Body
	todo.Done = updatedTodo.Done

	return todo
}

func (todo Todo) GetTitle() string {
	return todo.Title
}

func (todo Todo) Get() Todo {
	return todo
}
