package aws

// CreateClient interface with a method to create an AWS client
type CreateClient interface {
	CreateClient() (interface{}, error) // Returns an interface{} so it can be flexible for different services
}

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
