package domain_test

import (
	"testing"

	"github.com/akbarpambudi/todo/internal/app/domain"
	"github.com/stretchr/testify/assert"
)

func TestTodo_NewTodo_should_ReturnTodoWithRightIDTitleAndDescription(t *testing.T) {
	expectedTodoID := "3d641cde-6bde-4f10-bcb7-f83de4b20deb"
	expectedTodoTitle := "Read a book"
	expectedTodoDescription := "Read clean code book"
	todo := domain.NewTodo(expectedTodoID, expectedTodoTitle, expectedTodoDescription)
	assert.Equal(t, expectedTodoID, todo.GetID())
	assert.Equal(t, expectedTodoTitle, todo.GetTitle())
	assert.Equal(t, expectedTodoDescription, todo.GetDescription())
}

func TestTodo_NewTodo_should_ReturnTodoWithStatusNew(t *testing.T) {
	todoId := "3d641cde-6bde-4f10-bcb7-f83de4b20deb"
	todoTitle := "Read a book"
	todoDescription := "Read clean code book"
	todo := domain.NewTodo(todoId, todoTitle, todoDescription)
	assert.Equal(t, domain.TodoStatusNew, todo.GetStatus())
}

func TestTodo_NewTodo_should_AddTodoCreatedEvent(t *testing.T) {
	todoId := "3d641cde-6bde-4f10-bcb7-f83de4b20deb"
	todoTitle := "Read a book"
	todoDescription := "Read clean code book"

	todo := domain.NewTodo(todoId, todoTitle, todoDescription)
	assert.Condition(t, func() (success bool) {
		success = len(todo.GetUnCommittedEvents()) == 1
		return
	})
}
