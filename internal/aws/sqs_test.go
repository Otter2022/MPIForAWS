package aws

import (
	"testing"
)

// Test the CreateSQSQueue function
func TestCreateSQSQueue(t *testing.T) {
	queueName := "test-queue"
	err := CreateSQSQueue(queueName)
	if err != nil {
		t.Errorf("CreateSQSQueue(%s) failed with error: %v", queueName, err)
	}
}

// Test the SendMessage function
func TestSendMessage(t *testing.T) {
	queueURL := "http://localhost:4566/000000000000/test-queue" // Example localstack URL
	message := "Hello, world!"

	err := SendMessage(queueURL, message)
	if err != nil {
		t.Errorf("SendMessage() to %s failed with error: %v", queueURL, err)
	}
}

// Test the ReceiveMessage function
func TestReceiveMessage(t *testing.T) {
	queueURL := "http://localhost:4566/000000000000/test-queue" // Example localstack URL

	message, err := ReceiveMessage(queueURL)
	if err != nil {
		t.Errorf("ReceiveMessage() from %s failed with error: %v", queueURL, err)
	}

	if message == "" {
		t.Errorf("ReceiveMessage() returned an empty message, expected a valid message")
	}
}
