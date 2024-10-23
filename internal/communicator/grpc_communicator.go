// grpc_communicator.go
// This file handles the gRPC communication between nodes in the MPI-like framework.
// It implements the logic to send and receive messages using gRPC clients and servers.
package communicator

import (
	"context"
	"log"
	"time"

	pb "github.com/Otter2022/MPIForAWS/internal/mpi" // Replace with the correct path to your generated protobuf package
	"google.golang.org/grpc"
)

type Communicator struct {
	client pb.MPIServerClient
}

// NewCommunicator establishes a connection to another node's gRPC server using grpc.DialContext
func NewCommunicator(address string) (*Communicator, error) {
	// Create a context with timeout for dialing
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Connect to the gRPC server using DialContext
	conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
		return nil, err
	}

	// Create a new gRPC client
	client := pb.NewMPIServerClient(conn)
	return &Communicator{client: client}, nil
}

// SendMessage sends a message to the gRPC server on another node
func (c *Communicator) SendMessage(content string, nodeRank int) error {
	// Set up a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Call the SendMessage function on the remote gRPC server
	_, err := c.client.SendMessage(ctx, &pb.Message{
		Content:  content,
		NodeRank: int32(nodeRank),
	})
	if err != nil {
		log.Printf("Failed to send message: %v", err)
		return err
	}

	log.Printf("Sent message: %s to node %d", content, nodeRank)
	return nil
}
