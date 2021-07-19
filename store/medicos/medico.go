package medicos

import (
	"context"
	"database/sql"

	log "github.com/sirupsen/logrus"

	"github.com/GianGoulart/Clinica_backend/model"
	"github.com/jmoiron/sqlx"
)

type Store interface {
	GetAll(ctx context.Context) (*[]model.Medico, error)
	GetById(ctx context.Context, id string) (*model.Medico, error)
	GetByAnything(ctx context.Context, Medico *model.Medico) (*[]model.Medico, error)
	Set(ctx context.Context, Medico *model.Medico) (*model.Medico, error)
	Update(ctx context.Context, Medico *model.Medico) (*model.Medico, error)
	Delete(ctx context.Context, id string) error
}

func NewStore(db *sqlx.DB) Store {
	return &storeImpl{db}
}

type storeImpl struct {
	db *sqlx.DB
}

func (s *storeImpl) GetAll(ctx context.Context) (*[]model.Medico, error) {
	medicos := new([]model.Medico)
	query := `
			Select id, nome, especialidade
				From 
					BD_ClinicaAbrao.medicos
				Order by nome`

	err := s.db.SelectContext(ctx, medicos, query)
	if err != nil && err != sql.ErrNoRows {
		log.WithContext(ctx).Error("store.medicos.get_medicos_all ", err.Error())
		return nil, err
	}

	return medicos, nil
}

func (s *storeImpl) GetById(ctx context.Context, id string) (*model.Medico, error) {
	medico := new(model.Medico)
	query := `
			Select id, nome, especialidade
				From 
					BD_ClinicaAbrao.medicos 
				Where 
					id = ? 
				Order by nome`

	err := s.db.GetContext(ctx, medico, query, id)
	if err != nil && err != sql.ErrNoRows {
		log.WithContext(ctx).Error("store.medico.get_medico_by_id ", err.Error())
		return nil, err
	}

	return medico, nil
}

func (s *storeImpl) GetByAnything(ctx context.Context, medico *model.Medico) (*[]model.Medico, error) {
	medicos := new([]model.Medico)
	query := `
			Select id, nome, especialidade
				From 
					BD_ClinicaAbrao.medicos 
				Where 
					1 = 1 `

	if len(medico.Nome) > 0 {
		query += `and nome like '%` + medico.Nome + `%' `
	}
	if len(medico.Especialidade) > 0 {
		query += `and especialidade like '%` + medico.Especialidade + `%' `
	}

	err := s.db.SelectContext(ctx, medicos, query)
	if err != nil && err != sql.ErrNoRows {
		log.WithContext(ctx).Error("store.medico.get_medico_by_anything ", err.Error())
		return nil, err
	}

	return medicos, nil
}

func (s *storeImpl) Set(ctx context.Context, medico *model.Medico) (*model.Medico, error) {

	_, err := s.db.ExecContext(ctx, `INSERT INTO BD_ClinicaAbrao.medicos (id, nome, especialidade) VALUES (?,?,?)`,
		medico.Id, medico.Nome, medico.Especialidade)
	if err != nil {
		log.WithContext(ctx).Error("store.medico.set_paciente", err.Error())
		return nil, err
	}

	return medico, nil
}

func (s *storeImpl) Update(ctx context.Context, medico *model.Medico) (*model.Medico, error) {
	_, err := s.db.ExecContext(ctx, `Update BD_ClinicaAbrao.medicos SET nome=?, especialidade=? Where id=?`,
		medico.Nome, medico.Especialidade, medico.Id)
	if err != nil {
		log.WithContext(ctx).Error("store.medico.update", err.Error())
		return nil, err
	}

	return medico, nil
}

func (s *storeImpl) Delete(ctx context.Context, id string) error {
	_, err := s.db.ExecContext(ctx, `Delete FROm BD_ClinicaAbrao.medicos Where id=?`, id)
	if err != nil {
		log.WithContext(ctx).Error("store.medico.delete", err.Error())
		return err
	}

	return nil
}
