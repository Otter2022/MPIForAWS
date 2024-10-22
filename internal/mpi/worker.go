// worker.go contains the logic for each worker node in the MPI-like framework.
// It listens for incoming messages, processes tasks, and communicates results
// using MPI communication functions such as MPI_Send and MPI_Recv.
package mpi

import (
	"fmt"
	"log"
)

// StartWorker starts the worker process, receiving tasks and processing them.
func StartWorker(config *Config) error {
	// Log the initialization of the worker node
	log.Printf("Worker node %d started\n", config.NodeRank)

	// Simulate receiving and processing tasks in a loop (in real use, you'd implement a termination condition)
	for {
		// Receive a message from another node (simulating a task)
		task, err := MPI_Recv(config.NodeRank - 1) // Simulate receiving from the previous node
		if err != nil {
			return fmt.Errorf("worker node %d failed to receive message: %w", config.NodeRank, err)
		}

		log.Printf("Worker node %d received task: %s\n", config.NodeRank, task)

		// Process the received task
		result, err := ProcessNodeTask(task)
		if err != nil {
			return fmt.Errorf("worker node %d failed to process task: %w", config.NodeRank, err)
		}

		log.Printf("Worker node %d processed task, result: %s\n", config.NodeRank, result)

		// Simulate sending the result back to another node (e.g., master or next node)
		err = MPI_Send(result, config.NodeRank+1) // Send result to the next node
		if err != nil {
			return fmt.Errorf("worker node %d failed to send result: %w", config.NodeRank, err)
		}

		log.Printf("Worker node %d sent result to node %d\n", config.NodeRank, config.NodeRank+1)
	}

	return nil
}

// ProcessNodeTask processes the task assigned to the node and returns the result.
func ProcessNodeTask(task string) (string, error) {
	// Simulate task processing (customize this with your actual task logic)
	log.Printf("Processing task: %s\n", task)

	// Simulated result of the task
	result := fmt.Sprintf("Processed task: %s", task)

	// In a real implementation, you would include actual processing logic here
	// e.g., perform computation, process data, etc.

	return result, nil
}
