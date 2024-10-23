// worker.go
// This file defines the logic for each worker node in the distributed MPI-like framework.
// Worker nodes can either send or receive tasks based on their rank.

package mpi

// StartWorker starts the worker node logic. Master nodes send messages, while other nodes receive them.
func StartWorker(config *Config)
