// deployment_manager.go
// deployment_manager.go
package aws

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	ec2Types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

// InstanceInfo holds the instance ID and IP address
type InstanceInfo struct {
	InstanceID   string
	PrivateIP    string
	InstanceRank int
}

type SSMClientCreator struct{}

// GetInstanceIPs fetches the instance IDs and IP addresses of all instances in the specified subnet
func GetInstanceIPandIDs(client *ec2.Client, subnetID string) ([]InstanceInfo, error) {
	input := &ec2.DescribeInstancesInput{
		Filters: []ec2Types.Filter{
			{
				Name:   aws.String("subnet-id"),
				Values: []string{subnetID},
			},
		},
	}

	var instances []InstanceInfo
	paginator := ec2.NewDescribeInstancesPaginator(client, input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(context.TODO())
		if err != nil {
			return nil, err
		}
		for _, reservation := range page.Reservations {
			for _, instance := range reservation.Instances {
				if instance.InstanceId != nil && instance.PrivateIpAddress != nil {
					instances = append(instances, InstanceInfo{
						InstanceID: *instance.InstanceId,
						PrivateIP:  *instance.PrivateIpAddress,
					})
				}
			}
		}
	}
	return instances, nil
}

func InitializeEnviromentsAndBuild(client *ssm.Client, instances []InstanceInfo) ([]InstanceInfo, error) {
	n := len(instances)

	for i, instance := range instances {
		var commands []string
		commands = append(commands, fmt.Sprintf("EXPORT MPI_SIZE=%d", n))
		commands = append(commands, fmt.Sprintf("EXPORT MPI_RANK=%d", i))
		for x := range n {
			if x == i {
				commands = append(commands, fmt.Sprintf("EXPORT MPI_ADDRESS_%d=\"0.0.0.0:50051\"", i))
				instances[i].InstanceRank = i
			} else {
				commands = append(commands, fmt.Sprintf("EXPORT MPI_ADDRESS_%d=\"%v:50051\"", i, instance.PrivateIP))
			}
		}
		commands = append(commands, "cd cloud-native-mpi-for-aws")
		commands = append(commands, "go build -o mpi_program")
		commands = append(commands, "./mpi_program > ../output.txt")
		input := &ssm.SendCommandInput{
			DocumentName: aws.String("Makeenviromentvariables"),
			Parameters: map[string][]string{
				"commands": commands,
			},
			InstanceIds: []string{instances[i].InstanceID},
		}
		result, err := client.SendCommand(context.TODO(), input)
		if err != nil {
			return nil, err
		} else {
			fmt.Printf("%v", result)
		}
	}

	return instances, nil
}

// CreateClient method creates the EC2 client using AWS SDK v2
func (s *SSMClientCreator) CreateClient() (*ssm.Client, error) {
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

	client := ssm.NewFromConfig(cfg)
	return client, err
}
