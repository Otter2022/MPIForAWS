// ec2_manager.go
// This file handles the creation, management, and termination of EC2 instances used
// for running the gRPC nodes in the MPI-like framework.
package aws

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

// EC2ClientCreator implements the CreateClient interface for EC2
type EC2ClientCreator struct{}

// LaunchEC2Instances launches a specified number of EC2 instances with a given AMI and instance type.
func LaunchEC2Instances(svc *ec2.Client, count int32, ami, keyName string, instanceType types.InstanceType) []string {
	runResult, err := svc.RunInstances(context.TODO(), &ec2.RunInstancesInput{
		ImageId:      aws.String(ami),  // AMI ID
		InstanceType: instanceType,     // Instance type (e.g., t2.micro)
		MinCount:     aws.Int32(count), // Number of instances
		MaxCount:     aws.Int32(count),
		KeyName:      aws.String(keyName), // The name of your key pair
	})
	if err != nil {
		log.Fatalf("Failed to create EC2 instances: %v", err)
	}

	instanceIds := []string{}
	for _, instance := range runResult.Instances {
		instanceIds = append(instanceIds, *instance.InstanceId)
	}

	log.Printf("Created instances: %v", instanceIds)
	return instanceIds
}

// DescribeEC2Instances describes running EC2 instances
func DescribeEC2Instances(svc *ec2.Client, instanceIds []string) {
	// Example of getting details of instances
	// Add implementation for describing instances here
}

// TerminateEC2Instances terminates EC2 instances
func TerminateEC2Instances(svc *ec2.Client, instanceIds []string) {
	// Add termination logic here
}

// CreateClient method creates the EC2 client using AWS SDK v2
func (s *EC2ClientCreator) CreateClient() (*ec2.Client, error) {
	var cfg aws.Config
	var err error

	region := os.Getenv("AWS_REGION")
	if region != "" {
		cfg, err = config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	} else {
		cfg, err = config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	}

	if err != nil {
		return nil, fmt.Errorf("unable to load AWS config: %w", err)
	}

	client := ec2.NewFromConfig(cfg)
	return client, err
}
