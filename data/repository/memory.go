package repository

import "context"

type MemoryRepository[Model Identifiable] struct {
	IRepository[Model]

	Data map[string]*Model
}

func CreateMemoryRepository[Model Identifiable](initialData []Model) IRepository[Model] {
	repo := MemoryRepository[Model]{
		Data: map[string]*Model{},
	}

	for _, data := range initialData {
		repo.Data[data.GetID()] = &data
	}
	return &repo
}

func (r MemoryRepository[Model]) GetAll(context.Context) ([]*Model, error) {
	models := []*Model{}
	for _, data := range r.Data {
		models = append(models, data)
	}
	return models, nil
}

func (r MemoryRepository[Model]) GetById(ctx context.Context, id string) (*Model, error) {
	existingTodo := r.Data[id]

	if existingTodo == nil {
		return nil, NotFound{}
	}
	return existingTodo, nil
}

func (r MemoryRepository[Model]) Create(ctx context.Context, model Model) error {
	r.Data[model.GetID()] = &model
	return nil
}

func (r *MemoryRepository[Model]) Save(ctx context.Context, model Model) error {
	_, err := r.GetById(ctx, model.GetID())
	if err != nil {
		return err
	}
	r.Data[model.GetID()] = &model
	return nil
}
