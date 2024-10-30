// mpi.go
// This file implements MPI-like primitives using gRPC. Functions include
// initialization of the communication environment, sending and receiving messages,
// and finalizing the environment after tasks are completed.
package mpi

import (
	"fmt"
	"log"

	myaws "github.com/Otter2022/MPIForAWS/aws"
)

// MPI_Init initializes the gRPC server for MPI-like communication on a fixed port (50051).
// This setup allows the server to send and receive messages on the same port.
func MPI_Init(ip, bucketName string) error {
	// Initialize S3 client
	s3Client, err := myaws.NewS3Client(bucketName)
	if err != nil {
		return fmt.Errorf("failed to initialize S3 client: %v", err)
	}

	// Example: Download shared configuration file
	err = s3Client.DownloadFile("config.yaml", "/tmp/config.yaml")
	if err != nil {
		log.Printf("Failed to download config file: %v", err)
	} else {
		log.Println("Successfully downloaded config file from S3")
	}

	// Continue with the setup of the gRPC server
	port := "50051"
	address := ip + ":" + port

	go func() {
		if err := StartGRPCServer(address); err != nil {
			log.Printf("Failed to start gRPC server on %s: %v", address, err)
			return
		}
	}()

	log.Printf("MPI initialized and gRPC server started on %s", address)
	return nil
}

// MPI_Send sends a message to another node using gRPC
// It connects to the target node's gRPC server on the default port 50051.
func MPI_Send(targetIP, content string, nodeRank int, bucketName string) error {
	targetAddress := targetIP + ":50051"
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
