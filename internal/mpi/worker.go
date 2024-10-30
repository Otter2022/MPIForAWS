// worker.go
// This file defines the logic for each worker node in the distributed MPI-like framework.
// Worker nodes can either send or receive tasks based on their rank.

package mpi

import (
	"fmt"
	"log"
)

// StartWorker starts the worker node logic. Master nodes send messages, while other nodes receive them.
func StartWorker(nodeRank int, totalNodes int) {
	if nodeRank == 0 {
		// Master node logic
		for i := 1; i < totalNodes; i++ {
			targetIP := fmt.Sprintf("node%d.internal:50051", i) // Replace with actual IP or hostname
			err := MPI_Send(targetIP, fmt.Sprintf("Hello from master to node %d", i), nodeRank)
			if err != nil {
				log.Printf("Failed to send message to node %d: %v", i, err)
			}
		}
	} else {
		// Worker node logic
		MPI_Recv()
	}
}
