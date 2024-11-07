package domain

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	// GIVEN
	title := "Test Title"
	body := "Important task"
	todo := new(Todo)

	// WHEN
	newTodo := todo.Create(title, body)

	// THEN
	assert.NotEqual(t, "", newTodo.ID)
	assert.Equal(t, newTodo.Title, title)
	assert.Equal(t, newTodo.Body, body)
}

func TestUpdate(t *testing.T) {
	// GIVEN
	newID, _ := uuid.NewUUID()
	newTodo := &Todo{
		ID:    newID.String(),
		Title: "Some Title",
		Body:  "Some Body",
		Done:  true,
	}
	newTitle := "new title"

	// WHEN
	updatedTodo := newTodo.Update(Todo{Title: newTitle, Done: false})

	// THEN
	assert.Equal(t, updatedTodo.ID, newTodo.ID)
	assert.Equal(t, updatedTodo.Title, newTitle)
	assert.Equal(t, updatedTodo.Done, false)
	assert.Equal(t, updatedTodo.Body, newTodo.Body)
}

func TestGetTitle(t *testing.T) {
	// GIVEN
	testTitle := "Test Title"
	todo := &Todo{
		ID:    "ID",
		Title: testTitle,
		Body:  "Some Body",
		Done:  true,
	}

	// WHEN
	title := todo.GetTitle()

	// THEN
	assert.Equal(t, title, testTitle)
}

func TestGet(t *testing.T) {
	// GIVEN
	todo := &Todo{
		ID:    "ID",
		Title: "Some Title",
		Body:  "Some Body",
		Done:  true,
	}

	// WHEN
	gotTodo := todo.Get()

	// THEN
	assert.Equal(t, gotTodo, *todo)
}
