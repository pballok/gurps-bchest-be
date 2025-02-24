package utils

import (
	"log"
	"os"
)

// GetEnvOrFail returns the value of the env variable, or panics if the variable doesn't exist
func GetEnvOrFail(varName string) string {
	value := os.Getenv(varName)
	if value == "" { // coverage-ignore
		log.Fatalf("Environment variable %s must be set", varName)
	}
	return value
}
