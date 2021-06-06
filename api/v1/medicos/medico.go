package medicos

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

	g.GET("", h.getAllMedico, m.Auth.Public)
	g.GET("/:medico_id", h.getMedicoById, m.Auth.Public)
	g.POST("/anything", h.getMedicoByAnything, m.Auth.Public)
	g.POST("", h.setMedico, m.Auth.Public)
	g.PUT("", h.updateMedico, m.Auth.Public)
	g.DELETE("/:medico_id", h.deleteMedico, m.Auth.Public)
}

type handler struct {
	apps *app.Container
}

func (h handler) getAllMedico(c echo.Context) error {
	ctx := c.Request().Context()

	response, err := h.apps.Medico.GetAll(ctx)
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

func (h handler) getMedicoById(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("medico_id")

	response, err := h.apps.Medico.GetById(ctx, id)
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

func (h handler) getMedicoByAnything(c echo.Context) error {
	ctx := c.Request().Context()

	payload := new(model.Medico)

	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Data: nil,
			Err:  err,
		})
	}

	response, err := h.apps.Medico.GetByAnything(ctx, payload)
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

func (h handler) setMedico(c echo.Context) error {
	ctx := c.Request().Context()

	payload := new(model.Medico)

	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Data: nil,
			Err:  err,
		})
	}

	response, err := h.apps.Medico.Set(ctx, payload)
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

func (h handler) updateMedico(c echo.Context) error {
	ctx := c.Request().Context()

	payload := new(model.Medico)

	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Data: nil,
			Err:  err,
		})
	}

	response, err := h.apps.Medico.Update(ctx, payload)
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

func (h handler) deleteMedico(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("medico_id")

	err := h.apps.Medico.Delete(ctx, id)
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
