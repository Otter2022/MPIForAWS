package aws

import (
	"os"
	"testing"
)

var queueURL string

// Setup and teardown using TestMain for the whole test suite
func TestMain(m *testing.M) {
	// Setup - create the queue
	queueName := "test-queue"
	var err error
	queueURL, err = CreateSQSQueue(queueName)
	if err != nil {
		// If the queue creation fails, exit with a failure code
		panic("Failed to create SQS queue: " + err.Error())
	}

	// Run tests
	exitCode := m.Run()

	// Teardown - delete the queue after all tests have completed
	err = DeleteSQSQueue(queueURL)
	if err != nil {
		panic("Failed to delete SQS queue: " + err.Error())
	}

	// Exit the test process with the appropriate exit code
	os.Exit(exitCode)
}

// Test the CreateSQSQueue function
func TestCreateSQSQueue(t *testing.T) {
	if queueURL == "" {
		t.Errorf("Queue URL is empty, expected a valid queue URL")
	} else {
		t.Log("Queue created successfully:", queueURL)
	}
}

// Test the SendMessage function
func TestSendMessage(t *testing.T) {
	message := "Hello, world!"

	err := SendMessage(queueURL, message)
	if err != nil {
		t.Errorf("SendMessage() to %s failed with error: %v", queueURL, err)
	}
}

// Test the ReceiveMessage function
func TestReceiveMessage(t *testing.T) {
	message, err := ReceiveMessage(queueURL)
	if err != nil {
		t.Errorf("ReceiveMessage() from %s failed with error: %v", queueURL, err)
	}

	if message == "" {
		t.Errorf("ReceiveMessage() returned an empty message, expected a valid message")
	} else {
		t.Log("Message received:", message)
	}
}
