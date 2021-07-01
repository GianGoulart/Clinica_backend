package acompanhamento

import (
	"context"
	"fmt"
	"time"

	"github.com/GianGoulart/Clinica_backend/model"
	"github.com/GianGoulart/Clinica_backend/store"
)

type App interface {
	GetAll(ctx context.Context) (*[]model.Acompanhamento, error)
	GetById(ctx context.Context, id string) (*model.Acompanhamento, error)
	GetByIdProcedimento(ctx context.Context, id string) (*model.Acompanhamento, error)
	GetByAnything(ctx context.Context, acompanhamento *model.Acompanhamento) (*[]model.Acompanhamento, error)
	Set(ctx context.Context, acompanhamento *model.Acompanhamento) (*model.Acompanhamento, error)
	Update(ctx context.Context, acompanhamento *model.Acompanhamento) (*model.Acompanhamento, error)
	Delete(ctx context.Context, id string) error
}

func NewApp(stores *store.Container) App {
	return appImpl{
		store: stores,
	}
}

type appImpl struct {
	store *store.Container
}

func (s appImpl) GetAll(ctx context.Context) (*[]model.Acompanhamento, error) {
	res, err := s.store.Acompanhamento.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var acompanhamentos = []model.Acompanhamento{}
	for _, r := range *res {
		p := r.PreencheAcompanhamento(&r)

		procedimento, err := s.store.Procedimento.GetById(ctx, p.Id_Procedimento)
		if err != nil {
			return nil, err
		}
		procedimento = procedimento.PreencheProcedimentos(procedimento)

		p.Desc_Procedimento = fmt.Sprintf("%s - %s - %s - %v", procedimento.Nome_Paciente, procedimento.Nome_Medico, procedimento.NomeProcedimento, time.Unix(procedimento.Data, 0).Format("02-01-2006"))
		acompanhamentos = append(acompanhamentos, *p)
	}

	res = &acompanhamentos
	return res, nil
}

func (s appImpl) GetById(ctx context.Context, id string) (*model.Acompanhamento, error) {
	res, err := s.store.Acompanhamento.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	res = res.PreencheAcompanhamento(res)

	procedimento, err := s.store.Procedimento.GetById(ctx, res.Id_Procedimento)
	if err != nil {
		return nil, err
	}
	procedimento = procedimento.PreencheProcedimentos(procedimento)

	res.Desc_Procedimento = fmt.Sprintf("%s - %s - %s - %v", procedimento.Nome_Paciente, procedimento.Nome_Medico, procedimento.NomeProcedimento, time.Unix(procedimento.Data, 0).Format("02-01-2006"))

	return res, nil
}

func (s appImpl) GetByIdProcedimento(ctx context.Context, id string) (*model.Acompanhamento, error) {
	res, err := s.store.Acompanhamento.GetByIdProcedimento(ctx, id)
	if err != nil {
		return nil, err
	}
	res = res.PreencheAcompanhamento(res)

	procedimento, err := s.store.Procedimento.GetById(ctx, res.Id_Procedimento)
	if err != nil {
		return nil, err
	}
	procedimento = procedimento.PreencheProcedimentos(procedimento)

	res.Desc_Procedimento = fmt.Sprintf("%s - %s - %s - %v", procedimento.Nome_Paciente, procedimento.Nome_Medico, procedimento.NomeProcedimento, time.Unix(procedimento.Data, 0).Format("02-01-2006"))

	return res, nil
}

func (s appImpl) GetByAnything(ctx context.Context, acompanhamento *model.Acompanhamento) (*[]model.Acompanhamento, error) {
	res, err := s.store.Acompanhamento.GetByAnything(ctx, acompanhamento)
	if err != nil {
		return nil, err
	}
	var acompanhamentos = []model.Acompanhamento{}
	for _, r := range *res {
		p := r.PreencheAcompanhamento(&r)

		procedimento, err := s.store.Procedimento.GetById(ctx, p.Id_Procedimento)
		if err != nil {
			return nil, err
		}
		procedimento = procedimento.PreencheProcedimentos(procedimento)

		p.Desc_Procedimento = fmt.Sprintf("%s - %s - %s - %v", procedimento.Nome_Paciente, procedimento.Nome_Medico, procedimento.NomeProcedimento, time.Unix(procedimento.Data, 0).Format("02-01-2006"))
		acompanhamentos = append(acompanhamentos, *p)
	}

	res = &acompanhamentos
	return res, nil
}

func (s appImpl) Set(ctx context.Context, acompanhamento *model.Acompanhamento) (*model.Acompanhamento, error) {
	acompanhamento.Id = model.NewId()

	res, err := s.store.Acompanhamento.Set(ctx, acompanhamento)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s appImpl) Update(ctx context.Context, acompanhamento *model.Acompanhamento) (*model.Acompanhamento, error) {

	res, err := s.store.Acompanhamento.Update(ctx, acompanhamento)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s appImpl) Delete(ctx context.Context, id string) error {
	err := s.store.Acompanhamento.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
