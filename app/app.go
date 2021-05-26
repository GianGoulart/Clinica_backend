package app

import (
	"time"

	"github.com/nats-io/nats.go"

	"github.com/tradersclub/PocArquitetura/app/health"
	"github.com/tradersclub/PocArquitetura/app/item"
	"github.com/tradersclub/PocArquitetura/app/session"
	"github.com/tradersclub/PocArquitetura/store"
	"github.com/tradersclub/TCUtils/cache"
	"github.com/tradersclub/TCUtils/logger"
)

// Container modelo para exportação dos serviços instanciados
type Container struct {
	Health  health.App
	Item    item.App
	Session session.App
}

// Options struct de opções para a criação de uma instancia dos serviços
type Options struct {
	Stores *store.Container
	Cache  cache.Cache
	Nats   *nats.Conn

	StartedAt time.Time
	Version   string
}

// New cria uma nova instancia dos serviços
func New(opts Options) *Container {

	container := &Container{
		Health:  health.NewApp(opts.Stores, opts.Version, opts.StartedAt),
		Item:    item.NewApp(opts.Stores, opts.Nats),
		Session: session.NewApp(opts.Cache),
	}

	logger.Info("Registered -> App")

	return container

}
