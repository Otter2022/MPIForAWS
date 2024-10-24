// security_group_manager.go
// This file manages the creation and configuration of EC2 security groups.
// Security groups control inbound and outbound traffic, allowing gRPC communication between nodes.
package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

// Create a new security group
func CreateSecurityGroup(svc *ec2.Client, groupName, vpcId string) (*ec2.CreateSecurityGroupOutput, error) {
	input := &ec2.CreateSecurityGroupInput{
		GroupName:   aws.String(groupName),
		Description: aws.String("Security group for gRPC MPI project"),
		VpcId:       aws.String(vpcId),
	}

	result, err := svc.CreateSecurityGroup(context.TODO(), input)
	if err != nil {
		log.Printf("Failed to create security group: %v", err)
		return nil, err
	}

	log.Printf("Created security group with ID: %s", *result.GroupId)
	return result, nil
}

// Add ingress rule (e.g., allow SSH)
func AuthorizeSecurityGroupIngress(svc *ec2.Client, groupId string) {
	input := &ec2.AuthorizeSecurityGroupIngressInput{
		GroupId: aws.String(groupId),
		IpPermissions: []types.IpPermission{
			{
				IpProtocol: aws.String("tcp"),
				FromPort:   aws.Int32(22), // Port 22 for SSH
				ToPort:     aws.Int32(22),
				IpRanges: []types.IpRange{
					{
						CidrIp:      aws.String("0.0.0.0/0"), // Allow from all IPs
						Description: aws.String("Allow SSH"),
					},
				},
			},
		},
	}

	// Call AuthorizeSecurityGroupIngress API
	_, err := svc.AuthorizeSecurityGroupIngress(context.TODO(), input)
	if err != nil {
		log.Printf("Failed to authorize ingress for security group %s: %v", groupId, err)
		return
	}

	log.Printf("Successfully added ingress rule to security group %s", groupId)
}

// Delete a security group
func DeleteSecurityGroup(svc *ec2.Client, groupId string) error {
	input := &ec2.DeleteSecurityGroupInput{
		GroupId: aws.String(groupId),
	}

	// Call DeleteSecurityGroup API
	_, err := svc.DeleteSecurityGroup(context.TODO(), input)
	if err != nil {
		return fmt.Errorf("failed to delete security group %s: %v", groupId, err)
	}

	log.Printf("Successfully deleted security group: %s", groupId)
	return nil
}
