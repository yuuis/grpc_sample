package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"pancake.maker/gen/api"
	"pancake.maker/handler"
)

func main() {
	unaryPort := 50051
	streamPort := 50052
	unaryLis, err := net.Listen("tcp", fmt.Sprintf(":%d", unaryPort))
	if err != nil {
		log.Fatalf("failed to listern: %v", err)
	}
	streamLis, err := net.Listen("tcp", fmt.Sprintf(":%d", streamPort))
	if err != nil {
		log.Fatalf("failed to listern: %v", err)
	}

	// logger
	zapLogger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("failed to create logger: %v", err)
	}
	grpc_zap.ReplaceGrpcLoggerV2(zapLogger)

	// unary
	unaryServer := grpc.NewServer(grpc.UnaryInterceptor(
		grpc_middleware.ChainUnaryServer(
			grpc_zap.UnaryServerInterceptor(zapLogger),
			grpc_auth.UnaryServerInterceptor(auth),
		),
	))
	api.RegisterPancakeBakerServiceServer(unaryServer, handler.NewBakerHandler())
	reflection.Register(unaryServer)

	// stream
	streamServer := grpc.NewServer(grpc.StreamInterceptor(
		grpc_middleware.ChainStreamServer(
			grpc_zap.StreamServerInterceptor(zapLogger),
			//grpc_auth.StreamServerInterceptor(auth),
		),
	))
	api.RegisterImageUploadServiceServer(streamServer, handler.NewImageUploadHandler())
	reflection.Register(streamServer)

	go func() {
		log.Printf("start gRPC unary server prot %v\n", unaryPort)
		unaryServer.Serve(unaryLis)
	}()

	go func() {
		log.Printf("start gRPC stream server prot %v\n", streamPort)
		streamServer.Serve(streamLis)
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	unaryServer.GracefulStop()
	streamServer.GracefulStop()
}

func auth(ctx context.Context) (context.Context, error) {
	validToken := "hi/mi/tsu"
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}

	if token != validToken {
		return nil, status.Errorf(codes.Unauthenticated, "invalid bearer token")
	}

	return context.WithValue(ctx, "UserName", "God"), nil
}
