// sqs.go provides functions for interacting with AWS SQS, allowing nodes to send
// and receive messages. It abstracts the message-passing mechanisms required
// for node-to-node communication in the MPI-like framework.
package aws

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

// SQSClientCreator implements the CreateClient interface for SQS
type SQSClientCreator struct{}

// CreateSQSClient creates an AWS SQS client using default credentials and configuration.
func CreateSQSQueue(queueName string) (string, error) {
	var clientCreator CreateClient

	// Create SQS client
	clientCreator = &SQSClientCreator{}
	client, err := clientCreator.CreateClient()
	if err != nil {
		return "", fmt.Errorf("failed to create SQS client: %w", err)
	}

	// Type assert the client to *sqs.Client
	sqsClient, ok := client.(*sqs.Client)
	if !ok {
		log.Fatalf("failed to assert to *sqs.Client")
	}

	// Create the queue
	input := &sqs.CreateQueueInput{
		QueueName: aws.String(queueName),
	}

	result, err := sqsClient.CreateQueue(context.TODO(), input)
	if err != nil {
		return "", fmt.Errorf("failed to create queue: %w", err)
	}

	// Return the queue URL
	return aws.ToString(result.QueueUrl), nil
}

// DeleteSQSQueue deletes the specified SQS queue by its URL.
func DeleteSQSQueue(queueURL string) error {
	var clientCreator CreateClient

	// Load the default AWS configuration
	clientCreator = &SQSClientCreator{}
	client, err := clientCreator.CreateClient()
	if err != nil {
		return fmt.Errorf("failed to create SQS client: %w", err)
	}

	// Type assert the client to *sqs.Client
	sqsClient, ok := client.(*sqs.Client)
	if !ok {
		log.Fatalf("failed to assert to *sqs.Client")
	}

	// Delete the SQS queue
	_, err = sqsClient.DeleteQueue(context.TODO(), &sqs.DeleteQueueInput{
		QueueUrl: aws.String(queueURL),
	})
	if err != nil {
		return fmt.Errorf("failed to delete SQS queue %s: %w", queueURL, err)
	}

	fmt.Printf("SQS queue %s deleted successfully\n", queueURL)
	return nil
}

// CreateClient method for SQS which implements CreateClient interface
func (s *SQSClientCreator) CreateClient() (interface{}, error) {
	var cfg aws.Config
	var err error
	err = nil

	region := os.Getenv("AWS_REGION")
	if region != "" {
		cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
		if err != nil {
			return nil, fmt.Errorf("unable to load AWS config: %w", err)
		}
	} else {
		cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
		if err != nil {
			return nil, fmt.Errorf("unable to load AWS config: %w", err)
		}
	}

	client := sqs.NewFromConfig(cfg)
	return client, err
}
