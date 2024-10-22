package aws

import (
	"testing"
)

// Test the LaunchEC2Instance function
func TestLaunchEC2Instance(t *testing.T) {
	instanceType := "t2.micro"
	instanceID, err := LaunchEC2Instance(instanceType)
	if err != nil {
		t.Errorf("LaunchEC2Instance(%s) failed with error: %v", instanceType, err)
	}

	if instanceID == "" {
		t.Errorf("Expected a valid instance ID, but got an empty string")
	}
}

// Test the TerminateEC2Instance function
func TestTerminateEC2Instance(t *testing.T) {
	instanceID := "i-0123456789abcdef0" // Example instance ID
	err := TerminateEC2Instance(instanceID)
	if err != nil {
		t.Errorf("TerminateEC2Instance(%s) failed with error: %v", instanceID, err)
	}
}
