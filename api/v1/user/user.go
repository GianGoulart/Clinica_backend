package user

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

	g.POST("/getUser", h.getUser, m.Auth.Public)
	g.POST("", h.setUser, m.Auth.Public)
}

type handler struct {
	apps *app.Container
}

func (h handler) getUser(c echo.Context) error {
	ctx := c.Request().Context()

	payload := new(model.User)

	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Data: nil,
			Err:  err,
		})
	}

	response, err := h.apps.User.GetUser(ctx, payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Data: nil,
			Err:  err,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		Data: response,
	})

}

func (h handler) setUser(c echo.Context) error {
	ctx := c.Request().Context()

	payload := new(model.User)

	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Data: nil,
			Err:  err,
		})
	}

	response, err := h.apps.User.Set(ctx, payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Data: nil,
			Err:  err,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		Data: response,
	})

}
