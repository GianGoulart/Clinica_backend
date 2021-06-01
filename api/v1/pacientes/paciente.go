package pacientes

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

	g.GET("", h.getAllPaciente, m.Auth.Public)
	g.GET("/:paciente_id", h.getPacienteById, m.Auth.Public)
	g.POST("/anything", h.getPacienteByAnything, m.Auth.Public)
	g.POST("", h.setPaciente, m.Auth.Public)
	g.PUT("", h.updatePaciente, m.Auth.Public)
	g.DELETE("/:paciente_id", h.deletePaciente, m.Auth.Public)
}

type handler struct {
	apps *app.Container
}

func (h handler) getAllPaciente(c echo.Context) error {
	ctx := c.Request().Context()

	response, err := h.apps.Paciente.GetAll(ctx)
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

func (h handler) getPacienteById(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("paciente_id")

	response, err := h.apps.Paciente.GetById(ctx, id)
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

func (h handler) getPacienteByAnything(c echo.Context) error {
	ctx := c.Request().Context()

	payload := new(model.Paciente)

	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Data: nil,
			Err:  err,
		})
	}

	response, err := h.apps.Paciente.GetByAnything(ctx, payload)
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

func (h handler) setPaciente(c echo.Context) error {
	ctx := c.Request().Context()

	payload := new(model.Paciente)

	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Data: nil,
			Err:  err,
		})
	}

	response, err := h.apps.Paciente.Set(ctx, payload)
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

func (h handler) updatePaciente(c echo.Context) error {
	ctx := c.Request().Context()

	payload := new(model.Paciente)

	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Data: nil,
			Err:  err,
		})
	}

	response, err := h.apps.Paciente.Update(ctx, payload)
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

func (h handler) deletePaciente(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("paciente_id")

	err := h.apps.Paciente.Delete(ctx, id)
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
