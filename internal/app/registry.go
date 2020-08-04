package app

import (
	"log"
	"net/http"

	"github.com/akbarpambudi/todo/internal/app/todo"

	"github.com/go-chi/chi"
	"go.uber.org/dig"
)

func Register(container *dig.Container) {
	todo.Register(container)
	container.Provide(func() chi.Router {
		return chi.NewRouter()
	})
}

func Invoke(container *dig.Container) {
	todo.Invoke(container)
	err := container.Invoke(func(router chi.Router) error {
		return http.ListenAndServe(":8080", router)
	})

	if err != nil {
		log.Fatal(err)
	}
}
