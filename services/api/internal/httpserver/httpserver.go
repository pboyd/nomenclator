// Package httpserver provides an HTTP server for the API.
package httpserver

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"go.uber.org/fx"
)

// Config contains configuration options for the server.
type Config struct {
	Addr     string
	CertFile string
	KeyFile  string
}

// New starts a new HTTP server.
func New(lc fx.Lifecycle, config *Config, handler http.Handler) *http.Server {
	server := &http.Server{
		Handler:           handler,
		ReadHeaderTimeout: 5 * time.Second,
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			listener, err := net.Listen("tcp", config.Addr)
			if err != nil {
				return fmt.Errorf("failed to listen on %s: %w", config.Addr, err)
			}

			go func() {
				var err error
				if config.CertFile == "" && config.KeyFile == "" {
					err = server.Serve(listener)
				} else {
					err = server.ServeTLS(listener, config.CertFile, config.KeyFile)
				}

				if err != nil && !errors.Is(err, http.ErrServerClosed) {
					log.Fatal(err.Error())
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})

	return server
}
