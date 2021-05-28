package api

import (
	"github.com/GianGoulart/Clinica_backend/api/middleware"
	v1 "github.com/GianGoulart/Clinica_backend/api/v1"
	"github.com/GianGoulart/Clinica_backend/app"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
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

	log.Info("Registered -> Api")
}
