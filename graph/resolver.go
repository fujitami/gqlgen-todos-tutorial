package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//go:generate go run github.com/99designs/gqlgen

import (
	"gqlgen-todos-tutorial/graph/model"
)

type Resolver struct{
	todos []*model.Todo
}
