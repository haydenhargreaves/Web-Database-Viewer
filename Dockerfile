FROM golang:1.25.7-bookworm AS build

WORKDIR /src

COPY go.mod go.sum ./
COPY vendor ./vendor
COPY cmd ./cmd
COPY internal ./internal

RUN CGO_ENABLED=1 go build -mod=vendor -trimpath -ldflags="-s -w" -o /out/web_server ./cmd/web_server.go

FROM debian:bookworm-slim

RUN apt-get update \
    && apt-get install --no-install-recommends -y ca-certificates \
    && rm -rf /var/lib/apt/lists/* \
    && useradd --uid 10001 --create-home app

WORKDIR /app

COPY --from=build /out/web_server ./web_server
COPY web ./web
COPY assets ./assets

USER app

EXPOSE 3001

ENTRYPOINT ["./web_server"]
