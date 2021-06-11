package financeiro

import (
	"context"
	"fmt"

	"github.com/GianGoulart/Clinica_backend/app/comercial"
	"github.com/GianGoulart/Clinica_backend/model"
	"github.com/GianGoulart/Clinica_backend/store"
	"github.com/GianGoulart/Clinica_backend/store/financeiro"
)

type App interface {
	GetAll(ctx context.Context) (*[]model.Financeiro, error)
	GetById(ctx context.Context, id string) (*model.Financeiro, error)
	GetByAnything(ctx context.Context, financeiro *model.Financeiro) (*[]model.Financeiro, error)
	Set(ctx context.Context, financeiro *model.Financeiro) (*model.Financeiro, error)
	Update(ctx context.Context, financeiro *model.Financeiro) (*model.Financeiro, error)
	Delete(ctx context.Context, id string) error
}

func NewApp(stores *store.Container, comercial comercial.App) App {
	return appImpl{
		store:     stores.Financeiro,
		comercial: comercial,
	}
}

type appImpl struct {
	store     financeiro.Store
	comercial comercial.App
}

func (s appImpl) GetAll(ctx context.Context) (*[]model.Financeiro, error) {
	res, err := s.store.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var financeiros = []model.Financeiro{}
	for _, r := range *res {
		p := r.PreencheFinanceiro(&r)

		c, _ := s.comercial.GetById(ctx, p.Id_Comercial)

		p.Desc_Comercial = fmt.Sprintf("%s - %s", c.Desc_Procedimento, c.Nome_Medico_Part)

		financeiros = append(financeiros, *p)
	}

	res = &financeiros
	return res, nil
}

func (s appImpl) GetById(ctx context.Context, id string) (*model.Financeiro, error) {
	res, err := s.store.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	res = res.PreencheFinanceiro(res)

	return res, nil
}

func (s appImpl) GetByAnything(ctx context.Context, financeiro *model.Financeiro) (*[]model.Financeiro, error) {
	res, err := s.store.GetByAnything(ctx, financeiro)
	if err != nil {
		return nil, err
	}
	var financeiros = []model.Financeiro{}
	for _, r := range *res {
		p := r.PreencheFinanceiro(&r)
		c, _ := s.comercial.GetById(ctx, p.Id_Comercial)

		p.Desc_Comercial = fmt.Sprintf("%s - %s", c.Desc_Procedimento, c.Nome_Medico_Part)

		financeiros = append(financeiros, *p)
	}

	res = &financeiros
	return res, nil
}

func (s appImpl) Set(ctx context.Context, financeiro *model.Financeiro) (*model.Financeiro, error) {
	financeiro.Id = model.NewId()

	res, err := s.store.Set(ctx, financeiro)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s appImpl) Update(ctx context.Context, financeiro *model.Financeiro) (*model.Financeiro, error) {

	res, err := s.store.Update(ctx, financeiro)
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
