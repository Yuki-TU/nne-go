package graph

//go:generate go run -mod=mod github.com/99designs/gqlgen generate

import "github.com/Yuki-TU/nne-go/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// リゾルバーのルート部分

type Resolver struct{
	todos []*model.Todo
}
