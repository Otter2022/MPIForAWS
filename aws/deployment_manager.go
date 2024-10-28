// deployment_manager.go
// deployment_manager.go
package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	ec2Types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

// getInstanceIPs fetches IP addresses of all instances in the specified subnet
func getInstanceIPs(client *ec2.Client, subnetID string) ([]string, error) {
	input := &ec2.DescribeInstancesInput{
		Filters: []ec2Types.Filter{
			{
				Name:   aws.String("subnet-id"),
				Values: []string{subnetID},
			},
		},
	}

	var ips []string
	paginator := ec2.NewDescribeInstancesPaginator(client, input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(context.TODO())
		if err != nil {
			return nil, fmt.Errorf("failed to get page of instances: %v", err)
		}
		for _, reservation := range page.Reservations {
			for _, instance := range reservation.Instances {
				if instance.PrivateIpAddress != nil {
					ips = append(ips, *instance.PrivateIpAddress)
				}
			}
		}
	}
	return ips, nil
}
