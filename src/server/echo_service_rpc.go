package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"

	proto "github.com/savarin/echo-go/proto"
)

// EchoServiceRpc implements the gRPC server interface for EchoService.
type EchoServiceRpc struct {
	proto.UnimplementedEchoServiceServer
	echoService *EchoService
}

// NewEchoServiceRpc creates a new EchoServiceRpc.
func NewEchoServiceRpc(echoService *EchoService) *EchoServiceRpc {
	return &EchoServiceRpc{echoService: echoService}
}

// Echo implementation of the Echo RPC method. It delegates the actual logic to the EchoService.
func (s *EchoServiceRpc) Echo(ctx context.Context, request *proto.EchoRequest) (*proto.EchoResponse, error) {
	message := s.echoService.Echo(request.GetMessage())

	// Build and return the response by encapsulating the echoed message
	response := &proto.EchoResponse{
		Message: message,
	}
	return response, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("Failed to listen on port 8080: %v\n", err)
		return
	}
	grpcServer := grpc.NewServer()

	echoService := &EchoService{} // You'll need to define or import EchoService
	rpcService := NewEchoServiceRpc(echoService)
	proto.RegisterEchoServiceServer(grpcServer, rpcService)

	fmt.Println("Server started, listening on 8080")
	if err := grpcServer.Serve(lis); err != nil {
		fmt.Printf("Failed to serve gRPC server: %v\n", err)
	}
}
