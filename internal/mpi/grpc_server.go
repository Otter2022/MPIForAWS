// grpc_server.go
// This file implements the gRPC server logic, allowing nodes to receive messages
// sent by other nodes in the MPI-like framework.
package mpi

import (
	"fmt"
	"log"
)

// MPI_Init initializes the gRPC server and retrieves necessary files from S3
func MPI_Init(ip, bucketName string) error {
	// Initialize S3 client
	s3Client, err := NewS3Client(bucketName)
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
