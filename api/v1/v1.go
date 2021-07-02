package v1

import (
	"github.com/GianGoulart/Clinica_backend/api/middleware"
	"github.com/GianGoulart/Clinica_backend/api/v1/acompanhamento"
	"github.com/GianGoulart/Clinica_backend/api/v1/comercial"
	"github.com/GianGoulart/Clinica_backend/api/v1/dashboard"
	"github.com/GianGoulart/Clinica_backend/api/v1/health"
	"github.com/GianGoulart/Clinica_backend/api/v1/item"
	"github.com/GianGoulart/Clinica_backend/api/v1/medicos"
	"github.com/GianGoulart/Clinica_backend/api/v1/pacientes"
	"github.com/GianGoulart/Clinica_backend/api/v1/procedimentos"
	"github.com/GianGoulart/Clinica_backend/api/v1/user"
	"github.com/GianGoulart/Clinica_backend/app"
	"github.com/labstack/echo/v4"
)

// Register regristra as rotas v1
func Register(g *echo.Group, apps *app.Container, middleware *middleware.Middleware) {
	v1 := g.Group("/v1", middleware.Session.InjectSession)

	health.Register(v1.Group("/health"), apps, middleware)
	item.Register(v1.Group("/item"), apps, middleware)
	pacientes.Register(v1.Group("/pacientes"), apps, middleware)
	medicos.Register(v1.Group("/medicos"), apps, middleware)
	procedimentos.Register(v1.Group("/procedimentos"), apps, middleware)
	comercial.Register(v1.Group("/comercial"), apps, middleware)
	acompanhamento.Register(v1.Group("/acompanhamentos"), apps, middleware)
	dashboard.Register(v1.Group("/dashboard"), apps, middleware)
	user.Register(v1.Group("/user"), apps, middleware)
}
