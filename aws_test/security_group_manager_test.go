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

func (m *mockEC2Client) CreateSecurityGroup(ctx context.Context, params *ec2.CreateSecurityGroupInput, optFns ...func(*ec2.Options)) (*ec2.CreateSecurityGroupOutput, error) {
	return &ec2.CreateSecurityGroupOutput{
		GroupId: aws.String("sg-0123456789abcdef"),
	}, nil
}

func (m *mockEC2Client) AuthorizeSecurityGroupIngress(ctx context.Context, params *ec2.AuthorizeSecurityGroupIngressInput, optFns ...func(*ec2.Options)) (*ec2.AuthorizeSecurityGroupIngressOutput, error) {
	return &ec2.AuthorizeSecurityGroupIngressOutput{}, nil
}

func TestCreateSecurityGroup(t *testing.T) {
	// Mock EC2 client
	client := &mockEC2Client{}

	// Call the function to create security group
	sgID, err := aws.CreateSecurityGroup(client, "grpc-security-group", "vpc-123456")
	assert.NoError(t, err)
	assert.Equal(t, "sg-0123456789abcdef", sgID)
}
