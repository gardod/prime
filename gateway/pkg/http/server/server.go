package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/sirupsen/logrus"
)

type Config struct {
	Port         int `mapstructure:"port"`
	ReadTimeout  int `mapstructure:"read_timeout"`
	WriteTimeout int `mapstructure:"write_timeout"`
	IdleTimeout  int `mapstructure:"idle_timeout"`
}

func Serve(config Config, handler http.Handler) {
	logrus.Info("Server starting")

	h2s := &http2.Server{
		IdleTimeout: time.Second * time.Duration(config.IdleTimeout),
	}

	h1s := &http.Server{
		Addr:         ":" + strconv.Itoa(config.Port),
		ReadTimeout:  time.Second * time.Duration(config.ReadTimeout),
		WriteTimeout: time.Second * time.Duration(config.WriteTimeout),
		IdleTimeout:  time.Second * time.Duration(config.IdleTimeout),
		Handler:      h2c.NewHandler(handler, h2s),
	}

	go func() {
		if err := h1s.ListenAndServe(); err != http.ErrServerClosed {
			logrus.WithError(err).Fatal("Unable to start server")
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := h1s.Shutdown(ctx); err != nil {
		logrus.WithError(err).Fatal("Unable to gracefully shut down server")
	}

	logrus.Info("Server shut down")
}
