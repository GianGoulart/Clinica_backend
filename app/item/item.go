package item

import (
	"context"
	"time"

	"github.com/GianGoulart/Clinica_backend/model"
	"github.com/GianGoulart/Clinica_backend/store"
	"github.com/sirupsen/logrus"
)

// App interface de health para implementação
type App interface {
	GetItemById(ctx context.Context, id string) (*model.Item, error)
}

// NewApp cria uma nova instancia do serviço de item
func NewApp(stores *store.Container) App {
	return &appImpl{
		stores: stores,
	}
}

type appImpl struct {
	stores    *store.Container
	startedAt time.Time
	version   string
}

func (s *appImpl) GetItemById(ctx context.Context, id string) (*model.Item, error) {

	// exemplo de consulta em store
	result, err := s.stores.Item.GetItemById(ctx, id)
	if err != nil {
		logrus.Error(ctx, "app.item.get_item_by_id", err.Error())
		return nil, err
	}

	data, err := model.ToItem(result)
	if err != nil {
		logrus.Error(ctx, "app.health.check", err.Error())

		return nil, err
	}

	return data, nil
}
