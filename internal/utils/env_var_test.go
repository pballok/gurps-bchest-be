package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnvOrFail_VariableDefined(t *testing.T) {
	expectedValue := "value"
	_ = os.Setenv("MY_VAR", expectedValue)
	defer func() { _ = os.Unsetenv("MY_VAR") }()

	value := GetEnvOrFail("MY_VAR")
	assert.Equal(t, expectedValue, value)
}
