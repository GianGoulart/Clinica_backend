package procedimentos

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

	g.GET("", h.getAllProcedimentos, m.Auth.Public)
	g.GET("/:procedimento_id", h.getProcedimentoById, m.Auth.Public)
	g.POST("/anything", h.getProcedimentoByAnything, m.Auth.Public)
	g.POST("", h.setProcedimento, m.Auth.Public)
	g.PUT("", h.updateProcedimento, m.Auth.Public)
	g.DELETE("/:procedimento_id", h.deleteProcedimento, m.Auth.Public)
}

type handler struct {
	apps *app.Container
}

func (h handler) getAllProcedimentos(c echo.Context) error {
	ctx := c.Request().Context()

	response, err := h.apps.Procedimento.GetAll(ctx)
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

func (h handler) getProcedimentoById(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("procedimento_id")

	response, err := h.apps.Procedimento.GetById(ctx, id)
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

func (h handler) getProcedimentoByAnything(c echo.Context) error {
	ctx := c.Request().Context()

	payload := new(model.Procedimento)

	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Data: nil,
			Err:  err,
		})
	}

	response, err := h.apps.Procedimento.GetByAnything(ctx, payload)
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

func (h handler) setProcedimento(c echo.Context) error {
	ctx := c.Request().Context()

	payload := new(model.Procedimento)

	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Data: nil,
			Err:  err,
		})
	}

	response, err := h.apps.Procedimento.Set(ctx, payload)
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

func (h handler) updateProcedimento(c echo.Context) error {
	ctx := c.Request().Context()

	payload := new(model.Procedimento)

	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Data: nil,
			Err:  err,
		})
	}

	response, err := h.apps.Procedimento.Update(ctx, payload)
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

func (h handler) deleteProcedimento(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("procedimento_id")

	err := h.apps.Procedimento.Delete(ctx, id)
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
