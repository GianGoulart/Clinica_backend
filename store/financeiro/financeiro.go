package financeiro

import (
	"context"
	"database/sql"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/GianGoulart/Clinica_backend/model"
	"github.com/jmoiron/sqlx"
)

type Store interface {
	GetAll(ctx context.Context) (*[]model.Financeiro, error)
	GetById(ctx context.Context, id string) (*model.Financeiro, error)
	GetByAnything(ctx context.Context, financeiro *model.Financeiro) (*[]model.Financeiro, error)
	Set(ctx context.Context, financeiro *model.Financeiro) (*model.Financeiro, error)
	Update(ctx context.Context, financeiro *model.Financeiro) (*model.Financeiro, error)
	Delete(ctx context.Context, id string) error
}

func NewStore(db *sqlx.DB) Store {
	return &storeImpl{db}
}

type storeImpl struct {
	db *sqlx.DB
}

func (s *storeImpl) GetAll(ctx context.Context) (*[]model.Financeiro, error) {
	financeiro := new([]model.Financeiro)
	query := `
				SELECT 
					f.*, pa.nome nome_paciente, m.nome nome_medico
				FROM 
					BD_ClinicaAbrao.financeiro f
				Inner Join BD_ClinicaAbrao.comercial c
					ON( f.id_comercial = c.id)
				Inner Join BD_ClinicaAbrao.procedimentos pr
					ON( pr.id = c.id_procedimento)
				Inner Join BD_ClinicaAbrao.pacientes pa
					ON( pr.id_paciente = pa.id)
				Inner Join BD_ClinicaAbrao.medicos m
					On(pr.id_medico = m.id)`

	err := s.db.SelectContext(ctx, financeiro, query)
	if err != nil && err != sql.ErrNoRows {
		log.WithContext(ctx).Error("store.financeiro.get_financeiro_all ", err.Error())
		return nil, err
	}

	return financeiro, nil
}

func (s *storeImpl) GetById(ctx context.Context, id string) (*model.Financeiro, error) {
	financeiro := new(model.Financeiro)
	query := `
			Select  *
				From 
					BD_ClinicaAbrao.financeiro 
				Where 
					id = ? `

	err := s.db.GetContext(ctx, financeiro, query, id)
	if err != nil && err != sql.ErrNoRows {
		log.WithContext(ctx).Error("store.financeiro.get_financeiro_by_id ", err.Error())
		return nil, err
	}

	return financeiro, nil
}

func (s *storeImpl) GetByAnything(ctx context.Context, financeiro *model.Financeiro) (*[]model.Financeiro, error) {
	listafinanceiro := new([]model.Financeiro)
	query := `
			SELECT 
					f.*, pa.nome nome_paciente, m.nome nome_medico
				FROM 
					BD_ClinicaAbrao.financeiro f
				Inner Join BD_ClinicaAbrao.comercial c
					ON( f.id_comercial = c.id)
				Inner Join BD_ClinicaAbrao.procedimentos pr
					ON( pr.id = c.id_procedimento)
				Inner Join BD_ClinicaAbrao.pacientes pa
					ON( pr.id_paciente = pa.id)
				Inner Join BD_ClinicaAbrao.medicos m
					On(pr.id_medico = m.id) 
				Where 1 = 1 `

	if len(financeiro.Id_Comercial) > 0 {
		query += ` and f.id_comercial = '` + financeiro.Id_Comercial + `'`
	}
	if financeiro.Data_Pagamento > 0 {
		query += ` and f.data_pagamento <= ` + strconv.Itoa(int(financeiro.Data_Pagamento))
	}
	if financeiro.Data_Compensacao > 0 {
		query += ` and f.data_compensacao <=` + strconv.Itoa(int(financeiro.Data_Compensacao))
	}
	if financeiro.Plano_Contas > 0 {
		query += ` and f.plano_contas =` + strconv.Itoa(int(financeiro.Plano_Contas))
	}
	if financeiro.Conta > 0 {
		query += ` and f.conta =` + strconv.Itoa(int(financeiro.Conta))
	}

	err := s.db.SelectContext(ctx, listafinanceiro, query)
	if err != nil && err != sql.ErrNoRows {
		log.WithContext(ctx).Error("store.financeiro.get_financeiro_by_anything ", err.Error())
		return nil, err
	}

	return listafinanceiro, nil
}

func (s *storeImpl) Set(ctx context.Context, financeiro *model.Financeiro) (*model.Financeiro, error) {

	_, err := s.db.ExecContext(ctx, `INSERT INTO BD_ClinicaAbrao.financeiro (id, id_comercial, data_pagamento, data_compensacao, plano_contas, conta, valor_ajuste, valor_liquido, obs) VALUES (?,?,?,?,?,?,?,?,?)`,
		financeiro.Id, financeiro.Id_Comercial, financeiro.Data_Pagamento, financeiro.Data_Compensacao, financeiro.Plano_Contas, financeiro.Conta, financeiro.Valor_Ajuste, financeiro.Valor_Liquido, financeiro.Obs)
	if err != nil {
		log.WithContext(ctx).Error("store.financeiro.set_financeiro", err.Error())
		return nil, err
	}

	return financeiro, nil
}

func (s *storeImpl) Update(ctx context.Context, financeiro *model.Financeiro) (*model.Financeiro, error) {
	_, err := s.db.ExecContext(ctx, `Update BD_ClinicaAbrao.financeiro SET id_comercial=?, data_pagamento=?, data_compensacao=?, plano_contas=?, conta=?, valor_ajuste=?, valor_liquido=?, obs=? Where id = ?`,
		financeiro.Id_Comercial, financeiro.Data_Pagamento, financeiro.Data_Compensacao, financeiro.Plano_Contas, financeiro.Conta, financeiro.Valor_Ajuste, financeiro.Valor_Liquido, financeiro.Obs, financeiro.Id)
	if err != nil {
		log.WithContext(ctx).Error("store.financeiro.update", err.Error())
		return nil, err
	}

	return financeiro, nil
}

func (s *storeImpl) Delete(ctx context.Context, id string) error {
	_, err := s.db.ExecContext(ctx, `Delete FROm BD_ClinicaAbrao.financeiro Where id=?`, id)
	if err != nil {
		log.WithContext(ctx).Error("store.financeiro.delete", err.Error())
		return err
	}

	return nil
}
