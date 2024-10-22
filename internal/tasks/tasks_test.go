package tasks

import (
	"reflect"
	"testing"
)

// TestDistributeTask tests that tasks are correctly distributed to the nodes
func TestDistributeTask(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}

	// Call the function to distribute tasks
	err := DistributeTask(data)

	// Check if an error occurred
	if err != nil {
		t.Errorf("DistributeTask() failed with error: %v", err)
	}
}

// TestProcessTask tests that the task processing works correctly
func TestProcessTask(t *testing.T) {
	taskID := 1
	data := []int{5, 3, 1, 4, 2}

	// Call the function to process the task (e.g., sorting the data)
	result, err := ProcessTask(taskID, data)

	// Check if an error occurred
	if err != nil {
		t.Errorf("ProcessTask() failed with error: %v", err)
	}

	// Define the expected result (assuming ProcessTask sorts the data)
	expected := []int{1, 2, 3, 4, 5}

	// Check if the result matches the expected value
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ProcessTask() result = %v, want %v", result, expected)
	}
}

// TestCollectResults tests that results are correctly collected from the nodes
func TestCollectResults(t *testing.T) {
	// Call the function to collect results
	result, err := CollectResults()

	// Check if an error occurred
	if err != nil {
		t.Errorf("CollectResults() failed with error: %v", err)
	}

	// Define the expected result (example data)
	expected := []int{1, 2, 3}

	// Check if the result matches the expected value
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("CollectResults() result = %v, want %v", result, expected)
	}
}
