package server

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"net"
	"testing"

	proto "github.com/savarin/echo-go/proto"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestEchoRPC(t *testing.T) {
	// Create in-memory listener and server
	grpcServer := grpc.NewServer()
	s := NewEchoServiceRpc(&EchoService{}) // Adjust to your EchoService creation
	proto.RegisterEchoServiceServer(grpcServer, s)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			fmt.Println("Server exited with error:", err)
		}
	}()

	// Create client connection
	conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := proto.NewEchoServiceClient(conn)

	// Arrange: Define request message
	message := "Hello, World!"
	request := &proto.EchoRequest{Message: message}

	// Act: Invoke RPC call and get the response
	response, err := client.Echo(context.Background(), request)
	if err != nil {
		t.Fatalf("Failed to call Echo: %v", err)
	}

	// Assert: Verify that the response is as expected
	expectedResponse := &proto.EchoResponse{Message: "13:" + message}
	if response.GetMessage() != expectedResponse.GetMessage() {
		t.Errorf("Expected: %v, Got: %v", expectedResponse.GetMessage(), response.GetMessage())
	}
}
