package acompanhamento

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/GianGoulart/Clinica_backend/model"
	"github.com/jmoiron/sqlx"
)

type Store interface {
	GetAll(ctx context.Context) (*[]model.Acompanhamento, error)
	GetById(ctx context.Context, id string) (*model.Acompanhamento, error)
	GetByAnything(ctx context.Context, acompanhamento *model.Acompanhamento) (*[]model.Acompanhamento, error)
	Set(ctx context.Context, acompanhamento *model.Acompanhamento) (*model.Acompanhamento, error)
	Update(ctx context.Context, acompanhamento *model.Acompanhamento) (*model.Acompanhamento, error)
	Delete(ctx context.Context, id string) error
}

func NewStore(db *sqlx.DB) Store {
	return &storeImpl{db}
}

type storeImpl struct {
	db *sqlx.DB
}

func (s *storeImpl) GetAll(ctx context.Context) (*[]model.Acompanhamento, error) {
	acompanhamento := new([]model.Acompanhamento)
	query := `
				SELECT 
					a.id, a.id_procedimento, envio_protocolo, solicitacao_previa, confirmacao_solicitacao, finalizacao_previa, status_previa, envio_convenio, liberacao, repasse_paciente, repasse_clinica, status_reembolso, obs 
				FROM 
					BD_ClinicaAbrao.acompanhamentos a 
				Left Join BD_ClinicaAbrao.procedimentos pr
				ON( pr.id = a.id_procedimento)`

	err := s.db.SelectContext(ctx, acompanhamento, query)
	if err != nil && err != sql.ErrNoRows {
		log.WithContext(ctx).Error("store.acompanhamento.get_acompanhamento_all ", err.Error())
		return nil, err
	}

	return acompanhamento, nil
}

func (s *storeImpl) GetById(ctx context.Context, id string) (*model.Acompanhamento, error) {
	acompanhamento := new(model.Acompanhamento)
	query := `
			Select  
			a.id, a.id_procedimento, envio_protocolo, solicitacao_previa, confirmacao_solicitacao, finalizacao_previa, status_previa, envio_convenio, liberacao, repasse_paciente, repasse_clinica, status_reembolso, obs
				From 
					BD_ClinicaAbrao.acompanhamentos a
				Where 
					id = ? `

	err := s.db.GetContext(ctx, acompanhamento, query, id)
	if err != nil && err != sql.ErrNoRows {
		log.WithContext(ctx).Error("store.acompanhamento.get_acompanhamento_by_id ", err.Error())
		return nil, err
	}

	return acompanhamento, nil
}

func (s *storeImpl) GetByAnything(ctx context.Context, acompanhamento *model.Acompanhamento) (*[]model.Acompanhamento, error) {
	listaacompanhamento := new([]model.Acompanhamento)
	query := `
			SELECT 
			a.id, a.id_procedimento, envio_protocolo, solicitacao_previa, confirmacao_solicitacao, finalizacao_previa, status_previa, envio_convenio, liberacao, repasse_paciente, repasse_clinica, status_reembolso, obs
				FROM 
					BD_ClinicaAbrao.acompanhamentos a
				Left Join BD_ClinicaAbrao.procedimentos pr
					ON( pr.id = a.id_procedimento) 
				Where 1 = 1 `

	if acompanhamento.Status_Previa > 0 {
		query += ` and a.status_previa = ` + strconv.Itoa(int(acompanhamento.Status_Previa))
	}
	if acompanhamento.Status_Reembolso > 0 {
		query += ` and a.status_reembolso =` + strconv.Itoa(int(acompanhamento.Status_Reembolso))
	}
	if len(acompanhamento.Id_Procedimento) > 0 {
		query += ` and a.id_procedimento ='` + acompanhamento.Id_Procedimento + `'`
	}

	err := s.db.SelectContext(ctx, listaacompanhamento, query)
	if err != nil && err != sql.ErrNoRows {
		log.WithContext(ctx).Error("store.Acompanhamento.get_acompanhamento_by_anything ", err.Error())
		return nil, err
	}

	return listaacompanhamento, nil
}

func (s *storeImpl) Set(ctx context.Context, acompanhamento *model.Acompanhamento) (*model.Acompanhamento, error) {
	fmt.Println(acompanhamento)
	_, err := s.db.ExecContext(ctx, `INSERT INTO BD_ClinicaAbrao.acompanhamentos (id, id_procedimento, envio_protocolo, solicitacao_previa, confirmacao_solicitacao, finalizacao_previa, status_previa, envio_convenio, liberacao, repasse_paciente, repasse_clinica, status_reembolso, obs) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)`,
		acompanhamento.Id, acompanhamento.Id_Procedimento, acompanhamento.Envio_Protocolo, acompanhamento.Solicitacao_Previa, acompanhamento.Confirmacao_Solicitacao, acompanhamento.Finalizacao_Previa, acompanhamento.Status_Previa, acompanhamento.Envio_Convenio, acompanhamento.Liberacao, acompanhamento.Repasse_Paciente, acompanhamento.Repasse_Clinica, acompanhamento.Status_Reembolso, acompanhamento.Obs)
	if err != nil {
		log.WithContext(ctx).Error("store.Acompanhamento.set_acompanhamento", err.Error())
		return nil, err
	}

	return acompanhamento, nil
}

func (s *storeImpl) Update(ctx context.Context, acompanhamento *model.Acompanhamento) (*model.Acompanhamento, error) {
	_, err := s.db.ExecContext(ctx, `Update BD_ClinicaAbrao.acompanhamentos SET id_procedimento=?, envio_protocolo=?, solicitacao_previa=?, confirmacao_solicitacao=?, finalizacao_previa=?, status_previa=?, envio_convenio=?, liberacao=?, repasse_paciente=?, repasse_clinica=?, status_reembolso=?, obs=? Where id = ?`,
		acompanhamento.Id_Procedimento, acompanhamento.Envio_Protocolo, acompanhamento.Solicitacao_Previa, acompanhamento.Confirmacao_Solicitacao, acompanhamento.Finalizacao_Previa, acompanhamento.Status_Previa, acompanhamento.Envio_Convenio, acompanhamento.Liberacao, acompanhamento.Repasse_Paciente, acompanhamento.Repasse_Clinica, acompanhamento.Status_Reembolso, acompanhamento.Obs, acompanhamento.Id)
	if err != nil {
		log.WithContext(ctx).Error("store.Acompanhamento.update", err.Error())
		return nil, err
	}

	return acompanhamento, nil
}

func (s *storeImpl) Delete(ctx context.Context, id string) error {
	_, err := s.db.ExecContext(ctx, `Delete FROm BD_ClinicaAbrao.acompanhamentos Where id=?`, id)
	if err != nil {
		log.WithContext(ctx).Error("store.Acompanhamento.delete", err.Error())
		return err
	}

	return nil
}
