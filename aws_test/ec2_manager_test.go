package aws_test

import (
	"context"
	"testing"

	"github.com/Otter2022/MPIForAWS/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/stretchr/testify/assert"
)

// Mock EC2 client
type mockEC2Client struct{}

func (m *mockEC2Client) RunInstances(ctx context.Context, params *ec2.RunInstancesInput, optFns ...func(*ec2.Options)) (*ec2.RunInstancesOutput, error) {
	return &ec2.RunInstancesOutput{
		Instances: []types.Instance{
			{InstanceId: aws.String("i-0123456789abcdef")},
		},
	}, nil
}

func (m *mockEC2Client) TerminateInstances(ctx context.Context, params *ec2.TerminateInstancesInput, optFns ...func(*ec2.Options)) (*ec2.TerminateInstancesOutput, error) {
	return &ec2.TerminateInstancesOutput{}, nil
}

func TestLaunchEC2Instances(t *testing.T) {
	// Mock EC2 client
	client := &mockEC2Client{}

	// Call the function to launch EC2 instances
	instanceIDs, err := aws.LaunchEC2Instances(client, 2, "sg-123456", "my-keypair")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(instanceIDs))
	assert.Equal(t, "i-0123456789abcdef", instanceIDs[0])
}

func TestTerminateEC2Instances(t *testing.T) {
	// Mock EC2 client
	client := &mockEC2Client{}

	// Call the function to terminate EC2 instances
	err := aws.TerminateEC2Instances(client, []string{"i-0123456789abcdef"})
	assert.NoError(t, err)
}
