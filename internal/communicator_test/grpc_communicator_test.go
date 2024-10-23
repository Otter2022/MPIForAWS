package communicator_test

import (
	"testing"
	"time"

	"context"
	"net"

	"github.com/Otter2022/MPIForAWS/internal/communicator"
	pb "github.com/Otter2022/MPIForAWS/proto"
	"google.golang.org/grpc"
)

// Mock gRPC server for testing
type mockMPIServer struct {
	pb.UnimplementedMPIServerServer
}

func (s *mockMPIServer) SendMessage(ctx context.Context, msg *pb.Message) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}

func TestSendMessage(t *testing.T) {
	// Start a local gRPC server for testing
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		t.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterMPIServerServer(grpcServer, &mockMPIServer{})
	go grpcServer.Serve(lis)
	defer grpcServer.Stop()

	time.Sleep(1 * time.Second) // Allow time for server to start

	// Create the gRPC client
	comm, err := communicator.NewCommunicator("localhost:50051")
	if err != nil {
		t.Fatalf("Failed to create gRPC client: %v", err)
	}

	// Send message using the client
	err = comm.SendMessage("Test message", 1)
	if err != nil {
		t.Fatalf("Failed to send message: %v", err)
	}
}
