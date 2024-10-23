// communicator.go handles the low-level communication between nodes, abstracting
// message-passing using AWS SQS. It ensures that messages are sent and received
// correctly between distributed nodes in the cloud.
package mpi

import (
	"context"
	"fmt"
	"log"

	myaws "github.com/Otter2022/MPIForAWS/internal/aws"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

// SendMessage sends a message to the specified SQS queue
func SendMessage(queueURL string, message string) error {
	// Create SQS client
	clientCreator := &myaws.SQSClientCreator{}
	client, err := clientCreator.CreateClient()
	if err != nil {
		return fmt.Errorf("failed to create SQS client: %w", err)
	}

	// Type assert the client to *sqs.Client
	sqsClient, ok := client.(*sqs.Client)
	if !ok {
		log.Fatalf("failed to assert to *sqs.Client")
	}

	// Send the message
	input := &sqs.SendMessageInput{
		QueueUrl:    aws.String(queueURL),
		MessageBody: aws.String(message),
	}

	_, err = sqsClient.SendMessage(context.TODO(), input)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	log.Printf("Message sent to queue %s: %s\n", queueURL, message)
	return nil
}

// ReceiveMessage receives a message from the specified SQS queue
func ReceiveMessage(queueURL string) (string, error) {
	// Create SQS client
	clientCreator := &myaws.SQSClientCreator{}
	client, err := clientCreator.CreateClient()
	if err != nil {
		return "", fmt.Errorf("failed to create SQS client: %w", err)
	}

	// Type assert the client to *sqs.Client
	sqsClient, ok := client.(*sqs.Client)
	if !ok {
		log.Fatalf("failed to assert to *sqs.Client")
	}

	// Receive messages from the queue
	input := &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(queueURL),
		MaxNumberOfMessages: 1, // Limit to one message at a time
		WaitTimeSeconds:     5, // Long polling (wait for messages)
	}

	result, err := sqsClient.ReceiveMessage(context.TODO(), input)
	if err != nil {
		return "", fmt.Errorf("failed to receive message: %w", err)
	}

	// Check if any messages were received
	if len(result.Messages) == 0 {
		log.Printf("No messages received from queue: %s\n", queueURL)
		return "", nil
	}

	// Return the message body of the first message
	message := result.Messages[0].Body
	log.Printf("Message received from queue %s: %s\n", queueURL, *message)

	return *message, nil
}

// BroadcastMessage sends a message to all nodes.
func BroadcastMessage(queueURLs []string, message string) error {
	for _, queueURL := range queueURLs {
		err := SendMessage(queueURL, message)
		if err != nil {
			return fmt.Errorf("failed to broadcast message to %s: %w", queueURL, err)
		}
	}

	log.Printf("Message broadcasted to all nodes: %s\n", message)
	return nil
}
