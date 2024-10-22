// sqs.go provides functions for interacting with AWS SQS, allowing nodes to send
// and receive messages. It abstracts the message-passing mechanisms required
// for node-to-node communication in the MPI-like framework.
package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

// CreateSQSClient creates an AWS SQS client using default credentials and configuration.
func CreateSQSQueue(queueName string) (string, error) {
	// Create SQS client
	client, err := CreateSQSClient()
	if err != nil {
		return "", fmt.Errorf("failed to create SQS client: %w", err)
	}

	// Create the queue
	input := &sqs.CreateQueueInput{
		QueueName: aws.String(queueName),
	}

	result, err := client.CreateQueue(context.TODO(), input)
	if err != nil {
		return "", fmt.Errorf("failed to create queue: %w", err)
	}

	// Return the queue URL
	return aws.ToString(result.QueueUrl), nil
}

// DeleteSQSQueue deletes the specified SQS queue by its URL.
func DeleteSQSQueue(queueURL string) error {
	// Load the default AWS configuration
	client, err := CreateSQSClient()
	if err != nil {
		return fmt.Errorf("failed to create SQS client: %w", err)
	}

	// Delete the SQS queue
	_, err = client.DeleteQueue(context.TODO(), &sqs.DeleteQueueInput{
		QueueUrl: aws.String(queueURL),
	})
	if err != nil {
		return fmt.Errorf("failed to delete SQS queue %s: %w", queueURL, err)
	}

	fmt.Printf("SQS queue %s deleted successfully\n", queueURL)
	return nil
}

// CreateSQSClient creates an AWS SQS client using the default credentials and configuration.
func CreateSQSClient() (*sqs.Client, error) {
	// Load the AWS configuration (credentials, region, etc.)
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		return nil, fmt.Errorf("unable to load AWS config: %w", err)
	}

	// Create and return an SQS client
	client := sqs.NewFromConfig(cfg)
	return client, nil
}
