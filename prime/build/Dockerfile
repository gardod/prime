FROM golang:1.14-alpine3.11 AS build-env
WORKDIR /go/src/tratnik.net/prime
COPY . /go/src/tratnik.net/prime
RUN go build -o /server

FROM alpine:3.11
WORKDIR /opt/prime
COPY ./migrations ./migrations
COPY ./config/dev.yaml ./config.yaml
COPY --from=build-env /server .
EXPOSE 9000
ENTRYPOINT ["./server"]
