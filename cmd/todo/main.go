package main

import (
	app "github.com/akbarpambudi/todo/internal/app"

	"go.uber.org/dig"
)

var (
	container = dig.New()
)

func init() {
	app.Register(container)
}

func main() {
	app.Invoke(container)
}
