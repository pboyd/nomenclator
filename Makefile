.PHONY: all
all: deps secrets build

.PHONY: secrets
secrets: secrets/postgres_password.txt secrets/localhost.pem

secrets/postgres_password.txt:
	@mkdir -p secrets
	@./bin/rand > secrets/postgres_password.txt

secrets/localhost.pem:
	@mkdir -p secrets
	@cd secrets && ../bin/gencert localhost

.PHONY: build
build:
	@[ -n "$$(git status --porcelain)" ] && echo "git working directory is not clean" && exit 1
	BUILD_TAG=$$(git rev-parse --short HEAD) docker compose build

.PHONY: deps
deps:
	go install github.com/kyleconroy/sqlc/cmd/sqlc@v1.18.0
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.2

.PHONY: gen
gen:
	./bin/foreach_mod go generate

.PHONY: lint
lint:
	./bin/foreach_mod golangci-lint run ./...

.PHONY: tidy
tidy:
	./bin/foreach_mod go mod tidy

.PHONY: test
test:
	./bin/foreach_mod go test ./...
