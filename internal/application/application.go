package application

import (
	"context"
	"em-test-task/internal/application/config"
	"em-test-task/internal/application/product"
	"em-test-task/pkg/db/postgresql"
	"fmt"
	"github.com/AubSs/fasthttplogger"
	"github.com/fasthttp/router"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"golang.org/x/sync/errgroup"
	"strconv"
	"time"
)

type App interface {
	Run(ctx context.Context, log *logrus.Entry) error
	runHTTP(log *logrus.Entry) error
}

type app struct {
	endpoint   *product.Endpoint
	httpConfig *config.HTTPConfig
	log        *logrus.Entry
	appConfig  *config.AppConfig
}

func NewApp(ctx context.Context, log *logrus.Entry, appConfig *config.AppConfig) (App, error) {
	log.Println("Start creating application")

	postgresConfig, err := config.Postgre()
	if err != nil {
		log.Errorln(err)
		return nil, err
	}

	httpConfig, err := config.Http()
	if err != nil {
		log.Errorln(err)
		return nil, err
	}

	pgxPool, err := postgresql.NewPGXPool(
		ctx,
		&postgresql.ClientConfig{
			Host:                  postgresConfig.Host,
			Port:                  strconv.Itoa(postgresConfig.Port),
			Username:              postgresConfig.User,
			Password:              postgresConfig.Password,
			DatabaseName:          postgresConfig.DatabaseName,
			UseCA:                 postgresConfig.SSLMode,
			SimpleQueryMode:       true,
			WaitingDuration:       1000 * time.Second,
			MaxConnectionAttempts: 3,
			MaxConnections:        15,
		},
		log.WithField("location", "pgx"),
	)

	log.Infoln("Creating storages")
	storages := product.NewStorages(pgxPool, log.WithField("location", "storages"))

	log.Infoln("Creating services")
	services := product.NewServices(storages, log.WithField("location", "services"))

	log.Infoln("Creating gateways")
	gateways, err := product.NewGateways(services, log.WithField("location", "gateways"))
	if err != nil {
		return nil, err
	}

	log.Infoln("Creating controllers")
	controllers := product.NewControllers(gateways, log.WithField("location", "controllers"))

	log.Infoln("Creating endpoints")
	endpoint := product.NewEndpoint(controllers)

	return &app{
		endpoint:   endpoint,
		httpConfig: httpConfig,
		log:        log,
		appConfig:  appConfig,
	}, nil
}

func (app *app) Run(ctx context.Context, log *logrus.Entry) error {
	grp, ctx := errgroup.WithContext(ctx)

	grp.Go(func() error {
		return app.runHTTP(log.WithField("location", "runHTTP"))
	})

	return grp.Wait()
}

func (app *app) runHTTP(log *logrus.Entry) error {
	app.log.Infoln("Configure api routes")

	r := router.New()
	app.endpoint.RegisterRouter(r.Group("/api"))
	list := r.List()
	for key, routesGroup := range list {
		for _, route := range routesGroup {
			log.Infoln(fmt.Sprintf("%s : %s", key, route))
		}
	}

	err := fasthttp.ListenAndServe(app.httpConfig.HttpServer, fasthttplogger.CombinedColored(r.Handler))
	if err != nil {
		log.Fatalln(fmt.Sprintf("Error start listen and serve http server: %w", err))
		return err
	}
	return nil
}
