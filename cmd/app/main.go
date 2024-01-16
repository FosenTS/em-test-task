package main

import (
	"context"
	"em-test-task/internal/application"
	"em-test-task/internal/application/config"
	"fmt"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()

	err := config.Env()
	if err != nil {
		log.Fatalln(fmt.Errorf("fatal reading env config: %w", err))
		return
	}

	appConfig, err := config.App()
	if err != nil {
		log.Fatalln(err)
		return
	}

	log.SetLevel(appConfig.LogLevel)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	app, err := application.NewApp(ctx, log.WithField("location", "application"), appConfig)
	if err != nil {
		log.Fatalln("fatal error creating application")
		return
	}
	err = app.Run(ctx, log.WithField("location", "runner"))
	if err != nil {
		log.Fatalln("fatal run application")
		return
	}
}
