package financeiro

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

	g.GET("", h.getAllFinanceiro, m.Auth.Public)
	g.GET("/:financeiro_id", h.getFinanceiroById, m.Auth.Public)
	g.POST("/anything", h.getFinanceiroByAnything, m.Auth.Public)
	g.POST("", h.setFinanceiro, m.Auth.Public)
	g.PUT("", h.updateFinanceiro, m.Auth.Public)
	g.DELETE("/:financeiro_id", h.deleteFinanceiro, m.Auth.Public)
}

type handler struct {
	apps *app.Container
}

func (h handler) getAllFinanceiro(c echo.Context) error {
	ctx := c.Request().Context()

	response, err := h.apps.Financeiro.GetAll(ctx)
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

func (h handler) getFinanceiroById(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("financeiro_id")

	response, err := h.apps.Financeiro.GetById(ctx, id)
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

func (h handler) getFinanceiroByAnything(c echo.Context) error {
	ctx := c.Request().Context()

	payload := new(model.Financeiro)

	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Data: nil,
			Err:  err,
		})
	}

	response, err := h.apps.Financeiro.GetByAnything(ctx, payload)
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

func (h handler) setFinanceiro(c echo.Context) error {
	ctx := c.Request().Context()

	payload := new(model.Financeiro)

	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Data: nil,
			Err:  err,
		})
	}

	response, err := h.apps.Financeiro.Set(ctx, payload)
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

func (h handler) updateFinanceiro(c echo.Context) error {
	ctx := c.Request().Context()

	payload := new(model.Financeiro)

	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Data: nil,
			Err:  err,
		})
	}

	response, err := h.apps.Financeiro.Update(ctx, payload)
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

func (h handler) deleteFinanceiro(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param(":financeiro_id")

	err := h.apps.Financeiro.Delete(ctx, id)
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
