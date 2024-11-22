package config

import (
	"os"

	l "github.com/ogonna-anaekwe/matrixserver/internal/logger"
	"gopkg.in/yaml.v3"
)

// Server configs.
type Config struct {
	Port            string `yaml:"port"`              // Port on which server listens
	CSVFileLocation string `yaml:"csv_file_location"` // Path to CSV file containing matrix data
}

// Reads config file.
func (c *Config) ParseConfig(configFile string) {
	log := l.NewLogger()
	log.SetReportCaller(true)

	file, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Error reading config file %v", configFile)
	}

	cfg := Config{}
	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		log.Fatalf("Error unmarshalling config file %v", configFile)
	}

	c.Port = cfg.Port
	c.CSVFileLocation = cfg.CSVFileLocation
}
