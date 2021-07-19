package pacientes

import (
	"context"
	"database/sql"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/GianGoulart/Clinica_backend/model"
	"github.com/jmoiron/sqlx"
)

type Store interface {
	GetAll(ctx context.Context) (*[]model.Paciente, error)
	GetById(ctx context.Context, id string) (*model.Paciente, error)
	GetByAnything(ctx context.Context, paciente *model.Paciente) (*[]model.Paciente, error)
	Set(ctx context.Context, paciente *model.Paciente) (*model.Paciente, error)
	Update(ctx context.Context, paciente *model.Paciente) (*model.Paciente, error)
	Delete(ctx context.Context, id string) error
}

func NewStore(db *sqlx.DB) Store {
	return &storeImpl{db}
}

type storeImpl struct {
	db *sqlx.DB
}

func (s *storeImpl) GetAll(ctx context.Context) (*[]model.Paciente, error) {
	pacientes := new([]model.Paciente)
	query := `
			Select id, cpf, nome, telefone, telefone2, convenio, plano, acomodacao, status
				From 
					BD_ClinicaAbrao.pacientes
				Order by nome`

	err := s.db.SelectContext(ctx, pacientes, query)
	if err != nil && err != sql.ErrNoRows {
		log.WithContext(ctx).Error("store.paciente.get_paciente_all ", err.Error())
		return nil, err
	}

	return pacientes, nil
}

func (s *storeImpl) GetById(ctx context.Context, id string) (*model.Paciente, error) {
	paciente := new(model.Paciente)
	query := `
			Select id, cpf, nome, telefone, telefone2, convenio, plano, acomodacao, status
				From 
					BD_ClinicaAbrao.pacientes 
				Where 
					id = ? 
				Order by nome`

	err := s.db.GetContext(ctx, paciente, query, id)
	if err != nil && err != sql.ErrNoRows {
		log.WithContext(ctx).Error("store.paciente.get_paciente_by_id ", err.Error())
		return nil, err
	}

	return paciente, nil
}

func (s *storeImpl) GetByAnything(ctx context.Context, paciente *model.Paciente) (*[]model.Paciente, error) {
	pacientes := new([]model.Paciente)
	query := `
			Select id, cpf, nome, telefone, telefone2, convenio, plano, acomodacao, status
				From 
					BD_ClinicaAbrao.pacientes 
				Where 
					1 = 1 `

	if len(paciente.Nome) > 0 {
		query += `and nome like '%` + paciente.Nome + `%' `
	}
	if len(paciente.Convenio) > 0 {
		query += `and convenio like '%` + paciente.Convenio + `%' `
	}
	if len(paciente.Cpf) > 0 {
		query += `and cpf like '%` + paciente.Cpf + `%' `
	}
	if len(paciente.Plano) > 0 {
		query += `and plano like '%` + paciente.Plano + `%' `
	}
	if paciente.Status > 0 {
		query += fmt.Sprintf(`and status = %d `, paciente.Status)
	}

	err := s.db.SelectContext(ctx, pacientes, query)
	if err != nil && err != sql.ErrNoRows {
		log.WithContext(ctx).Error("store.paciente.get_paciente_by_anything ", err.Error())
		return nil, err
	}

	return pacientes, nil
}

func (s *storeImpl) Set(ctx context.Context, paciente *model.Paciente) (*model.Paciente, error) {

	_, err := s.db.ExecContext(ctx, `INSERT INTO BD_ClinicaAbrao.pacientes (id, cpf, nome, telefone, telefone2, convenio, plano, acomodacao, status) VALUES (?,?,?,?,?,?,?,?,?)`,
		paciente.Id, paciente.Cpf, paciente.Nome, paciente.Telefone, paciente.Telefone2, paciente.Convenio, paciente.Plano, paciente.Acomodacao, paciente.Status)
	if err != nil {
		log.WithContext(ctx).Error("store.paciente.set_paciente", err.Error())
		return nil, err
	}

	return paciente, nil
}

func (s *storeImpl) Update(ctx context.Context, paciente *model.Paciente) (*model.Paciente, error) {
	_, err := s.db.ExecContext(ctx, `Update BD_ClinicaAbrao.pacientes SET cpf=?, nome=?, telefone=?, telefone2=?, convenio=?, plano=?, acomodacao=?, status=? Where id=?`,
		paciente.Cpf, paciente.Nome, paciente.Telefone, paciente.Telefone2, paciente.Convenio, paciente.Plano, paciente.Acomodacao, paciente.Status, paciente.Id)
	if err != nil {
		log.WithContext(ctx).Error("store.paciente.update", err.Error())
		return nil, err
	}

	return paciente, nil
}

func (s *storeImpl) Delete(ctx context.Context, id string) error {
	_, err := s.db.ExecContext(ctx, `Delete FROm BD_ClinicaAbrao.pacientes Where id=?`, id)
	if err != nil {
		log.WithContext(ctx).Error("store.paciente.delete", err.Error())
		return err
	}

	return nil
}
