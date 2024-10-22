// worker.go contains the logic for each worker node in the MPI-like framework.
// It listens for incoming messages, processes tasks, and communicates results
// using MPI communication functions such as MPI_Send and MPI_Recv.
package mpi

import (
	"fmt"
)

// StartWorker starts the worker logic for a given node.
func StartWorker(config *Config) {
	fmt.Printf("Node %d started, ready to process tasks.\n", config.NodeRank)

	// Example: Receive a task from the master node
	task, err := ReceiveMessage(config.SQSQueue)
	if err != nil {
		fmt.Printf("Failed to receive task on node %d: %v\n", config.NodeRank, err)
		return
	}

	fmt.Printf("Node %d received task: %s\n", config.NodeRank, task)

	// Process the task
	result, err := ProcessNodeTask(task)
	if err != nil {
		fmt.Printf("Node %d failed to process task: %v\n", config.NodeRank, err)
		return
	}

	// Send the result back to the master node or next node
	fmt.Printf("Node %d processed task, result: %v\n", config.NodeRank, result)
	err = SendMessage(config.SQSQueue, result)
	if err != nil {
		fmt.Printf("Failed to send result from node %d: %v\n", config.NodeRank, err)
	}
}

// ProcessNodeTask processes the task assigned to the worker node.
func ProcessNodeTask(task string) (string, error) {
	// Example task processing logic (could be any distributed computation)
	fmt.Printf("Processing task: %s\n", task)
	return "Task result", nil
}
