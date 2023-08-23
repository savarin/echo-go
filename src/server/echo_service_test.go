package server

import (
	"testing"
)

func TestEchoShouldReturnTheSameMessage(t *testing.T) {
	// Arrange: Initialize EchoService and define the expected message
	echoService := &EchoService{}
	message := "Hello, World!"

	// Act: Invoke the echo method with the input message
	result := echoService.Echo(message)

	// Assert: Verify that the result is equal to the expected message
	expectedResult := "13:" + message
	if result != expectedResult {
		t.Errorf("Expected %s, but got %s", expectedResult, result)
	}
}
