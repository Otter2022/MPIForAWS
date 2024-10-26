// grpc_server.go
// This file implements the gRPC server logic, allowing nodes to receive messages
// sent by other nodes in the MPI-like framework.
package mpi

import (
	"log"

	"google.golang.org/grpc"
)

// MPIServer is the gRPC server that handles incoming messages
type MPIServer struct {
	UnimplementedMPIServerServer
}

// SendMessage is called whenever another node sends a message
// func (s *MPIServer) SendMessage(ctx context.Context, req *Message) (*Empty, error) {
// 	log.Printf("Received message from node %d: %s", req.NodeRank, req.Content)
// 	return &Empty{}, nil
// }

// StartGRPCServer starts the gRPC server to listen for incoming messages on the specified address
func StartGRPCServer(address string) error {
	// Create a TCP listener on the given address (e.g., ":50051")
	// lis, err := net.Listen("tcp", ":50051")
	// if err != nil {
	// 	log.Printf("Failed to listen: %v", err)
	// 	return err
	// }

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Register the gRPC service (MPIServer)
	RegisterMPIServerServer(grpcServer, &MPIServer{})

	log.Printf("Starting gRPC server on %s", address)

	// // Start serving incoming connections
	// if err := grpcServer.Serve(lis); err != nil {
	// 	log.Printf("Failed to serve: %v", err)
	// 	return err
	// }

	return nil
}
