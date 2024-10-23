// grpc_communicator.go
// This file handles the gRPC communication between nodes in the MPI-like framework.
// It implements the logic to send and receive messages using gRPC clients and servers.

package communicator

// NewCommunicator creates a new gRPC client to communicate with other nodes.
func NewCommunicator(address string) (*Communicator, error)

// SendMessage sends a message to the gRPC server running on another node.
func (c *Communicator) SendMessage(content string, nodeRank int) error
