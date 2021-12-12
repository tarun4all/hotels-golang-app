# build stage
FROM golang:alpine AS build-env
RUN apk --no-cache add build-base
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -o goapp ./...

# final stage
FROM alpine:latest AS production
COPY --from=build-env /app/goapp /app/
ENTRYPOINT ./goapp