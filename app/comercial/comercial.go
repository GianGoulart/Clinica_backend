package comercial

import (
	"context"
	"fmt"
	"time"

	"github.com/GianGoulart/Clinica_backend/model"
	"github.com/GianGoulart/Clinica_backend/store"
)

type App interface {
	GetAll(ctx context.Context) (*[]model.Comercial, error)
	GetById(ctx context.Context, id string) (*model.Comercial, error)
	GetByAnything(ctx context.Context, comercial *model.Comercial) (*[]model.Comercial, error)
	Set(ctx context.Context, comercial *model.Comercial) (*model.Comercial, error)
	Update(ctx context.Context, comercial *model.Comercial) (*model.Comercial, error)
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

func (s appImpl) GetAll(ctx context.Context) (*[]model.Comercial, error) {
	res, err := s.store.Comercial.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var comercials = []model.Comercial{}
	for _, r := range *res {
		p := r.PreencheComercial(&r)

		medico_part, err := s.store.Medico.GetById(ctx, p.Id_Medico_Part)
		if err != nil {
			return nil, err
		}
		p.Nome_Medico_Part = medico_part.Nome

		procedimento, err := s.store.Procedimento.GetById(ctx, p.Id_Procedimento)
		if err != nil {
			return nil, err
		}
		procedimento = procedimento.PreencheProcedimentos(procedimento)

		p.Desc_Procedimento = fmt.Sprintf("%s - %s - %s - %v", procedimento.Nome_Paciente, procedimento.Nome_Medico, procedimento.NomeProcedimento, time.Unix(procedimento.Data, 0).Format("02-01-2006"))
		comercials = append(comercials, *p)
	}

	res = &comercials
	return res, nil
}

func (s appImpl) GetById(ctx context.Context, id string) (*model.Comercial, error) {
	res, err := s.store.Comercial.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	res = res.PreencheComercial(res)
	medico_part, err := s.store.Medico.GetById(ctx, res.Id_Medico_Part)
	if err != nil {
		return nil, err
	}
	res.Nome_Medico_Part = medico_part.Nome

	procedimento, err := s.store.Procedimento.GetById(ctx, res.Id_Procedimento)
	if err != nil {
		return nil, err
	}
	procedimento = procedimento.PreencheProcedimentos(procedimento)

	res.Desc_Procedimento = fmt.Sprintf("%s - %s - %s - %v", procedimento.Nome_Paciente, procedimento.Nome_Medico, procedimento.NomeProcedimento, time.Unix(procedimento.Data, 0).Format("02-01-2006"))

	return res, nil
}

func (s appImpl) GetByAnything(ctx context.Context, comercial *model.Comercial) (*[]model.Comercial, error) {
	res, err := s.store.Comercial.GetByAnything(ctx, comercial)
	if err != nil {
		return nil, err
	}
	var comercials = []model.Comercial{}
	for _, r := range *res {
		p := r.PreencheComercial(&r)

		medico_part, err := s.store.Medico.GetById(ctx, p.Id_Medico_Part)
		if err != nil {
			return nil, err
		}
		p.Nome_Medico_Part = medico_part.Nome

		procedimento, err := s.store.Procedimento.GetById(ctx, p.Id_Procedimento)
		if err != nil {
			return nil, err
		}
		procedimento = procedimento.PreencheProcedimentos(procedimento)

		p.Desc_Procedimento = fmt.Sprintf("%s - %s - %s - %v", procedimento.Nome_Paciente, procedimento.Nome_Medico, procedimento.NomeProcedimento, time.Unix(procedimento.Data, 0).Format("02-01-2006"))
		comercials = append(comercials, *p)
	}

	res = &comercials
	return res, nil
}

func (s appImpl) Set(ctx context.Context, comercial *model.Comercial) (*model.Comercial, error) {
	comercial.Id = model.NewId()

	res, err := s.store.Comercial.Set(ctx, comercial)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s appImpl) Update(ctx context.Context, comercial *model.Comercial) (*model.Comercial, error) {

	res, err := s.store.Comercial.Update(ctx, comercial)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s appImpl) Delete(ctx context.Context, id string) error {
	err := s.store.Comercial.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
