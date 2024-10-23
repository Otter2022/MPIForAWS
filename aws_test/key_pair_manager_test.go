package aws_test

import (
	"context"
	"testing"

	"github.com/Otter2022/MPIForAWS/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/stretchr/testify/assert"
)

// Mock EC2 client
type mockEC2Client struct{}

func (m *mockEC2Client) CreateKeyPair(ctx context.Context, params *ec2.CreateKeyPairInput, optFns ...func(*ec2.Options)) (*ec2.CreateKeyPairOutput, error) {
	return &ec2.CreateKeyPairOutput{
		KeyMaterial: aws.String("mock-private-key-material"),
	}, nil
}

func TestCreateKeyPair(t *testing.T) {
	// Mock EC2 client
	client := &mockEC2Client{}

	// Call the function to create key pair
	keyMaterial, err := aws.CreateKeyPair(client, "my-keypair")
	assert.NoError(t, err)
	assert.Equal(t, "mock-private-key-material", keyMaterial)
}
