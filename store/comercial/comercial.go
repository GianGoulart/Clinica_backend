package comercial

import (
	"context"
	"database/sql"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/GianGoulart/Clinica_backend/model"
	"github.com/jmoiron/sqlx"
)

type Store interface {
	GetAll(ctx context.Context) (*[]model.Comercial, error)
	GetById(ctx context.Context, id string) (*model.Comercial, error)
	GetByAnything(ctx context.Context, comercial *model.Comercial) (*[]model.Comercial, error)
	Set(ctx context.Context, comercial *model.Comercial) (*model.Comercial, error)
	Update(ctx context.Context, comercial *model.Comercial) (*model.Comercial, error)
	Delete(ctx context.Context, id string) error
	GetByIdProcedimento(ctx context.Context, id string) (*[]model.Comercial, error)
}

func NewStore(db *sqlx.DB) Store {
	return &storeImpl{db}
}

type storeImpl struct {
	db *sqlx.DB
}

func (s *storeImpl) GetAll(ctx context.Context) (*[]model.Comercial, error) {
	comercial := new([]model.Comercial)
	query := `
				SELECT 
					c.*, ifnull(pa.nome ,"") nome_paciente,	ifnull(m.nome ,"")  nome_medico 
				FROM 
					BD_ClinicaAbrao.comercial c
				Left Join BD_ClinicaAbrao.procedimentos pr
				ON( pr.id = c.id_procedimento)
				Left Join BD_ClinicaAbrao.pacientes pa
				ON( pr.id_paciente = pa.id)
				Left Join BD_ClinicaAbrao.medicos m
				On(pr.id_medico = m.id)`

	err := s.db.SelectContext(ctx, comercial, query)
	if err != nil && err != sql.ErrNoRows {
		log.WithContext(ctx).Error("store.comercial.get_comercial_all ", err.Error())
		return nil, err
	}

	return comercial, nil
}

func (s *storeImpl) GetById(ctx context.Context, id string) (*model.Comercial, error) {
	comercial := new(model.Comercial)
	query := `
			Select  *
				From 
					BD_ClinicaAbrao.comercial 
				Where 
					id = ? `

	err := s.db.GetContext(ctx, comercial, query, id)
	if err != nil && err != sql.ErrNoRows {
		log.WithContext(ctx).Error("store.comercial.get_comercial_by_id ", err.Error())
		return nil, err
	}

	return comercial, nil
}

func (s *storeImpl) GetByIdProcedimento(ctx context.Context, id string) (*[]model.Comercial, error) {
	comercial := new([]model.Comercial)
	query := `SELECT 
					c.*, ifnull(pa.nome ,"") nome_paciente,	ifnull(m.nome ,"")  nome_medico
				FROM 
					BD_ClinicaAbrao.comercial c
				Left Join BD_ClinicaAbrao.procedimentos pr
				ON( pr.id = c.id_procedimento)
				Left Join BD_ClinicaAbrao.pacientes pa
				ON( pr.id_paciente = pa.id)
				Left Join BD_ClinicaAbrao.medicos m
				On(pr.id_medico = m.id)
				Where c.id_procedimento = ? `

	err := s.db.SelectContext(ctx, comercial, query, id)
	if err != nil && err != sql.ErrNoRows {
		log.WithContext(ctx).Error("store.comercial.get_comercial_by_id ", err.Error())
		return nil, err
	}

	return comercial, nil
}

