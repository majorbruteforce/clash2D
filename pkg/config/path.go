package config

import (
	"log"
	"os"
)

var RootDir string

func init() {

	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get working directory: %v", err)
	}
	RootDir = dir
}
