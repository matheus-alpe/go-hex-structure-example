package rpc

import (
	"log"
	"net"

	"github.com/matheus-alpe/go-hex-structure-example/internal/adapters/framework/left/grpc/pb"
	"github.com/matheus-alpe/go-hex-structure-example/internal/ports"
	"google.golang.org/grpc"
)

type Adapter struct {
	api ports.APIPort
}

func NewAdapter(api ports.APIPort) *Adapter {
	return &Adapter{api: api}
}

func (grpca Adapter) Run() {
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen on port 9000: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterArithmeticServiceServer(grpcServer, grpca)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve gRPC server: %v", err)
	}
}
