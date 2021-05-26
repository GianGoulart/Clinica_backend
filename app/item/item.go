package item

import (
	"context"
	"net/http"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/tradersclub/PocArquitetura/model"
	"github.com/tradersclub/PocArquitetura/store"
	nmodel "github.com/tradersclub/TCNatsModel/TCStartKit"
	"github.com/tradersclub/TCUtils/cache"
	"github.com/tradersclub/TCUtils/logger"
	"github.com/tradersclub/TCUtils/tcerr"
)

// App interface de health para implementação
type App interface {
	RequestItemById(ctx context.Context, id string) (*model.Item, error)
	GetItemById(ctx context.Context, id string) (*model.Item, error)
}

// NewApp cria uma nova instancia do serviço de item
func NewApp(stores *store.Container, nc *nats.Conn) App {
	return &appImpl{
		stores: stores,
		nc:     nc,
	}
}

type appImpl struct {
	stores    *store.Container
	startedAt time.Time
	version   string
	nc        *nats.Conn
	cache     cache.Cache
}

//GetItemById função de exemplo de request no nats
func (s *appImpl) RequestItemById(ctx context.Context, id string) (*model.Item, error) {

	obj := new(nmodel.GetItemById)
	obj.Id = id
	msg, err := s.nc.Request(nmodel.TCSTARTKIT_GET_ITEM_BY_ID, obj.ToBytes(), time.Second*100)
	if err != nil {
		return nil, tcerr.NewError(http.StatusInternalServerError, "erro ao resgatar o item", nil)
	}

	obj, err = nmodel.GetItemByIdFromBytes(msg.Data)
	if err != nil {
		return nil, tcerr.NewError(http.StatusInternalServerError, "erro ao realizar o parser do item", nil)
	}

	return obj.Data, nil
}

func (s *appImpl) GetItemById(ctx context.Context, id string) (*model.Item, error) {
	item := new(model.Item)

	// exemplo de consulta em cache
	if err := s.cache.Get(ctx, id, item); err != nil {
		logger.ErrorContext(ctx, "app.session.read_by_id", "não encontrei a sessão com id: "+id, err.Error())

		return nil, tcerr.New(http.StatusNotFound, "não encontrei a sessão", map[string]string{"id": id})
	}

	// exemplo de consulta em store
	result := <-s.stores.Item.GetItemById(ctx, id)
	if result.Error != nil {
		logger.ErrorContext(ctx, "app.item.get_item_by_id", result.Error.Error())
		return nil, result.Error
	}

	data, err := model.ToItem(result.Data)
	if err != nil {
		logger.ErrorContext(ctx, "app.health.check", err.Error())

		return nil, err
	}

	return data, nil
}
