// main.go
// This file is the main entry point for the MPI-like framework that initializes AWS resources
// (EC2 instances, security groups, and key pairs) and establishes gRPC communication between nodes.
// It orchestrates the setup, execution, and cleanup of the distributed system.
package main

import (
	"log"
	"time"

	myaws "github.com/Otter2022/MPIForAWS/aws" // Replace with the correct package path to your aws functions
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func main() {
	// Step 1: Create the EC2 client using your EC2ClientCreator
	creator := myaws.EC2ClientCreator{}
	ec2Client, err := creator.CreateClient()
	if err != nil {
		log.Fatalf("Failed to create EC2 client: %v", err)
	}

	// Step 2: Use a predefined VPC ID (replace with your actual VPC ID)
	vpcId := "vpc-xxxxxx x" // Replace with your actual VPC ID

	// Step 3: Create a security group for the EC2 instances
	securityGroupName := "mpi-security-group"

	// Use the VPC ID directly
	securityGroup, err := myaws.CreateSecurityGroup(ec2Client, securityGroupName, vpcId)
	if err != nil {
		log.Fatalf("Failed to create security group: %v", err)
	}
	groupId := *securityGroup.GroupId

	// Step 4: Authorize SSH and other necessary ports in the security group
	myaws.AuthorizeSecurityGroupIngress(ec2Client, groupId)

	// Step 5: Create a Key Pair
	keyName := "my-key-pair" // Replace with your desired key pair name
	myaws.CreateKeyPair(ec2Client, keyName)

	defer func() {
		// Ensure the key pair is deleted after the program ends
		err = myaws.DeleteKeyPair(ec2Client, keyName)
		if err != nil {
			log.Fatalf("Failed to delete key pair: %v", err)
		}
		log.Printf("Deleted key pair: %s", keyName)
	}()

	// Step 6: Launch EC2 instances
	instanceCount := int32(2)      // Number of instances to launch
	ami := "ami-0dba2cb6798deb6d8" // Example AMI, replace with yours
	instanceType := types.InstanceTypeT2Micro

	instanceIds := myaws.LaunchEC2Instances(ec2Client, instanceCount, ami, keyName, instanceType)
	log.Printf("Launched EC2 instances: %v", instanceIds)

	// Step 7: Wait for a few seconds to let the instances fully start
	time.Sleep(30 * time.Second)

	// Step 8: Describe the EC2 instances
	myaws.DescribeEC2Instances(ec2Client, instanceIds)

	// Step 9: Do any additional framework-related tasks (e.g., gRPC communication)

	// Step 10: Terminate EC2 instances after testing
	myaws.TerminateEC2Instances(ec2Client, instanceIds)
	log.Printf("Terminated EC2 instances: %v", instanceIds)

	// Step 11: Clean up - delete the security group
	myaws.DeleteSecurityGroup(ec2Client, groupId)
	log.Printf("Deleted security group: %s", groupId)

	// Optional: You can delete the VPC if you want, but since it's a predefined VPC, you might not want to delete it.
}
