FROM golang:1.14-alpine3.11 AS build-env
WORKDIR /go/src/tratnik.net/gateway
COPY . /go/src/tratnik.net/gateway
RUN go build -o /server

FROM alpine:3.11
WORKDIR /opt/gateway
COPY ./config/dev.yaml ./config.yaml
COPY --from=build-env /server .
ENTRYPOINT ["./server"]
