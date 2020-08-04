package chi

import (
	"encoding/json"
	"net/http"
)

type CreateTodoRequest struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (payload *CreateTodoRequest) Bind(req *http.Request) error {
	if err := json.NewDecoder(req.Body).Decode(payload); err != nil {
		return err
	}
	return nil
}

type CreateTodoResponse struct{}

func (resp *CreateTodoResponse) Render(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(202)
	return nil
}

type GetTodoResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func (resp *GetTodoResponse) Render(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(200)
	return nil
}
