package grpc

import (
	"context"
	"net"
	"os"
	"os/signal"

	v1 "github.com/martinyonathann/grpc-http-rest-microservice-go/pkg/api/v1"
	"github.com/martinyonathann/grpc-http-rest-microservice-go/pkg/logger"
	"github.com/martinyonathann/grpc-http-rest-microservice-go/pkg/protocol/grpc/middleware"
	"google.golang.org/grpc"
)

//RunServer runs gRPC service to publish ToDo service
func RunServer(ctx context.Context, v1API v1.ToDoServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	//gRPC server statup options
	opts := []grpc.ServerOption{}

	//add middleware
	opts = middleware.AddLogging(logger.Log, opts)

	//register service
	server := grpc.NewServer()
	v1.RegisterToDoServiceServer(server, v1API)

	//graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			//sig is a ^C, handle it
			logger.Log.Warn("shutting down gRPC server...")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	// start gRPC server
	logger.Log.Warn("Starting gRPC server...")
	return server.Serve(listen)

}
