// grpc_communicator.go
// This file handles the gRPC communication between nodes in the MPI-like framework.
// It implements the logic to send and receive messages using gRPC clients and servers.
package mpi

import (
	"context"
	"crypto/tls"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Communicator struct holds the gRPC client and the connection
type Communicator struct {
	client MPIServerClient
	conn   *grpc.ClientConn // Store the connection to close later
}

// NewCommunicator establishes a secure connection to another node's gRPC server
func NewCommunicator(address string) (*Communicator, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // Increased timeout for reliability
	defer cancel()

	// Load secure TLS credentials
	tlsCredentials := credentials.NewTLS(&tls.Config{
		InsecureSkipVerify: true, // This disables server certificate validation (insecure, use only for local testing)
	})

	// Connect to the gRPC server using secure TLS credentials
	conn, err := grpc.DialContext(ctx, address, grpc.WithTransportCredentials(tlsCredentials), grpc.WithBlock())
	if err != nil {
		log.Printf("Failed to connect: %v", err)
		return nil, err
	}

	// Create a new gRPC client
	client := NewMPIServerClient(conn)
	return &Communicator{client: client, conn: conn}, nil
}

// Close closes the gRPC connection
func (c *Communicator) Close() error {
	if err := c.conn.Close(); err != nil {
		log.Printf("Failed to close connection: %v", err)
		return err
	}
	return nil
}

// SendMessage sends a message to the gRPC server on another node
func (c *Communicator) SendMessage(content string, nodeRank int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // Increased timeout for sending
	defer cancel()

	// Call the SendMessage function on the remote gRPC server
	_, err := c.client.SendMessage(ctx, &Message{
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
