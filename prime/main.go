package main

import (
	"net"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"tratnik.net/prime/internal/config"
	"tratnik.net/prime/internal/handler"
	"tratnik.net/prime/internal/repository"
	"tratnik.net/prime/internal/service"
	"tratnik.net/prime/pkg/pb"
	"tratnik.net/prime/pkg/postgres"
)

func main() {
	c := config.GetFromFile()

	// A workaround for lack of a deployment script. Migrate should not be called here.
	// Wanted everything to fit into a single clean docker-compose.
	time.Sleep(time.Second * 5)
	postgres.Migrate(c.Database)

	db := postgres.New(c.Database)

	validationRepo := repository.NewValidation(db)
	primeService := service.NewPrime(validationRepo)
	primeHandler := handler.NewPrime(primeService)

	logrus.WithField("port", c.Server.Port).Info("Server starting")
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(c.Server.Port))
	if err != nil {
		logrus.WithError(err).Fatal("Unable to start server")
	}
	grpcServer := grpc.NewServer()
	pb.RegisterPrimeServer(grpcServer, primeHandler)
	grpcServer.Serve(listener)
}
