package item

import (
	"net/http"

	"github.com/GianGoulart/Clinica_backend/api/middleware"
	"github.com/GianGoulart/Clinica_backend/app"
	"github.com/GianGoulart/Clinica_backend/model"
	"github.com/labstack/echo/v4"
)

// Register group item check
func Register(g *echo.Group, apps *app.Container, m *middleware.Middleware) {
	h := &handler{
		apps: apps,
	}

	g.POST("", h.postItem, m.Auth.Private)
}

type handler struct {
	apps *app.Container
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
	// ctx := c.Request().Context()

	body, err := c.Request().GetBody()
	if err != nil {
		return err
	}

	// recuperando body
	item := model.ItemFromJson(body)

	// resp, err := h.apps.Item.RequestItemById(ctx, item.ID)
	// if err != nil {
	// 	return err
	// }

	return c.JSON(http.StatusOK, item)
}
