package gonotion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_loadEnvironmentVariables(t *testing.T) {
	// Set up the environment variables for the test
	t.Setenv("INTEGRATION_TOKEN", "test")
	v, e := loadEnvironmentVariables()

	assert.Equal(t, "test", v, "ENV_VAR returns the correct value")
	assert.Nil(t, e, "ENV_VAR returns no error")

	// TODO: Complete the rest of the tests
	// t.Setenv("INTEGRATION_TOKEN", "")
	// assert.Error(t, e, "ENV_VAR_FAIL returns an error")
	// assert.Nil(t, v, "ENV_VAR_FAIL returns no value")
}
