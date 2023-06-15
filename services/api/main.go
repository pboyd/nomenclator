// Package main service entrypoint
package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"go.uber.org/fx"

	"github.com/pboyd/nomenclator/api/internal/database"
	"github.com/pboyd/nomenclator/api/internal/domain"
	"github.com/pboyd/nomenclator/api/internal/httpserver"
)

func main() {
	fxApp := fx.New(
		fx.Provide(readHTTPServerConfig),
		fx.Provide(readDatabaseConfig),
		fx.Provide(database.Connect),

		fx.Provide(domain.NewBundle),

		fx.Provide(httpserver.NewHandler),
		fx.Provide(httpserver.New),
		fx.Invoke(database.Migrate),
		fx.Invoke(func(s *http.Server) {}),
	)

	start, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	if err := fxApp.Start(start); err != nil {
		log.Fatalf("error starting service %v", err)
	}

	<-fxApp.Done()

	stop, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	if err := fxApp.Stop(stop); err != nil {
		log.Fatalf("error stopping service %v", err)
	}
}
