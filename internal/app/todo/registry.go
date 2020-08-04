package todo

import (
	"github.com/akbarpambudi/todo/internal/app/domain"
	todoChi "github.com/akbarpambudi/todo/internal/app/todo/handler/web/chi"
	"github.com/akbarpambudi/todo/internal/app/todo/repository/memory"

	"github.com/go-chi/chi"
	"go.uber.org/dig"
)

func Register(container *dig.Container) {
	container.Provide(func() domain.TodoRepository {
		return memory.New()
	})

	container.Provide(func(repository domain.TodoRepository) *todoChi.TodoHandlerWebChi {
		handler := todoChi.NewTodoHandlerWebChi(repository)
		return handler
	})

}

func Invoke(container *dig.Container) {
	container.Invoke(func(router chi.Router, handler *todoChi.TodoHandlerWebChi) {
		handler.Bind(router)
	})
}
