BIN := "./bin/banner"

GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

build:
	go build -v -o $(BIN) -ldflags "$(LDFLAGS)" ./cmd/banner

run: build
	$(BIN) -config ./configs

version: build
	$(BIN) version

docker-up:
	docker compose -f deployments/docker-compose.yaml up -d

docker-down:
	docker compose -f deployments/docker-compose.yaml down --remove-orphans

# Применить миграции с помощью goose
migrate-up:
	goose up -dir migrations

# Откатить миграции с помощью goose
migrate-down:
	goose down -dir migrations