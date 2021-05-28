package middleware

import (
	"errors"

	"github.com/GianGoulart/Clinica_backend/model"
	"github.com/labstack/echo/v4"
)

// AuthMiddleware é a interface para geração dos middlewares
type AuthMiddleware interface {
	Public(next echo.HandlerFunc) echo.HandlerFunc
	Private(next echo.HandlerFunc) echo.HandlerFunc
}

// newAuthMiddleware cria uma implementação da interface AuthMiddleware
func newAuthMiddleware(opts Options) AuthMiddleware {
	return &middlewareAuthImpl{}
}

type middlewareAuthImpl struct{}

func (m *middlewareAuthImpl) Public(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(c)
	}
}

func (m *middlewareAuthImpl) Private(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := m.validateSession(c, true); err != nil {
			return err
		}

		return next(c)
	}
}

func (m *middlewareAuthImpl) validateSession(c echo.Context, logged bool, roles ...string) error {
	if logged {
		authsession, ok := c.Get("session").(*model.Session)
		if !ok || authsession == nil {
			return errors.New("você precisa de autenticação para realizar esta operação")
		}
	}

	return nil
}
