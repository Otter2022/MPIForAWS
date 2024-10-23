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
	input := &ec2.DescribeInstancesInput{
		InstanceIds: instanceIds,
	}

	// Call DescribeInstances API
	result, err := svc.DescribeInstances(context.TODO(), input)
	if err != nil {
		log.Fatalf("Failed to describe EC2 instances: %v", err)
	}

	// Iterate over the instances and log details
	for _, reservation := range result.Reservations {
		for _, instance := range reservation.Instances {
			log.Printf("Instance ID: %s", *instance.InstanceId)
			log.Printf("Instance State: %s", instance.State.Name)
			log.Printf("Instance Type: %s", instance.InstanceType)
			if instance.PublicIpAddress != nil {
				log.Printf("Public IP: %s", *instance.PublicIpAddress)
			} else {
				log.Println("Public IP: None")
			}
		}
	}
}

// TerminateEC2Instances terminates EC2 instances
func TerminateEC2Instances(svc *ec2.Client, instanceIds []string) {
	input := &ec2.TerminateInstancesInput{
		InstanceIds: instanceIds, // No need to use aws.StringSlice here as it already expects []string
	}

	// Call the TerminateInstances API
	result, err := svc.TerminateInstances(context.TODO(), input)
	if err != nil {
		log.Fatalf("Failed to terminate EC2 instances: %v", err)
	}

	for _, instance := range result.TerminatingInstances {
		log.Printf("Terminating instance: %s, current state: %s, previous state: %s",
			*instance.InstanceId,
			instance.CurrentState.Name,
			instance.PreviousState.Name)
	}
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
