package procedimentos

import (
	"context"

	"github.com/GianGoulart/Clinica_backend/model"
	"github.com/GianGoulart/Clinica_backend/store"
	"github.com/GianGoulart/Clinica_backend/store/procedimentos"
)

type App interface {
	GetAll(ctx context.Context) (*[]model.Procedimento, error)
	GetById(ctx context.Context, id string) (*model.Procedimento, error)
	GetByAnything(ctx context.Context, procedimento *model.Procedimento) (*[]model.Procedimento, error)
	Set(ctx context.Context, procedimento *model.Procedimento) (*model.Procedimento, error)
	Update(ctx context.Context, procedimento *model.Procedimento) (*model.Procedimento, error)
	Delete(ctx context.Context, id string) error
}

func NewApp(stores *store.Container) App {
	return appImpl{
		store: stores.Procedimento,
	}
}

type appImpl struct {
	store procedimentos.Store
}

func (s appImpl) GetAll(ctx context.Context) (*[]model.Procedimento, error) {
	res, err := s.store.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var procedimentos = []model.Procedimento{}
	for _, r := range *res {
		p := r.PreencheProcedimentos(&r)
		procedimentos = append(procedimentos, *p)
	}

	res = &procedimentos
	return res, nil
}

func (s appImpl) GetById(ctx context.Context, id string) (*model.Procedimento, error) {
	res, err := s.store.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	res = res.PreencheProcedimentos(res)

	return res, nil
}

func (s appImpl) GetByAnything(ctx context.Context, procedimento *model.Procedimento) (*[]model.Procedimento, error) {
	res, err := s.store.GetByAnything(ctx, procedimento)
	if err != nil {
		return nil, err
	}
	var procedimentos = []model.Procedimento{}
	for _, r := range *res {
		p := r.PreencheProcedimentos(&r)
		procedimentos = append(procedimentos, *p)
	}

	res = &procedimentos
	return res, nil
}

func (s appImpl) Set(ctx context.Context, procedimento *model.Procedimento) (*model.Procedimento, error) {
	procedimento.Id = model.NewId()

	if err := procedimento.Validate(); err != nil {
		return nil, err

	}
	res, err := s.store.Set(ctx, procedimento)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s appImpl) Update(ctx context.Context, procedimento *model.Procedimento) (*model.Procedimento, error) {
	if err := procedimento.Validate(); err != nil {
		return nil, err

	}

	res, err := s.store.Update(ctx, procedimento)
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
