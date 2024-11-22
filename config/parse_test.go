package config

import (
	"testing"
)

func TestParseConfig(t *testing.T) {
	cfg := Config{}
	cfg.ParseConfig("../config.yml")
	const expectedPort string = ":8080"
	if cfg.Port != expectedPort {
		t.Errorf("Expected %v. Got %v", expectedPort, cfg.Port)
	}

	const expectedCSVFileLocation string = "./matrix.csv"
	if cfg.CSVFileLocation != expectedCSVFileLocation {
		t.Errorf("Expected %v. Got %v", expectedCSVFileLocation, cfg.CSVFileLocation)
	}
}
