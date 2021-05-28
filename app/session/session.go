package session

import (
	"context"
	"errors"

	"github.com/sirupsen/logrus"

	"github.com/GianGoulart/Clinica_backend/model"
)

// App interface de session para implementação
type App interface {
	ReadByID(ctx context.Context, id string) (*model.Session, error)
}

// NewApp cria uma nova instancia do serviço de session
func NewApp(cache model.Cache) App {
	return &appImpl{
		cache: cache,
	}
}

type appImpl struct {
	cache model.Cache
}

// ReadByID faz a leitura de uma sessão baseada no seu id
func (s *appImpl) ReadByID(ctx context.Context, id string) (*model.Session, error) {
	sess := new(model.Session)
	if err := s.cache.Get(ctx, id, sess); err != nil {
		logrus.Error(ctx, "app.session.read_by_id", "não encontrei a sessão com id: "+id, err.Error())

		return nil, errors.New("não encontrei a sessão")
	}

	return sess, nil
}
