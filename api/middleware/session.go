package middleware

import (
	"net/http"
	"strings"

	"github.com/tradersclub/PocArquitetura/app"
	"github.com/tradersclub/PocArquitetura/model"
	"github.com/tradersclub/TCUtils/tcerr"

	"github.com/labstack/echo/v4"
)

// SessionMiddleware é a interface para geração dos middlewares
type SessionMiddleware interface {
	InjectSession(next echo.HandlerFunc) echo.HandlerFunc
}

// newSessionMiddleware cria uma implementação da interface SessionMiddleware
func newSessionMiddleware(opts Options) SessionMiddleware {
	return &middlewareSessionImpl{
		apps: opts.Apps,
	}
}

type middlewareSessionImpl struct {
	apps *app.Container
}

func (m *middlewareSessionImpl) InjectSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authorization := c.Request().Header.Get("Authorization")

		if authorization != "" {
			splitedToken := strings.Split(authorization, " ")
			if len(splitedToken) != 2 {
				return tcerr.New(http.StatusUnauthorized, "não foi possível decodificar o token", map[string]string{
					"authorization": authorization,
				})
			}

			session, err := m.apps.Session.ReadByID(c.Request().Context(), splitedToken[1])
			if err != nil {
				return err
			}

			c.Set("session", session)
			c.SetRequest(c.Request().WithContext(model.SetSession(c.Request().Context(), session)))
		}

		return next(c)
	}
}
