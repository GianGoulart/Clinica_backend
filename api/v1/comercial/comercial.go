package comercial

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

	g.GET("", h.getAllComercial, m.Auth.Public)
	g.GET("/:comercial_id", h.getComercialById, m.Auth.Public)
	g.GET("/byProcedimento/:procedimento_id", h.getComercialByIdProcedimento, m.Auth.Public)
	g.POST("/anything", h.getComercialByAnything, m.Auth.Public)
	g.POST("", h.setComercial, m.Auth.Public)
	g.PUT("", h.updateComercial, m.Auth.Public)
	g.DELETE("/:comercial_id", h.deleteComercial, m.Auth.Public)
}

type handler struct {
	apps *app.Container
}

func (h handler) getAllComercial(c echo.Context) error {
	ctx := c.Request().Context()

	response, err := h.apps.Comercial.GetAll(ctx)
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

func (h handler) getComercialById(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("comercial_id")

	response, err := h.apps.Comercial.GetById(ctx, id)
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

func (h handler) getComercialByIdProcedimento(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("procedimento_id")

	response, err := h.apps.Comercial.GetByIdProcedimento(ctx, id)
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

func (h handler) getComercialByAnything(c echo.Context) error {
	ctx := c.Request().Context()

	payload := new(model.Comercial)

	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Data: nil,
			Err:  err,
		})
	}

	response, err := h.apps.Comercial.GetByAnything(ctx, payload)
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

func (h handler) setComercial(c echo.Context) error {
	ctx := c.Request().Context()

	payload := new(model.Comercial)

	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Data: nil,
			Err:  err,
		})
	}

	response, err := h.apps.Comercial.Set(ctx, payload)
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

func (h handler) updateComercial(c echo.Context) error {
	ctx := c.Request().Context()

	payload := new(model.Comercial)

	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Data: nil,
			Err:  err,
		})
	}

	response, err := h.apps.Comercial.Update(ctx, payload)
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

func (h handler) deleteComercial(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("comercial_id")

	err := h.apps.Comercial.Delete(ctx, id)
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
