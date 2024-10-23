// main.go
// This file is the main entry point for the MPI-like framework that initializes AWS resources
// (EC2 instances, security groups, and key pairs) and establishes gRPC communication between nodes.
// It orchestrates the setup, execution, and cleanup of the distributed system.

package main

import (
	"github.com/Otter2022/MPIForAWS/aws"
)

func main() {
	// Step 1: Create the key pair
	CreateKeyPair(ec2Client, keyName)

	// Step 2: Launch EC2 instances with the created key pair
	ami := "ami-0ea3c35c5c3284d82" // Example AMI ID
	instanceType := types.InstanceTypeT2Micro

	instanceIds := LaunchEC2Instances(ec2Client, 2, ami, keyName, instanceType)
}
