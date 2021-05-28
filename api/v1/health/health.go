package health

import (
	"net/http"

	"github.com/GianGoulart/Clinica_backend/api/middleware"
	"github.com/GianGoulart/Clinica_backend/app"
	"github.com/GianGoulart/Clinica_backend/model"
	"github.com/labstack/echo/v4"
)

// Register group health check
func Register(g *echo.Group, apps *app.Container, m *middleware.Middleware) {
	h := &handler{
		apps: apps,
	}

	g.GET("", h.ping, m.Auth.Public)
	g.GET("/check", h.check, m.Auth.Public)
}

type handler struct {
	apps *app.Container
}

// ping swagger document
// @Description Essa rota é privada com o token valido (Bearer)
// @Tags health
// @Accept  json
// @Produce  json
// @Param item body model.Item true "add Item"
// @Success 200 {object} model.Item
// @Failure 400 {object} string
// @Router /v1/health [get]
func (h *handler) ping(c echo.Context) error {
	ctx := c.Request().Context()

	status, err := h.apps.Health.Ping(ctx)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.Response{
		Data: status,
	})
}

// check swagger document
// @Tags health
// @Accept  json
// @Produce  json
// @Param item body model.Item true "add Item"
// @Success 200 {object} model.Item
// @Failure 400 {object} string
// @Router /v1/health/check [get]
func (h *handler) check(c echo.Context) error {
	ctx := c.Request().Context()

	status, err := h.apps.Health.Check(ctx)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.Response{
		Data: status,
	})
}
