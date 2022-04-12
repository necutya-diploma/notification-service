package transport

import (
	"context"
	"fmt"
	"net"

	"github.com/necutya-diploma/notification-service/internal/transport/handlers"
	"github.com/necutya-diploma/notification-service/pkg/grpc/gen"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcLogrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	log "github.com/sirupsen/logrus"
)

type GRPCServer struct {
	addr   string
	server *grpc.Server
}

func NewGRPC(host string, port int, emailService handlers.EmailService) *GRPCServer {
	grpcServer := &GRPCServer{
		addr: fmt.Sprintf("%v:%v", host, port),
		server: grpc.NewServer(
			grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(
				grpcLogrus.UnaryServerInterceptor(log.NewEntry(log.StandardLogger())),
				grpcRecovery.UnaryServerInterceptor(),
			)),
			grpc.StreamInterceptor(grpcMiddleware.ChainStreamServer(
				grpcLogrus.StreamServerInterceptor(log.NewEntry(log.StandardLogger())),
				grpcRecovery.StreamServerInterceptor(),
			)),
		),
	}

	gen.RegisterMailerServer(grpcServer.server, handlers.New(emailService))

	return grpcServer
}

func (s *GRPCServer) Run(ctx context.Context) {
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Error("grpc srv: run error: ", err)

		return
	}

	g, gCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		log.Infof("Starting grpc server: addr=%v", s.addr)
		return s.server.Serve(listener)
	})

	g.Go(func() error {
		<-gCtx.Done()
		log.Info("Stopping grpc server...")
		s.server.GracefulStop()
		return nil
	})

	if err := g.Wait(); err != nil {
		log.Info("Grpc server has been stopped successfully")
	}
}
