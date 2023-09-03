package server

import (
	"fmt"
)

// EchoService struct to provide echoing functionality
type EchoService struct {
}

// Echo method takes a message string as input and returns a transformed
// message containing the length followed by the original message.
func (e *EchoService) Echo(message string) string {
	return fmt.Sprintf("%d:%s", len(message), message)
}