func (s *storeImpl) GetByAnything(ctx context.Context, comercial *model.Comercial) (*[]model.Comercial, error) {
	listacomercial := new([]model.Comercial)
	query := `
			SELECT 
				c.*, ifnull(pa.nome ,"") nome_paciente,	ifnull(m.nome ,"")  nome_medico 
			FROM 
					BD_ClinicaAbrao.comercial c
				Left Join BD_ClinicaAbrao.procedimentos pr
					ON( pr.id = c.id_procedimento)
				Left Join BD_ClinicaAbrao.pacientes pa
					ON( pr.id_paciente = pa.id)
				Left Join BD_ClinicaAbrao.medicos m
					On(pr.id_medico = m.id) 
				Where 1 = 1 `

	if len(comercial.Id_Procedimento) > 0 {
		query += ` and c.id_procedimento = '` + comercial.Id_Procedimento + `'`
	}
	if comercial.Tipo_Pagamento > 0 {
		query += ` and c.tipo_pagamento = ` + strconv.Itoa(int(comercial.Tipo_Pagamento))
	}
	if comercial.Data_Vencimento > 0 {
		query += ` and c.data_vencimento <=` + strconv.Itoa(int(comercial.Data_Vencimento))
	}
	if comercial.Data_Emissao_NF > 0 {
		query += ` and c.data_emissao_nf <=` + strconv.Itoa(int(comercial.Data_Emissao_NF))
	}

	err := s.db.SelectContext(ctx, listacomercial, query)
	if err != nil && err != sql.ErrNoRows {
		log.WithContext(ctx).Error("store.comercial.get_comercial_by_anything ", err.Error())
		return nil, err
	}

	return listacomercial, nil
}

func (s *storeImpl) Set(ctx context.Context, comercial *model.Comercial) (*model.Comercial, error) {

	_, err := s.db.ExecContext(ctx, `INSERT INTO BD_ClinicaAbrao.comercial (id, id_procedimento, id_medico_part, funcao_medico_part, qtd_parcelas, valor_parcelas, tipo_pagamento, forma_pagamento, data_emissao_nf, data_vencimento, data_pagamento, data_compensacao, plano_contas, conta, valor_ajuste, valor_liquido, obs) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
		comercial.Id, comercial.Id_Procedimento, comercial.Id_Medico_Part, comercial.Funcao_Medico_Part, comercial.Qtd_Parcelas, comercial.Valor_Parcelas, comercial.Tipo_Pagamento, comercial.Forma_Pagamento, comercial.Data_Emissao_NF, comercial.Data_Vencimento, comercial.Data_Pagamento, comercial.Data_Compensacao, comercial.Plano_Contas, comercial.Conta, comercial.Valor_Ajuste, comercial.Valor_Liquido, comercial.Obs)
	if err != nil {
		log.WithContext(ctx).Error("store.comercial.set_comercial", err.Error())
		return nil, err
	}

	return comercial, nil
}

func (s *storeImpl) Update(ctx context.Context, comercial *model.Comercial) (*model.Comercial, error) {
	_, err := s.db.ExecContext(ctx, `Update BD_ClinicaAbrao.comercial SET id_procedimento=? , id_medico_part=? , funcao_medico_part=? , qtd_parcelas=? , valor_parcelas=? , tipo_pagamento=? , forma_pagamento=? , data_emissao_nf=? , data_vencimento=?, data_pagamento=?, data_compensacao=?, plano_contas=?, conta=?, valor_ajuste=?, valor_liquido=?, obs=? Where id = ?`,
		comercial.Id_Procedimento, comercial.Id_Medico_Part, comercial.Funcao_Medico_Part, comercial.Qtd_Parcelas, comercial.Valor_Parcelas, comercial.Tipo_Pagamento, comercial.Forma_Pagamento, comercial.Data_Emissao_NF, comercial.Data_Vencimento, comercial.Data_Pagamento, comercial.Data_Compensacao, comercial.Plano_Contas, comercial.Conta, comercial.Valor_Ajuste, comercial.Valor_Liquido, comercial.Obs, comercial.Id)
	if err != nil {
		log.WithContext(ctx).Error("store.Comercial.update", err.Error())
		return nil, err
	}

	return comercial, nil
}

func (s *storeImpl) Delete(ctx context.Context, id string) error {
	_, err := s.db.ExecContext(ctx, `Delete FROm BD_ClinicaAbrao.comercial Where id=?`, id)
	if err != nil {
		log.WithContext(ctx).Error("store.Comercial.delete", err.Error())
		return err
	}

	return nil
}
