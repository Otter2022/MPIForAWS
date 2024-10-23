// grpc_server.go
// This file implements the gRPC server logic, allowing nodes to receive messages
// sent by other nodes in the MPI-like framework.

package mpi

import (
	"context"

	pb "github.com/Otter2022/MPIForAWS/proto"
)

// MPIServer is the gRPC server for receiving messages in the MPI-like framework.
type MPIServer struct {
	// UnimplementedMPIServerServer provides methods that must be implemented for the gRPC server.
	pb.UnimplementedMPIServerServer
}

// SendMessage receives a message sent by another node.
func (s *MPIServer) SendMessage(ctx context.Context, in *pb.Message) (*pb.Empty, error)

// StartGRPCServer starts the gRPC server to listen for incoming messages.
func StartGRPCServer()
