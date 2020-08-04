package web

import "net/http"

type TodoHandlerWeb interface {
	HandleCreateTodo(w http.ResponseWriter, r *http.Request)
	HandleMoveToDone(w http.ResponseWriter, r *http.Request)
	HandleMoveToInProgress(w http.ResponseWriter, r *http.Request)
	HandleReviseTodo(w http.ResponseWriter, r *http.Request)
	HandleGetTodo(w http.ResponseWriter, r *http.Request)
}
