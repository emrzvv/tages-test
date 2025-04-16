package main

import (
	"fmt"
	"log"
	"net"

	"github.com/emrzvv/tages-test/cfg"
	"github.com/emrzvv/tages-test/internal/app/limiter"
	"github.com/emrzvv/tages-test/internal/app/server"
	"github.com/emrzvv/tages-test/internal/app/service"
	"github.com/emrzvv/tages-test/internal/app/storage"
	pb "github.com/emrzvv/tages-test/proto"
	"google.golang.org/grpc"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("failed to run: %v", err)
	}
}

func run() error {
	config := cfg.LoadNewDefaultConfig()
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", config.Host, config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	limiter := limiter.NewCounterLimiter(config)
	s := grpc.NewServer(
		grpc.UnaryInterceptor(limiter.UnaryInterceptor),
		grpc.StreamInterceptor(limiter.StreamInterceptor),
	)

	fs := storage.NewSimpleFileStorage(config.StoragePath)
	ms, err := storage.NewSQLiteMetaStorage(config)
	if err != nil {
		return err
	}
	// ms := storage.NewInMemoryMetaStorage(config)

	service := service.NewImgService(config, fs, ms)
	srvr := server.NewServer(config, service)

	pb.RegisterImageServiceServer(s, srvr)
	log.Printf("server listening at %v\n", lis.Addr())

	if err := s.Serve(lis); err != nil {
		return err
	}
	return nil
}
