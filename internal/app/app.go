package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/necutya-diploma/notification-service/internal/config"
	"github.com/necutya-diploma/notification-service/internal/service"
	"github.com/necutya-diploma/notification-service/internal/transport"
	"github.com/necutya-diploma/notification-service/pkg/mailer"
	log "github.com/sirupsen/logrus"
)

func Run(configPath string) {
	cfg, err := config.Load(configPath)
	if err != nil {
		log.Fatalln(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	go gracefulShutdown(cancel)

	emailManager := mailer.NewSMTPMailer(
		cfg.Email.Host,
		cfg.Email.Port,
		cfg.Email.Username,
		cfg.Email.Password,
	)
	if cfg.IsDev() {
		emailManager.SetDevEnv()
	}

	services := service.New(emailManager)

	grpcServer := transport.NewGRPC(cfg.GRPC.Host, cfg.GRPC.Port, services.Email)
	grpcServer.Run(ctx)
}

func gracefulShutdown(stop func()) {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	<-signalChannel
	stop()
}
