package main

import (
	"context"
	"github.com/husanmusa/bpro_api_gateway/api"
	"github.com/husanmusa/bpro_api_gateway/api/handlers"
	"github.com/husanmusa/bpro_api_gateway/config"
	"github.com/husanmusa/bpro_api_gateway/grpc/client"
	"github.com/saidamir98/udevs_pkg/logger"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var loggerLevel string
	cfg := config.Load()

	switch cfg.Environment {
	case config.DebugMode:
		loggerLevel = logger.LevelDebug
	case config.TestMode:
		loggerLevel = logger.LevelDebug
	default:
		loggerLevel = logger.LevelInfo
	}

	log := logger.NewLogger(cfg.ServiceName, loggerLevel)
	defer func() {
		if err := logger.Cleanup(log); err != nil {
			log.Error("Failed to cleanup logger", logger.Error(err))
		}
	}()

	svcs, err := client.NewGrpcClients(cfg)
	if err != nil {
		log.Panic("client.NewGrpcClients", logger.Error(err))
	}

	log.Info("SUCCESSFUL connect to redis")

	h := handlers.NewHandler(cfg, log, svcs)

	r := api.SetUpRouter(h, cfg, log)

	group, _ := errgroup.WithContext(context.Background())

	group.Go(func() error {
		if r.Listen(cfg.HTTPPort) != nil {
			log.Error("could not start http server", logger.Error(err))
			return nil
		}
		return nil
	})
	log.Info("server started")

	shutdownChan := make(chan os.Signal, 1)
	defer close(shutdownChan)
	signal.Notify(shutdownChan, syscall.SIGTERM, syscall.SIGINT)
	sig := <-shutdownChan

	log.Info("received os signal", logger.Any("signal", sig))
	if err := r.Shutdown(); err != nil {
		log.Error("could not shutdown http server", logger.Error(err))
		return
	}

	log.Info("server shutdown successfully")
}
