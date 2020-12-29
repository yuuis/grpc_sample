package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
	"pancake.maker/gen/api"
	"pancake.maker/handler"
)

func main() {
	port := 50051
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listern: %v", err)
	}

	server := grpc.NewServer()

	// .protoファイルで定義したPancakeBakerServiceに対応している。
	// gRPCのサーバーにPancakeBakerServiceへのリクエストがあった場合に、渡したハンドラの対応したメソッドを呼び出すよう登録する。
	api.RegisterPancakeBakerServiceServer(server, handler.NewBakerHandler())
	reflection.Register(server)

	go func() {
		log.Printf("start gRPC server prot %v", port)
		server.Serve(lis)
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	server.GracefulStop()
}
