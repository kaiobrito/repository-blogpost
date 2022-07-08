package repository

import "context"

type Identifiable interface {
	GetID() string
}

type IRepository[Model Identifiable] interface {
	GetAll(context.Context) ([]*Model, error)
	GetById(context.Context, string) (*Model, error)
	Create(context.Context, Model) error
	Save(context.Context, Model) error
}
