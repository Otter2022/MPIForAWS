// main.go is the entry point of the application. It initializes the MPI environment,
// launches the worker nodes, and orchestrates the execution of distributed tasks
// by calling the appropriate MPI primitives and communication functions.
package main

import (
	"fmt"
	"log"

	"github.com/Otter2022/MPIForAWS/internal/mpi"
)

func main() {
	fmt.Println("Initializing MPI-like environment...")

	// Load configuration
	config, err := mpi.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize communication (e.g., SQS queues)
	if err := mpi.MPI_Init(); err != nil {
		log.Fatalf("Failed to initialize MPI: %v", err)
	}

	// Start the worker logic
	mpi.StartWorker(config)

	// Finalize the environment
	if err := mpi.MPI_Finalize(); err != nil {
		log.Fatalf("Failed to finalize MPI: %v", err)
	}
}
