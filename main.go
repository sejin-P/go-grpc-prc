package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"github.com/sejin-P/go-grpc/api"
	sejinpb "github.com/sejin-P/go-grpc/protos/v1/sejin"
)

const (
	portNum     = "9002"
	httpPortNum = "9000"
)

type sejinServer struct {
	sejinpb.SejinServer
}

func (s *sejinServer) GetUserInfo(ctx context.Context, req *sejinpb.GetInfoRequest) (*sejinpb.GetInfoResponse, error) {
	id := req.Id

	if id != api.SejinData.Id {
		return nil, errors.New("Not found")
	}

	return &sejinpb.GetInfoResponse{Message: &api.SejinData}, nil
}

func main() {
	ctx := context.Background()
	mux := runtime.NewServeMux()

	options := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	grpcServer := grpc.NewServer()
	go func() {
		lis, err := net.Listen("tcp", ":"+portNum)

		if err != nil {
			log.Panicf("failed to listen, %v", err)
			return
		}
		log.Printf("grpc server started on %s port", portNum)

		if err := grpcServer.Serve(lis); err != nil {
			log.Panicf("failed to serve grpc server, %v", err)
		}
	}()

	httpServer := &http.Server{
		Addr:    ":" + httpPortNum,
		Handler: mux,
	}

	go func() {
		if err := sejinpb.RegisterSejinHandlerFromEndpoint(
			ctx,
			mux,
			"localhost:"+portNum,
			options,
		); err != nil {
			log.Panicf("failed to register grpc gateway %v", err)
		}

		log.Printf("started HTTP server on %v", httpPortNum)
		if err := http.ListenAndServe(":"+httpPortNum, mux); err != nil {
			log.Panicf("failed to serve %v", err)
		}

	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(quit)

	<-quit

	time.Sleep(time.Duration(15) * time.Second)

	log.Println("stopping heenam HTTP server")
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	log.Print("stopping heenam gRPC server")
	grpcServer.GracefulStop()

}
