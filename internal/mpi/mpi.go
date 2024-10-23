// mpi.go
// This file implements MPI-like primitives using gRPC. Functions include
// initialization of the communication environment, sending and receiving messages,
// and finalizing the environment after tasks are completed.

package mpi

import (
	"log"
)

// MPI_Init initializes the MPI environment and starts the gRPC server
func MPI_Init(address string) error {
	// Start the gRPC server to handle incoming messages
	go func() {
		if err := StartGRPCServer(address); err != nil {
			log.Fatalf("Failed to start gRPC server: %v", err)
		}
	}()
	log.Printf("MPI initialized and gRPC server started on %s", address)
	return nil
}

// MPI_Send sends a message to another node using gRPC
func MPI_Send(targetAddress string, content string, nodeRank int) error {
	comm, err := NewCommunicator(targetAddress)
	if err != nil {
		return err
	}

	// Send the message to the target node
	err = comm.SendMessage(content, nodeRank)
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
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
