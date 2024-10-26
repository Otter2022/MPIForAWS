// mpi.go
// This file implements MPI-like primitives using gRPC. Functions include
// initialization of the communication environment, sending and receiving messages,
// and finalizing the environment after tasks are completed.
package mpi

import (
	"log"
)

// MPI_Init initializes the gRPC server for MPI-like communication on a fixed port (50051).
// This setup allows the server to send and receive messages on the same port.
func MPI_Init(ip string) error {
	port := "50051"
	address := ip + ":" + port // Bind to the specific private IP and port

	// Start the gRPC server to handle incoming messages
	go func() {
		if err := StartGRPCServer(address); err != nil {
			log.Printf("Failed to start gRPC server on %s: %v", address, err)
			return
		}
	}()

	log.Printf("MPI initialized and gRPC server started on %s", address)
	return nil // Return from the main function
}

// MPI_Send sends a message to another node using gRPC
// It connects to the target node's gRPC server on the default port 50051.
func MPI_Send(targetIP string, content string, nodeRank int) error {
	targetAddress := targetIP + ":50051" // Use the fixed port 50051 for all nodes
	comm, err := NewCommunicator(targetAddress)
	if err != nil {
		return err
	}

	// Send the message to the target node
	err = comm.SendMessage(content, nodeRank)
	if err != nil {
		log.Printf("Failed to send message: %v", err)
		return err
	}
	return nil
}

// MPI_Recv waits for a message (this is handled by the gRPC server)
func MPI_Recv() {
	log.Println("Waiting to receive a message...")
	// This is handled by the gRPC server in grpc_server.go
}

// MPI_Finalize cleans up resources
func MPI_Finalize() {
	log.Println("Finalizing MPI environment...")
	// Add any necessary cleanup code here
}
