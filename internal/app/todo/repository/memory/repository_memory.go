package memory

import (
	"sync"

	"github.com/akbarpambudi/todo/internal/app/domain"
)

type TodoRepositoryInMemory struct {
	datastore map[string][]domain.TodoEvent
	mux       sync.Mutex
}

func New() *TodoRepositoryInMemory {
	repo := new(TodoRepositoryInMemory)
	repo.init()
	return repo
}

func (repo *TodoRepositoryInMemory) init() {
	repo.datastore = map[string][]domain.TodoEvent{}
}

func (repo *TodoRepositoryInMemory) Load(id string) (domain.TodoAggregate, error) {
	todo := new(domain.Todo)
	events := repo.datastore[id]
	if events == nil {
		events = []domain.TodoEvent{}
	}
	for _, event := range events {

		todo.ApplyEvent(event)
	}
	return todo, nil
}

func (repo *TodoRepositoryInMemory) Save(todo domain.TodoAggregate) {
	repo.mux.Lock()
	defer repo.mux.Unlock()
	events := repo.datastore[todo.GetID()]
	if events == nil {
		events = []domain.TodoEvent{}
	}
	repo.datastore[todo.GetID()] = append(events, todo.GetUnCommittedEvents()...)
	todo.CommitEvents()
}
