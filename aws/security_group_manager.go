// security_group_manager.go
// This file manages the creation and configuration of EC2 security groups.
// Security groups control inbound and outbound traffic, allowing gRPC communication between nodes.
package aws

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// Create a new security group
func CreateSecurityGroup(svc *ec2.EC2, groupName, vpcId string) (*ec2.CreateSecurityGroupOutput, error) {
	input := &ec2.CreateSecurityGroupInput{
		GroupName:   aws.String(groupName),
		Description: aws.String("Security group for gRPC MPI project"),
		VpcId:       aws.String(vpcId),
	}

	result, err := svc.CreateSecurityGroup(input)
	if err != nil {
		log.Fatalf("Failed to create security group: %v", err)
		return nil, err
	}

	log.Printf("Created security group with ID: %s", *result.GroupId)
	return result, nil
}

// Add ingress rule (e.g., allow SSH)
func AuthorizeSecurityGroupIngress(svc *ec2.EC2, groupId string) {
	// Add logic to allow traffic on specific ports
}

// Delete a security group
func DeleteSecurityGroup(svc *ec2.EC2, groupId string) {
	// Add logic for deletion
}
