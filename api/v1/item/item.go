package item

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tradersclub/PocArquitetura/api/middleware"
	"github.com/tradersclub/PocArquitetura/app"
	"github.com/tradersclub/PocArquitetura/model"
)

// Register group item check
func Register(g *echo.Group, apps *app.Container, m *middleware.Middleware) {
	h := &handler{
		apps: apps,
	}

	g.GET("", h.getItem, m.Auth.Private)
	g.POST("", h.postItem, m.Auth.Private)
}

type handler struct {
	apps *app.Container
}

// getItem swagger document
// @Summary Exembro de Buscar item por id
// @Description Essa rota é privada com o token valido (Bearer)
// @Tags item
// @Accept  json
// @Produce  json
// @Param id query string true "id"
// @Success 200 {object} model.Item
// @Failure 400 {object} string
// @Security ApiKeyAuth
// @Router /v1/item [get]
func (h *handler) getItem(c echo.Context) error {
	ctx := c.Request().Context()

	// recuperando query param
	id := c.QueryParam("id")

	resp, err := h.apps.Item.RequestItemById(ctx, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

// postItem swagger document
// @Summary Exembro de como postar algum item
// @Description Essa rota é privada com o token valido (Bearer)
// @Tags item
// @Accept  json
// @Produce  json
// @Param item body model.Item true "add Item"
// @Success 200 {object} model.Item
// @Failure 400 {object} string
// @Security ApiKeyAuth
// @Router /v1/item [post]
func (h *handler) postItem(c echo.Context) error {
	ctx := c.Request().Context()

	body, err := c.Request().GetBody()
	if err != nil {
		return err
	}

	// recuperando body
	item := model.ItemFromJson(body)

	resp, err := h.apps.Item.RequestItemById(ctx, item.ID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}
