// mpi.go implements the core MPI-like primitives such as MPI_Init, MPI_Send, MPI_Recv,
// and MPI_Finalize. These functions allow distributed nodes to communicate by
// passing messages via AWS SQS and managing their rank and communication environment.
package mpi

import (
	"fmt"
	"log"

	myaws "github.com/Otter2022/MPIForAWS/internal/aws"
)

// MPI_Init initializes the communication environment for the MPI framework
func MPI_Init() error {
	// Load configuration (node rank, total nodes, and queue information)
	config, err := LoadConfig()
	if err != nil {
		return fmt.Errorf("MPI_Init: failed to load configuration: %w", err)
	}

	// Example: You might want to initialize other components here, like creating queues
	// or setting up worker nodes.
	log.Printf("Node %d initialized out of %d total nodes\n", config.NodeRank, config.TotalNodes)

	return nil
}

// MPI_Send sends a message to another node
func MPI_Send(message string, destinationRank int) error {
	// Load configuration to get queue URL for destination node
	config, err := LoadConfig()
	if err != nil {
		return fmt.Errorf("MPI_Send: failed to load configuration: %w", err)
	}

	// Construct destination queue URL (in a real system, you would map ranks to queue URLs)
	queueURL := fmt.Sprintf("%s-%d", config.SQSQueue, destinationRank)

	// Send the message using the SendMessage function
	err = SendMessage(queueURL, message)
	if err != nil {
		return fmt.Errorf("MPI_Send: failed to send message to node %d: %w", destinationRank, err)
	}

	log.Printf("Message sent from node %d to node %d: %s\n", config.NodeRank, destinationRank, message)
	return nil
}

// MPI_Recv receives a message from another node
func MPI_Recv(sourceRank int) (string, error) {
	// Load configuration to get queue URL for this node
	config, err := LoadConfig()
	if err != nil {
		return "", fmt.Errorf("MPI_Recv: failed to load configuration: %w", err)
	}

	// Construct source queue URL (in a real system, you would map ranks to queue URLs)
	queueURL := fmt.Sprintf("%s-%d", config.SQSQueue, sourceRank)

	// Receive the message using the ReceiveMessage function
	message, err := ReceiveMessage(queueURL)
	if err != nil {
		return "", fmt.Errorf("MPI_Recv: failed to receive message from node %d: %w", sourceRank, err)
	}

	log.Printf("Message received by node %d from node %d: %s\n", config.NodeRank, sourceRank, message)
	return message, nil
}

// MPI_Comm_rank returns the rank of the current node
func MPI_Comm_rank() int {
	// Return node's rank
	return 0
}

// MPI_Comm_size returns the total number of nodes in the system
func MPI_Comm_size() int {
	// Return the total number of nodes
	return 0
}

// MPI_Finalize finalizes the MPI environment, performing any necessary cleanup
func MPI_Finalize() error {
	// Load configuration to get queue URL
	config, err := LoadConfig()
	if err != nil {
		return fmt.Errorf("MPI_Finalize: failed to load configuration: %w", err)
	}

	// Construct the queue URL for this node (or other nodes if necessary)
	queueURL := fmt.Sprintf("%s-%d", config.SQSQueue, config.NodeRank)

	// Delete the SQS queue associated with this node
	err = myaws.DeleteSQSQueue(queueURL)
	if err != nil {
		return fmt.Errorf("MPI_Finalize: failed to delete SQS queue %s: %w", queueURL, err)
	}

	log.Printf("MPI environment finalized and SQS queue %s deleted\n", queueURL)
	return nil
}
