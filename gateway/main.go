package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"tratnik.net/gateway/internal/config"
	"tratnik.net/gateway/internal/handler"
	"tratnik.net/gateway/internal/middleware"
	"tratnik.net/gateway/internal/service"
	"tratnik.net/gateway/pkg/http/response"
	"tratnik.net/gateway/pkg/http/server"
	"tratnik.net/gateway/pkg/pb"
)

func main() {
	c := config.GetFromFile()

	router := InitRouter()

	conn, err := grpc.Dial(c.Prime.Address, grpc.WithInsecure())
	if err != nil {
		logrus.WithError(err).Fatal("Unable to open connection to Prime")
	}
	defer conn.Close()
	primeClient := pb.NewPrimeClient(conn)

	primeSrvc := service.NewPrime(primeClient)
	_ = handler.NewPrime(router, primeSrvc)

	server.Serve(c.Server, router)
}

func InitRouter() *mux.Router {
	r := mux.NewRouter()

	r.Use(
		middleware.Recoverer(response.JSON),
	)

	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		response.JSON(w, nil, http.StatusNotFound)
	})

	r.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		response.JSON(w, nil, http.StatusMethodNotAllowed)
	})

	return r
}
