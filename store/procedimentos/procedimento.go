package procedimentos

import (
	"context"
	"database/sql"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/GianGoulart/Clinica_backend/model"
	"github.com/jmoiron/sqlx"
)

type Store interface {
	GetAll(ctx context.Context) (*[]model.Procedimento, error)
	GetById(ctx context.Context, id string) (*model.Procedimento, error)
	GetByAnything(ctx context.Context, Procedimento *model.Procedimento) (*[]model.Procedimento, error)
	Set(ctx context.Context, procedimento *model.Procedimento) (*model.Procedimento, error)
	Update(ctx context.Context, procedimento *model.Procedimento) (*model.Procedimento, error)
	Delete(ctx context.Context, id string) error
}

func NewStore(db *sqlx.DB) Store {
	return &storeImpl{db}
}

type storeImpl struct {
	db *sqlx.DB
}

func (s *storeImpl) GetAll(ctx context.Context) (*[]model.Procedimento, error) {
	procedimentos := new([]model.Procedimento)
	query := `
				SELECT 
					pr.id, pr.id_paciente, pa.nome nome_paciente, pr.id_medico, m.nome nome_medico, pr.desc_procedimento, pr.procedimento, pr.local_procedimento, pr.status, pr.data, pr.valor, pr.esteira 
				FROM 
					BD_ClinicaAbrao.procedimentos pr
				Inner Join BD_ClinicaAbrao.pacientes pa
				ON( pr.id_paciente = pa.id)
				Inner Join BD_ClinicaAbrao.medicos m
				On(pr.id_medico = m.id)`

	err := s.db.SelectContext(ctx, procedimentos, query)
	if err != nil && err != sql.ErrNoRows {
		log.WithContext(ctx).Error("store.procedimentos.get_procedimento_all ", err.Error())
		return nil, err
	}

	return procedimentos, nil
}

func (s *storeImpl) GetById(ctx context.Context, id string) (*model.Procedimento, error) {
	procedimento := new(model.Procedimento)
	query := `
			Select  id, id_paciente, id_medico, desc_procedimento, procedimento, local_procedimento, status, data, valor, esteira
				From 
					BD_ClinicaAbrao.procedimentos 
				Where 
					id = ? `

	err := s.db.GetContext(ctx, procedimento, query, id)
	if err != nil && err != sql.ErrNoRows {
		log.WithContext(ctx).Error("store.procedimento.get_procedimento_by_id ", err.Error())
		return nil, err
	}

	return procedimento, nil
}

func (s *storeImpl) GetByAnything(ctx context.Context, procedimento *model.Procedimento) (*[]model.Procedimento, error) {
	procedimentos := new([]model.Procedimento)
	query := `
			Select  id, id_paciente, id_medico, desc_procedimento, procedimento, local_procedimento, status, data, valor, esteira
				From 
					BD_ClinicaAbrao.procedimentos 
				Where 
					1 = 1 `

	if procedimento.Procedimento > 0 {
		query += `and procedimento = ` + strconv.Itoa(int(procedimento.Procedimento))
	}
	if procedimento.Data > 0 {
		query += `and data <=` + strconv.Itoa(int(procedimento.Data))
	}
	if procedimento.Local_Procedimento > 0 {
		query += `and local_procedimento =` + strconv.Itoa(int(procedimento.Local_Procedimento))
	}
	if len(procedimento.Id_Medico) > 0 {
		query += `and id_medico = '` + procedimento.Id_Medico + `' `
	}
	if len(procedimento.Id_Paciente) > 0 {
		query += `and id_paciente = '` + procedimento.Id_Paciente + `' `
	}

	err := s.db.SelectContext(ctx, procedimentos, query)
	if err != nil && err != sql.ErrNoRows {
		log.WithContext(ctx).Error("store.procedimento.get_procedimento_by_anything ", err.Error())
		return nil, err
	}

	return procedimentos, nil
}

func (s *storeImpl) Set(ctx context.Context, procedimento *model.Procedimento) (*model.Procedimento, error) {

	_, err := s.db.ExecContext(ctx, `INSERT INTO BD_ClinicaAbrao.procedimentos (id, id_paciente, id_medico, desc_procedimento, procedimento, local_procedimento, status, data, valor, esteira) VALUES (?,?,?,?,?,?,?,?,?,?)`,
		procedimento.Id, procedimento.Id_Paciente, procedimento.Id_Medico, procedimento.Desc_Procedimento, procedimento.Procedimento, procedimento.Local_Procedimento, procedimento.Status, procedimento.Data, procedimento.Valor, procedimento.Esteira)
	if err != nil {
		log.WithContext(ctx).Error("store.procedimento.set_procedimento", err.Error())
		return nil, err
	}

	return procedimento, nil
}

func (s *storeImpl) Update(ctx context.Context, procedimento *model.Procedimento) (*model.Procedimento, error) {
	_, err := s.db.ExecContext(ctx, `Update BD_ClinicaAbrao.procedimentos SET Id_Paciente = ?, id_medico = ?, desc_procedimento = ?, procedimento = ?, local_procedimento = ?, status = ?, data = ?, valor = ?, esteira = ? Where id = ?`,
		procedimento.Id_Paciente, procedimento.Id_Medico, procedimento.Desc_Procedimento, procedimento.Procedimento, procedimento.Local_Procedimento, procedimento.Status, procedimento.Data, procedimento.Valor, procedimento.Esteira, procedimento.Id)
	if err != nil {
		log.WithContext(ctx).Error("store.procedimento.update", err.Error())
		return nil, err
	}

	return procedimento, nil
}

func (s *storeImpl) Delete(ctx context.Context, id string) error {
	_, err := s.db.ExecContext(ctx, `Delete FROm BD_ClinicaAbrao.procedimentos Where id=?`, id)
	if err != nil {
		log.WithContext(ctx).Error("store.procedimentos.delete", err.Error())
		return err
	}

	return nil
}
