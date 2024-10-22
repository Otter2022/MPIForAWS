// config.go is responsible for loading and managing the configuration required
// for setting up the MPI environment, such as AWS SQS queue URLs, node ranks,
// and other settings needed for communication between nodes.
package mpi

import (
	"os"
	"strconv"
)

type Config struct {
	NodeRank   int
	TotalNodes int
	SQSQueue   string
}

// LoadConfig loads the configuration from environment variables or other sources.
func LoadConfig() (*Config, error) {
	nodeRank, err := strconv.Atoi(os.Getenv("NODE_RANK"))
	if err != nil {
		return nil, err
	}

	totalNodes, err := strconv.Atoi(os.Getenv("TOTAL_NODES"))
	if err != nil {
		return nil, err
	}

	sqsQueue := os.Getenv("SQS_QUEUE_URL")

	return &Config{
		NodeRank:   nodeRank,
		TotalNodes: totalNodes,
		SQSQueue:   sqsQueue,
	}, nil
}

// GetNodeRank returns the rank of the current node.
func GetNodeRank(config *Config) int {
	return config.NodeRank
}

// GetTotalNodes returns the total number of nodes.
func GetTotalNodes(config *Config) int {
	return config.TotalNodes
}
