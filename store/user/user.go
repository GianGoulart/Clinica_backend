package user

import (
	"context"
	"database/sql"

	log "github.com/sirupsen/logrus"

	"github.com/GianGoulart/Clinica_backend/model"
	"github.com/jmoiron/sqlx"
)

type Store interface {
	GetUser(ctx context.Context, user *model.User) (*model.User, error)
	Set(ctx context.Context, user *model.User) (*model.User, error)
	// Update(ctx context.Context, user *model.User) (*model.User, error)
	// Delete(ctx context.Context, id string) error
}

func NewStore(db *sqlx.DB) Store {
	return &storeImpl{db}
}

type storeImpl struct {
	db *sqlx.DB
}

func (s *storeImpl) GetUser(ctx context.Context, user *model.User) (*model.User, error) {
	res := new(model.User)
	query := `	SELECT
					*
				FROM 
					BD_ClinicaAbrao.users u
				Where 
					u.nome = ?`

	err := s.db.GetContext(ctx, res, query, user.Nome)
	if err != nil && err != sql.ErrNoRows {
		log.WithContext(ctx).Error("store.user.get_user ", err.Error())
		return nil, err
	}

	return res, nil
}

func (s *storeImpl) Set(ctx context.Context, user *model.User) (*model.User, error) {

	_, err := s.db.ExecContext(ctx, `INSERT INTO BD_ClinicaAbrao.users (id, nome, senha, roles, email) VALUES (?,?,?,?,?)`,
		user.Id, user.Nome, user.Senha, user.Roles, user.Email)
	if err != nil {
		log.WithContext(ctx).Error("store.user.set", err.Error())
		return nil, err
	}

	return user, nil
}
