// mpi.go
// This file implements MPI-like primitives using gRPC. Functions include
// initialization of the communication environment, sending and receiving messages,
// and finalizing the environment after tasks are completed.

package mpi

// MPI_Init initializes the gRPC communication between nodes.
func MPI_Init() error

// MPI_Send sends a message to another node using gRPC.
func MPI_Send(content string, nodeRank int) error

// MPI_Recv receives a message from another node using gRPC.
func MPI_Recv() (string, error)

// MPI_Finalize cleans up resources and finalizes the MPI-like environment.
func MPI_Finalize() error
