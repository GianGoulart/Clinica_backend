package health

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/GianGoulart/Clinica_backend/model"
	"github.com/jmoiron/sqlx"
)

// Store interface para implementação do health
type Store interface {
	Ping(ctx context.Context) (*model.Health, error)
	Check(ctx context.Context) (*model.Health, error)
}

// NewStore cria uma nova instancia do repositorio de health
func NewStore(reader *sqlx.DB) Store {
	return &storeImpl{reader}
}

type storeImpl struct {
	reader *sqlx.DB
}

// Ping checa se o banco está online
func (r *storeImpl) Ping(ctx context.Context) (*model.Health, error) {
	err := r.reader.PingContext(ctx)
	if err != nil {
		logrus.Error(ctx, "store.health.ping", err.Error())
		return nil, err
	}

	return &model.Health{DatabaseStatus: "OK"}, nil
}

// Check checa se o banco está com status OK
func (r *storeImpl) Check(ctx context.Context) (*model.Health, error) {
	data := new(model.Health)

	err := r.reader.GetContext(ctx, data, `SELECT 'DB OK' AS database_status`)
	if err != nil {
		logrus.Error(ctx, "store.health.check", err.Error())
		return nil, err
	}

	return data, nil
}
