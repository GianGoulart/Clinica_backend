package api

import (
	"github.com/labstack/echo/v4"
	"github.com/tradersclub/PocArquitetura/api/middleware"
	v1 "github.com/tradersclub/PocArquitetura/api/v1"
	"github.com/tradersclub/PocArquitetura/app"
	"github.com/tradersclub/TCUtils/logger"
)

// Options struct de opções para a criação de uma instancia das rotas
type Options struct {
	Group      *echo.Group
	Apps       *app.Container
	Middleware *middleware.Middleware
}

// Register api instance
func Register(opts Options) {
	v1.Register(opts.Group, opts.Apps, opts.Middleware)

	logger.Info("Registered -> Api")
}
