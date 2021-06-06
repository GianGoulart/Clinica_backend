package medicos

import (
	"context"

	"github.com/GianGoulart/Clinica_backend/model"
	"github.com/GianGoulart/Clinica_backend/store"
	"github.com/GianGoulart/Clinica_backend/store/medicos"
)

type App interface {
	GetAll(ctx context.Context) (*[]model.Medico, error)
	GetById(ctx context.Context, id string) (*model.Medico, error)
	GetByAnything(ctx context.Context, medico *model.Medico) (*[]model.Medico, error)
	Set(ctx context.Context, medico *model.Medico) (*model.Medico, error)
	Update(ctx context.Context, medico *model.Medico) (*model.Medico, error)
	Delete(ctx context.Context, id string) error
}

func NewApp(stores *store.Container) App {
	return appImpl{
		store: stores.Medico,
	}
}

type appImpl struct {
	store medicos.Store
}

func (s appImpl) GetAll(ctx context.Context) (*[]model.Medico, error) {
	res, err := s.store.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s appImpl) GetById(ctx context.Context, id string) (*model.Medico, error) {
	res, err := s.store.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s appImpl) GetByAnything(ctx context.Context, medico *model.Medico) (*[]model.Medico, error) {
	res, err := s.store.GetByAnything(ctx, medico)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s appImpl) Set(ctx context.Context, medico *model.Medico) (*model.Medico, error) {
	medico.Id = model.NewId()

	if err := medico.Validate(); err != nil {
		return nil, err

	}
	res, err := s.store.Set(ctx, medico)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s appImpl) Update(ctx context.Context, medico *model.Medico) (*model.Medico, error) {
	if err := medico.Validate(); err != nil {
		return nil, err

	}

	res, err := s.store.Update(ctx, medico)
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
