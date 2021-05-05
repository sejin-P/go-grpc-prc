package main

import (
	"context"
	"errors"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/sejin-P/go-grpc/api"
	"github.com/sejin-P/go-grpc/park-client/client"
	parkpb "github.com/sejin-P/go-grpc/protos/v1/park"
	sejinpb "github.com/sejin-P/go-grpc/protos/v1/sejin"
)

const portNum = "9001"

type parkServer struct {
	parkpb.ParkServer

	sejinCli sejinpb.SejinClient
}

func (s *parkServer) GetPark(ctx context.Context, req *parkpb.GetParkRequest) (*parkpb.GetParkResponse, error) {
	id := req.Id

	if id != api.ParkData.Id {
		return nil, errors.New("not found")
	}

	return &parkpb.GetParkResponse{Message: &api.ParkData}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":"+portNum)
	if err != nil {
		log.Panicf("cannot listen")
	}

	sejinCli := client.GetSejinClient("localhost:9000")
	grpcServer := grpc.NewServer()
	parkpb.RegisterParkServer(grpcServer, &parkServer{
		sejinCli: sejinCli,
	})

	log.Printf("park server started on %s", portNum)

	if err := grpcServer.Serve(lis); err != nil {
		log.Panicf("failed to serve %v", err)
	}
}
