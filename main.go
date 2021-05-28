package main

import (
	"net/http"
	"time"

	"github.com/GianGoulart/Clinica_backend/api/middleware"
	"github.com/GianGoulart/Clinica_backend/api/swagger"
	"github.com/sirupsen/logrus"

	"github.com/GianGoulart/Clinica_backend/model"

	"github.com/GianGoulart/Clinica_backend/api"
	"github.com/GianGoulart/Clinica_backend/app"
	"github.com/GianGoulart/Clinica_backend/store"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	emiddleware "github.com/labstack/echo/v4/middleware"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/spf13/viper/remote"
)

// main configure swagger
//
// method of use bearer token in requests
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	startedAt := time.Now()

	model.Watch(func(c model.Config, quit chan bool) {
		e := echo.New()
		e.Validator = model.New()
		e.Debug = c.GetString("tc") != "prod"
		e.HideBanner = true

		e.Use(emiddleware.Logger())
		e.Use(emiddleware.BodyLimit("2M"))
		e.Use(emiddleware.Recover())
		e.Use(emiddleware.RequestID())

		dbWriter := sqlx.MustConnect("mysql", c.GetString("database.writer.url"))
		dbReader := sqlx.MustConnect("mysql", c.GetString("database.reader.url"))

		// criação dos stores com a injeção do banco de escrita e leitura
		stores := store.New(store.Options{
			Writer: dbWriter,
			Reader: dbReader,
		})

		// criação dos serviços
		apps := app.New(app.Options{
			Stores:    stores,
			Version:   c.GetString("version"),
			StartedAt: startedAt})

		// registros dos handlers
		api.Register(api.Options{
			Group: e.Group(""),
			Apps:  apps,

			// criação e injeção dos middlewares
			Middleware: middleware.New(middleware.Options{
				Apps: apps,
			}),
		})

		port := c.GetString("server.port")
		if e.Debug {
			swagger.Register(swagger.Options{
				Port:      port,
				Group:     e.Group("/swagger"),
				AccessKey: c.GetString("docs.key"),
			})
		}

		// funcão padrão pra tratamento de erros da camada http
		e.HTTPErrorHandler = func(err error, c echo.Context) {
			if c.Response().Committed {
				return
			}

			if err := c.JSON(http.StatusInternalServerError, model.Response{Err: err}); err != nil {
				logrus.Error(c.Request().Context(), err)
			}
		}

		// função para fechar as conexões
		go func() {
			<-quit

			dbReader.Close()
			dbWriter.Close()
			e.Close()
		}()

		go e.Start(port)

		logrus.Info("Microservice started!")
	})
}
