package domain

import (
	"encoding/json"
	"time"
)

type TodoState struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	Status      TodoStatus `json:"status"`
}

func (state TodoState) Equals(stateInCheck TodoState) bool {
	return state == stateInCheck
}

type TodoEvent interface {
	GetState() TodoState
	GetID() string
	GetEventName() string
	BindFromJson([]byte) error
	ToJson() ([]byte, error)
	Equals(event TodoEvent) bool
}

type TodoEvents []TodoEvent

func (events TodoEvents) Contains(todoEvent TodoEvent) (isContains bool) {
	for _, event := range events {
		isContains = todoEvent.Equals(event)
		if isContains {
			return
		}
	}
	return
}

type TodoCreated struct {
	state TodoState
}

func NewTodoCreatedEvent(state TodoState) *TodoCreated {
	return &TodoCreated{
		state: state,
	}
}

func (event TodoCreated) Equals(eventInCheck TodoEvent) bool {
	return event.GetState().Equals(eventInCheck.GetState())
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

func NewTodoRevisedEvent(state TodoState) *TodoRevised {
	return &TodoRevised{
		state: state,
	}
}

func (event TodoRevised) Equals(eventInCheck TodoEvent) bool {
	return event.GetState().Equals(eventInCheck.GetState())
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

func NewTodoMovedEvent(state TodoState) *TodoMoved {
	return &TodoMoved{state: state}
}

func (event TodoMoved) Equals(eventInCheck TodoEvent) bool {
	return event.GetState().Equals(eventInCheck.GetState())
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
