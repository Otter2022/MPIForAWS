// config.go
// This file handles the configuration for the MPI-like framework, including
// loading environment variables such as node rank and the total number of nodes.
package mpi

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	AWS struct {
		Region          string `yaml:"region"`
		AMI             string `yaml:"ami"`
		InstanceType    string `yaml:"instanceType"`
		InstanceCount   int32  `yaml:"instanceCount"`
		SecurityGroupId string `yaml:"securityGroupId"`
		SubnetId        string `yaml:"subnetId"`
		KeyName         string `yaml:"keyName"`
		BucketName      string `yaml:"bucketName"`
	} `yaml:"aws"`

	S3 struct {
		ConfigFileKey string `yaml:"configFileKey"`
		TempFilePath  string `yaml:"tempFilePath"`
	} `yaml:"s3"`

	Network struct {
		GRPCPort int `yaml:"grpcPort"`
	} `yaml:"network"`

	Logging struct {
		LogFilePath string `yaml:"logFilePath"`
		LogLevel    string `yaml:"logLevel"`
	} `yaml:"logging"`
}

// LoadConfig reads the YAML config file and unmarshals it into a Config struct
func LoadConfig(configFile string) (*Config, error) {
	data, err := os.ReadFile(configFile) // Updated to use os.ReadFile
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config file: %w", err)
	}

	return &config, nil
}
