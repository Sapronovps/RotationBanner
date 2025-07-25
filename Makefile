BIN := "./bin/banner"

GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

build:
	go build -v -o $(BIN) -ldflags "$(LDFLAGS)" ./cmd/banner

run: build
	$(BIN) -config ./configs

version: build
	$(BIN) version

# Поднять проект в docker (приложение + зависимости)
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

# Сгенерировать реализацию Proto
proto-gen:
	cd api && protoc --go_out=../internal/server/grpc/protobuf --go_opt=paths=source_relative --go-grpc_out=../internal/server/grpc/protobuf --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative BannerService.proto

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.64.8

lint: install-lint-deps
	golangci-lint run ./...

test:
	go test -race ./internal/...