package pacientes

import (
	"context"

	"github.com/GianGoulart/Clinica_backend/model"
	"github.com/GianGoulart/Clinica_backend/store"
	"github.com/GianGoulart/Clinica_backend/store/pacientes"
)

type App interface {
	GetAll(ctx context.Context) (*[]model.Paciente, error)
	GetById(ctx context.Context, id string) (*model.Paciente, error)
	GetByAnything(ctx context.Context, paciente *model.Paciente) (*[]model.Paciente, error)
	Set(ctx context.Context, paciente *model.Paciente) (*model.Paciente, error)
	Update(ctx context.Context, paciente *model.Paciente) (*model.Paciente, error)
	Delete(ctx context.Context, id string) error
}

func NewApp(stores *store.Container) App {
	return appImpl{
		store: stores.Paciente,
	}
}

type appImpl struct {
	store pacientes.Store
}

func (s appImpl) GetAll(ctx context.Context) (*[]model.Paciente, error) {
	res, err := s.store.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s appImpl) GetById(ctx context.Context, id string) (*model.Paciente, error) {
	res, err := s.store.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s appImpl) GetByAnything(ctx context.Context, paciente *model.Paciente) (*[]model.Paciente, error) {
	res, err := s.store.GetByAnything(ctx, paciente)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s appImpl) Set(ctx context.Context, paciente *model.Paciente) (*model.Paciente, error) {
	paciente.Id = model.NewId()

	if err := paciente.Validate(); err != nil {
		return nil, err

	}
	res, err := s.store.Set(ctx, paciente)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s appImpl) Update(ctx context.Context, paciente *model.Paciente) (*model.Paciente, error) {
	if err := paciente.Validate(); err != nil {
		return nil, err

	}

	res, err := s.store.Update(ctx, paciente)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s appImpl) Delete(ctx context.Context, id string) error {
	err := s.store.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
