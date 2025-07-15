FROM golang:1.24 AS build

ENV BIN_FILE /opt/banner/banner-app
ENV CODE_DIR /go/src

WORKDIR ${CODE_DIR}

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . ${CODE_DIR}

ARG LDFLAGS
RUN CGO_ENABLED=0 go build \
        -ldflags "$LDFLAGS" \
        -o ${BIN_FILE} cmd/banner/*

FROM alpine:3.9

LABEL ORGANIZATION="Rotation Banner Microservice"
LABEL SERVICE="banner"

ENV BIN_FILE "/opt/banner/banner-app"
COPY --from=build ${BIN_FILE} ${BIN_FILE}

ENV CONFIG_FILE /etc/banner/config.yaml
COPY configs/config.yaml ${CONFIG_FILE}

CMD ${BIN_FILE} -config /etc/banner