// security_group_manager.go
// This file manages the creation and configuration of EC2 security groups.
// Security groups control inbound and outbound traffic, allowing gRPC communication between nodes.

package aws

// EC2Manager defines the interface for EC2 operations
type EC2Manager interface {
	LaunchEC2Instances(count int64) []string
	DescribeEC2Instances(instanceIds []string)
	TerminateEC2Instances(instanceIds []string)
}

// KeyPairManager defines the interface for key pair operations
type KeyPairManager interface {
	CreateKeyPair(keyName string)
	DeleteKeyPair(keyName string)
	DescribeKeyPair(keyName string)
}

// SecurityGroupManager defines the interface for security group operations
type SecurityGroupManager interface {
	CreateSecurityGroup(groupName, vpcId string) string
	AuthorizeSecurityGroupIngress(groupId string)
	DeleteSecurityGroup(groupId string)
}
