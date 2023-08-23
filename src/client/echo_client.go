package main

import (
    "context"
    "fmt"
    "google.golang.org/grpc"
    "os"

    proto "github.com/savarin/echo-go/proto"
)

type EchoClient struct {
    client proto.EchoServiceClient
}

func NewEchoClient(conn *grpc.ClientConn) *EchoClient {
    return &EchoClient{
        client: proto.NewEchoServiceClient(conn),
    }
}

func (c *EchoClient) Echo(message string) (string, error) {
    request := &proto.EchoRequest{
        Message: message,
    }
    response, err := c.client.Echo(context.Background(), request)
    if err != nil {
        return "", err
    }
    return response.GetMessage(), nil
}

func main() {
    conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
    if err != nil {
        fmt.Printf("Failed to connect: %v\n", err)
        return
    }
    defer conn.Close()

    client := NewEchoClient(conn)

    // Retrieve the message from the command-line arguments or use a default message
    message := "default-message"
    if len(os.Args) > 1 {
        message = os.Args[1]
    }

    // Send the message to the server and print the response
    response, err := client.Echo(message)
    if err != nil {
        fmt.Printf("Error calling echo: %v\n", err)
        return
    }
    fmt.Printf("Server responded with: %s\n", response)
}
