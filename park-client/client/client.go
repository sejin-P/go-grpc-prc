package client

import (
	"sync"

	"google.golang.org/grpc"

	sejinpb "github.com/sejin-P/go-grpc/protos/v1/sejin"
)

var (
	once sync.Once
	cli  sejinpb.SejinClient
)

func GetSejinClient(servicehost string) sejinpb.SejinClient {
	once.Do(func() {
		conn, _ := grpc.Dial(servicehost,
			grpc.WithInsecure(),
			grpc.WithBlock(),
		)

		cli = sejinpb.NewSejinClient(conn)
	})

	return cli
}
