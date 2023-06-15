package main

import (
	"log"
	"net/url"
	"os"

	"github.com/pboyd/nomenclator/api/internal/database"
	"github.com/pboyd/nomenclator/api/internal/httpserver"
)

func readHTTPServerConfig() *httpserver.Config {
	return &httpserver.Config{
		Addr:     envString("LISTEN_ADDR", ":8080"),
		CertFile: envString("CERT_FILE", ""),
		KeyFile:  envString("KEY_FILE", ""),
	}
}

func readDatabaseConfig() *database.Config {
	dsn := envString("DATABASE_DSN", "")
	if dsn == "" {
		passwordFile := envString("POSTGRES_PASSWORD_FILE", "/run/secrets/postgres_password")
		//nolint:gosec
		password, err := os.ReadFile(passwordFile)
		if err != nil {
			log.Fatalf("unable to read postgres password file: %v", err)
		}

		dsn = (&url.URL{
			Scheme:   "postgres",
			User:     url.UserPassword("postgres", string(password)),
			Host:     envString("DATABASE_HOST", "db"),
			Path:     envString("DATABASE_NAME", "nomenclator"),
			RawQuery: "sslmode=disable",
		}).String()
	}

	return &database.Config{
		DSN:           dsn,
		MigrationsDir: envString("DATABASE_MIGRATIONS_PATH", "/app/migrations"),
	}
}

func envString(key, def string) string {
	value := os.Getenv(key)
	if value == "" {
		return def
	}
	return value
}
