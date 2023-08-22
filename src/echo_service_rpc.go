package main

import (
	"context"

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
