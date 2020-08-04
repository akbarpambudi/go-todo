package domain

import (
	"encoding/json"
	"errors"
	"strings"
	"time"
)

var (
	ErrTodoNotFound error = errors.New("todo not found")
)

type TodoStatus int

const (
	TodoStatusUnknown TodoStatus = iota
	TodoStatusNew
	TodoStatusInProgress
	TodoStatusDone
)

func (status TodoStatus) IsDone() bool {
	return status == TodoStatusDone
}

func (status TodoStatus) IsNew() bool {
	return status == TodoStatusNew
}

func (status TodoStatus) IsInProgress() bool {
	return status == TodoStatusInProgress
}

type TodoRevision struct {
	Title       string
	Description string
}

type TodoState struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	Status      TodoStatus `json:"status"`
}

type TodoEvent interface {
	GetState() TodoState
	GetID() string
	GetEventName() string
	BindFromJson([]byte) error
	ToJson() ([]byte, error)
}

type TodoCreated struct {
	state TodoState
}

func (event TodoCreated) GetState() TodoState {
	return event.state
}

func (event TodoCreated) GetID() string {
	return event.state.ID
}

func (event TodoCreated) GetEventName() string {
	return "todo.event.todo-created"
}

func (event *TodoCreated) BindFromJson(eventInJson []byte) error {
	return json.Unmarshal(eventInJson, event)
}

func (event *TodoCreated) ToJson() ([]byte, error) {
	return json.Marshal(event)
}

type TodoRevised struct {
	state TodoState
}

func (event TodoRevised) GetState() TodoState {
	return event.state
}

func (event TodoRevised) GetID() string {
	return event.state.ID
}

func (event TodoRevised) GetEventName() string {
	return "todo.event.todo-revised"
}

func (event *TodoRevised) BindFromJson(eventInJson []byte) error {
	return json.Unmarshal(eventInJson, event)
}

func (event *TodoRevised) ToJson() ([]byte, error) {
	return json.Marshal(event)
}

type TodoMoved struct {
	state TodoState
}

func (event TodoMoved) GetState() TodoState {
	return event.state
}

func (event TodoMoved) GetID() string {
	return event.state.ID
}

func (event TodoMoved) GetEventName() string {
	return "todo.event.todo-moved"
}

func (event *TodoMoved) BindFromJson(eventInJson []byte) error {
	return json.Unmarshal(eventInJson, event)
}

func (event *TodoMoved) ToJson() ([]byte, error) {
	return json.Marshal(event)
}

type TodoAggregate interface {
	GetID() string
	GetTitle() string
	GetDescription() string
	GetCreatedAt() time.Time
	GetStatus() TodoStatus
	addEvent(event TodoEvent)
	GetUnCommittedEvents() []TodoEvent
	CommitEvents()
	MoveToInProgress()
	MoveToDone()
	Revise(revision TodoRevision)
}

type Todo struct {
	id             string
	title          string
	description    string
	createdAt      time.Time
	status         TodoStatus
	unCommitEvents []TodoEvent
}

func NewTodo(id string, title string, description string) *Todo {
	todo := new(Todo)
	todo.addEvent(&TodoCreated{
		state: TodoState{
			ID:          id,
			Title:       title,
			Description: description,
			CreatedAt:   time.Now(),
			Status:      TodoStatusNew,
		},
	})
	return todo
}

func (todo *Todo) MoveToInProgress() {
	if !todo.status.IsNew() {
		return
	}
	todo.addEvent(&TodoMoved{
		state: TodoState{
			ID:     todo.id,
			Status: TodoStatusInProgress,
		},
	})
}

func (todo *Todo) MoveToDone() {
	if !todo.status.IsInProgress() {
		return
	}
	todo.addEvent(&TodoMoved{
		state: TodoState{
			ID:     todo.id,
			Status: TodoStatusDone,
		},
	})
}

func (todo *Todo) Revise(revision TodoRevision) {
	state := TodoState{}
	if strings.TrimSpace(revision.Title) != "" {
		state.Title = revision.Title
	}

	if strings.TrimSpace(revision.Description) != "" {
		state.Description = revision.Description
	}
	if (state == TodoState{}) {
		return
	}
	state.ID = todo.id
	todo.addEvent(&TodoRevised{
		state: state,
	})
}

func (todo Todo) GetTitle() string {
	return todo.title
}

func (todo Todo) GetDescription() string {
	return todo.description
}

func (todo Todo) GetID() string {
	return todo.id
}

func (todo Todo) GetCreatedAt() time.Time {
	return todo.createdAt
}

func (todo Todo) GetStatus() TodoStatus {
	return todo.status
}

func (todo *Todo) addEvent(event TodoEvent) {
	todo.unCommitEvents = append(todo.unCommitEvents, event)
	todo.ApplyEvent(event)
}

func (todo *Todo) CommitEvents() {
	todo.unCommitEvents = []TodoEvent{}
}

func (todo *Todo) GetUnCommittedEvents() []TodoEvent {
	return todo.unCommitEvents
}

func (todo *Todo) ApplyEvent(event TodoEvent) error {
	switch v := event.(type) {
	case *TodoCreated:
		todo.id = v.GetID()
		todo.title = v.GetState().Title
		todo.description = v.GetState().Description
		todo.status = v.GetState().Status
		todo.createdAt = v.GetState().CreatedAt
		return nil
	case *TodoMoved:
		todo.status = v.GetState().Status
		return nil
	case *TodoRevised:
		todo.description = v.GetState().Description
		todo.title = v.GetState().Title
		return nil
	}
	return errors.New("unknown event")
}

type TodoRepository interface {
	Load(id string) (TodoAggregate, error)
	Save(aggregate TodoAggregate)
}
