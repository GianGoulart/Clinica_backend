package main

import (
	"time"

	"github.com/tradersclub/TCUtils/natstan"

	"github.com/tradersclub/PocArquitetura/api/middleware"
	"github.com/tradersclub/PocArquitetura/api/swagger"
	"github.com/tradersclub/PocArquitetura/event"

	"github.com/tradersclub/TCUtils/cache"

	"github.com/tradersclub/TCUtils/logger"

	"github.com/tradersclub/PocArquitetura/model"
	"github.com/tradersclub/TCUtils/tcerr"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	emiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/tradersclub/PocArquitetura/api"
	"github.com/tradersclub/PocArquitetura/app"
	"github.com/tradersclub/PocArquitetura/store"
	"github.com/tradersclub/TCUtils/config"
	"github.com/tradersclub/TCUtils/validator"

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

	config.Watch(func(c config.Config, quit chan bool) {
		e := echo.New()
		e.Validator = validator.New()
		e.Debug = c.GetString("tc") != "prod"
		e.HideBanner = true

		e.Use(emiddleware.Logger())
		e.Use(emiddleware.BodyLimit("2M"))
		e.Use(emiddleware.Recover())
		e.Use(emiddleware.RequestID())

		// // middleware do prometheus
		p := prometheus.NewPrometheus("PocArquitetura", nil)
		p.Use(e)

		dbWriter := sqlx.MustConnect("mysql", c.GetString("database.writer.url"))
		dbReader := sqlx.MustConnect("mysql", c.GetString("database.reader.url"))

		natsConn := natstan.New(natstan.Options{
			URL: c.GetString("nats.url"),
		})

		// criação dos stores com a injeção do banco de escrita e leitura
		stores := store.New(store.Options{
			Writer: dbWriter,
			Reader: dbReader,
		})

		// criação dos serviços
		apps := app.New(app.Options{
			Stores:    stores,
			Version:   c.GetString("version"),
			StartedAt: startedAt,

			// criação e injeção da conexção com o cache
			Cache: cache.NewMemcache(cache.Options{
				URL:        c.GetString("cache.url"),
				Expiration: c.GetDuration("cache.expiration"),
			}),

			// criação e injeção das conexções do stan e nats
			Nats: natsConn,
		})

		event.Register(event.Options{
			Apps: apps,
			Nats: natsConn,
		})

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

			if err := c.JSON(tcerr.GetHTTPCode(err), model.Response{Err: err}); err != nil {
				logger.ErrorContext(c.Request().Context(), err)
			}
		}

		// função para fechar as conexões
		go func() {
			<-quit

			dbReader.Close()
			dbWriter.Close()
			natsConn.Close()
			e.Close()
		}()

		go e.Start(port)

		logger.Info("Microservice started!")
	})
}
