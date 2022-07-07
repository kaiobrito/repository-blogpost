package repository

type Identifiable interface {
	GetID() string
}

type IRepository[Model Identifiable] interface {
	GetAll() ([]*Model, error)
	GetById(string) (*Model, error)
	Create(Model) error
	Save(Model) error
}
