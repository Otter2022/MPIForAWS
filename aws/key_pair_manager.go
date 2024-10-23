// key_pair_manager.go
// This file handles the creation of EC2 key pairs, which allow SSH access to EC2 instances.
// The private key material is returned to the caller for secure access.
package aws

import (
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

// Create a new key pair
func CreateKeyPair(svc *ec2.Client, keyName string) {
	result, err := svc.CreateKeyPair(&ec2.CreateKeyPairInput{
		KeyName: aws.String(keyName),
	})
	if err != nil {
		log.Fatalf("Failed to create key pair: %v", err)
	}
	log.Printf("Created key pair: %s", *result.KeyName)
}

// Delete a key pair
func DeleteKeyPair(svc *ec2.Client, keyName string) {
	// Add deletion logic here
}

// Describe a key pair
func DescribeKeyPair(svc *ec2.Client, keyName string) {
	// Add logic to fetch key pair details
}
