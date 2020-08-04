package chi

import (
	"net/http"

	"github.com/akbarpambudi/todo/internal/app/domain"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type TodoHandlerWebChi struct {
	repository domain.TodoRepository
}

func NewTodoHandlerWebChi(repository domain.TodoRepository) *TodoHandlerWebChi {
	return &TodoHandlerWebChi{repository: repository}
}

func (handler TodoHandlerWebChi) Bind(router chi.Router) {
	router.Get("/todo/{id}", handler.HandleGetTodo)
	router.Post("/todo", handler.HandleCreateTodo)
}

func (handler TodoHandlerWebChi) HandleCreateTodo(w http.ResponseWriter, r *http.Request) {
	request := new(CreateTodoRequest)
	request.Bind(r)
	todo := domain.NewTodo(request.ID, request.Title, request.Description)

	handler.repository.Save(todo)
}

func (handler TodoHandlerWebChi) HandleMoveToDone(w http.ResponseWriter, r *http.Request) {

}

func (handler TodoHandlerWebChi) HandleMoveToInProgress(w http.ResponseWriter, r *http.Request) {
}

func (handler TodoHandlerWebChi) HandleReviseTodo(w http.ResponseWriter, r *http.Request) {
}

func (handler TodoHandlerWebChi) HandleGetTodo(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	todo, err := handler.repository.Load(id)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("something went wrong"))
		return
	}
	render.Render(w, r, &GetTodoResponse{
		ID:          todo.GetID(),
		Description: todo.GetDescription(),
		Status:      "",
		Title:       todo.GetTitle(),
	})
}
