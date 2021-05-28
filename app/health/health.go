package health

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/GianGoulart/Clinica_backend/model"
	"github.com/GianGoulart/Clinica_backend/store"
)

// App interface de health para implementação
type App interface {
	Ping(ctx context.Context) (*model.Health, error)
	Check(ctx context.Context) (*model.Health, error)
}

// NewApp cria uma nova instancia do serviço de health
func NewApp(stores *store.Container, version string, startedAt time.Time) App {
	return &appImpl{
		stores:    stores,
		version:   version,
		startedAt: startedAt,
	}
}

type appImpl struct {
	stores    *store.Container
	startedAt time.Time
	version   string
}

func (s *appImpl) Ping(ctx context.Context) (*model.Health, error) {
	result, err := s.stores.Health.Ping(ctx)
	if err != nil {
		logrus.Error(ctx, "app.health.ping", err.Error())

		return nil, err
	}

	data, err := model.ToHealth(result)
	if err != nil {
		logrus.Error(ctx, "app.health.ping", err.Error())

		return nil, err
	}
	data.ServerStatedAt = s.startedAt.UTC().String()
	data.Version = s.version

	return data, nil
}

func (s *appImpl) Check(ctx context.Context) (*model.Health, error) {
	result, err := s.stores.Health.Check(ctx)
	if err != nil {
		logrus.Error(ctx, "app.health.check", err.Error())

		return nil, err
	}

	data, err := model.ToHealth(result)
	if err != nil {
		logrus.Error(ctx, "app.health.check", err.Error())

		return nil, err
	}
	data.ServerStatedAt = s.startedAt.UTC().String()
	data.Version = s.version

	return data, nil
}
