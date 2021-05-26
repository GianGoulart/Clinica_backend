package session

import (
	"context"
	"net/http"

	"github.com/tradersclub/TCUtils/logger"

	"github.com/tradersclub/TCUtils/cache"
	"github.com/tradersclub/TCUtils/tcerr"

	"github.com/tradersclub/PocArquitetura/model"
)

// App interface de session para implementação
type App interface {
	ReadByID(ctx context.Context, id string) (*model.Session, error)
}

// NewApp cria uma nova instancia do serviço de session
func NewApp(cache cache.Cache) App {
	return &appImpl{
		cache: cache,
	}
}

type appImpl struct {
	cache cache.Cache
}

// ReadByID faz a leitura de uma sessão baseada no seu id
func (s *appImpl) ReadByID(ctx context.Context, id string) (*model.Session, error) {
	sess := new(model.Session)
	if err := s.cache.Get(ctx, id, sess); err != nil {
		logger.ErrorContext(ctx, "app.session.read_by_id", "não encontrei a sessão com id: "+id, err.Error())

		return nil, tcerr.New(http.StatusNotFound, "não encontrei a sessão", map[string]string{"id": id})
	}

	return sess, nil
}
