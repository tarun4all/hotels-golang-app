# build stage for server
FROM golang:alpine AS build-env-server
RUN apk --no-cache add build-base
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -o go_server_app ./cmd/server/

# build stage for server
FROM golang:alpine AS build-env-cli
RUN apk --no-cache add build-base
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -o go_cli_app ./cmd/cli/

# final stage
FROM alpine:latest AS production
COPY --from=build-env /app/go_server_app /app/
COPY --from=build-env /app/go_cli_app /app/
ENTRYPOINT ./app/go_server_app