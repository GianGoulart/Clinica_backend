package item

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/GianGoulart/Clinica_backend/model"
	"github.com/jmoiron/sqlx"
)

// Store interface para implementação do item
type Store interface {
	GetItemById(ctx context.Context, id string) (*model.Item, error)
}

// NewStore cria uma nova instancia do repositorio de item
func NewStore(writer, reader *sqlx.DB) Store {
	return &storeImpl{writer, reader}
}

type storeImpl struct {
	writer *sqlx.DB
	reader *sqlx.DB
}

// GetItemById - pega o item no banco
func (r *storeImpl) GetItemById(ctx context.Context, id string) (*model.Item, error) {
	data := new(model.Item)

	err := r.reader.GetContext(ctx, data, `SELECT 'DB OK' AS database_status`)
	if err != nil {
		// Muito importante não retornar o erro do banco, apenas logar e retornar erro personalizado
		logrus.Error(ctx, "store.item.check", err.Error())
		return nil, err
	}

	return data, nil
}
