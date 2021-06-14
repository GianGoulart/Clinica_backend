package acompanhamento

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

	g.GET("", h.getAllAcompanhamento, m.Auth.Public)
	g.GET("/:acompanhamento_id", h.getAcompanhamentoById, m.Auth.Public)
	g.POST("/anything", h.getAcompanhamentoByAnything, m.Auth.Public)
	g.POST("", h.setAcompanhamento, m.Auth.Public)
	g.PUT("", h.updateAcompanhamento, m.Auth.Public)
	g.DELETE("/:acompanhamento_id", h.deleteAcompanhamento, m.Auth.Public)
}

type handler struct {
	apps *app.Container
}

func (h handler) getAllAcompanhamento(c echo.Context) error {
	ctx := c.Request().Context()

	response, err := h.apps.Acompanhamento.GetAll(ctx)
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

func (h handler) getAcompanhamentoById(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("acompanhamento_id")

	response, err := h.apps.Acompanhamento.GetById(ctx, id)
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

func (h handler) getAcompanhamentoByAnything(c echo.Context) error {
	ctx := c.Request().Context()

	payload := new(model.Acompanhamento)

	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Data: nil,
			Err:  err,
		})
	}

	response, err := h.apps.Acompanhamento.GetByAnything(ctx, payload)
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

func (h handler) setAcompanhamento(c echo.Context) error {
	ctx := c.Request().Context()

	payload := new(model.Acompanhamento)

	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Data: nil,
			Err:  err,
		})
	}

	response, err := h.apps.Acompanhamento.Set(ctx, payload)
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

func (h handler) updateAcompanhamento(c echo.Context) error {
	ctx := c.Request().Context()

	payload := new(model.Acompanhamento)

	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Data: nil,
			Err:  err,
		})
	}

	response, err := h.apps.Acompanhamento.Update(ctx, payload)
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

func (h handler) deleteAcompanhamento(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("acompanhamento_id")

	err := h.apps.Acompanhamento.Delete(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Data: nil,
			Err:  err,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		Data: "ok",
	})

}
