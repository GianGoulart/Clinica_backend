package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/tradersclub/PocArquitetura/api/middleware"
	"github.com/tradersclub/PocArquitetura/api/v1/health"
	"github.com/tradersclub/PocArquitetura/api/v1/item"
	"github.com/tradersclub/PocArquitetura/app"
)

// Register regristra as rotas v1
func Register(g *echo.Group, apps *app.Container, middleware *middleware.Middleware) {
	v1 := g.Group("/v1", middleware.Session.InjectSession)

	health.Register(v1.Group("/health"), apps, middleware)
	item.Register(v1.Group("/item"), apps, middleware)
}
