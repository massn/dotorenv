package envordot

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testDotEnvFile string = ".test.env"
const testParam = "ENVORDOT_TEST_PARAM"
const testDotEnvVal = "test_value_in_dot_env"
const testEnvVal = "test_value_in_env"

func TestDotEnvFirstRead(t *testing.T) {
	setupTest()
	defer teardownTest()
	m, err := Read(true, testDotEnvFile)
	assert.Nil(t, err)
	assert.Equal(t, testDotEnvVal, m[testParam])
}

func TestEnvFirstRead(t *testing.T) {
	setupTest()
	defer teardownTest()
	m, err := Read(false, testDotEnvFile)
	assert.Nil(t, err)
	assert.Equal(t, testEnvVal, m[testParam])
}

func TestDotEnvFirstLoad(t *testing.T) {
	setupTest()
	defer teardownTest()
	err := Load(true, testDotEnvFile)
	assert.Nil(t, err)
	assert.Equal(t, testDotEnvVal, os.Getenv(testParam))
}

func TestEnvFirstLoad(t *testing.T) {
	setupTest()
	defer teardownTest()
	err := Load(false, testDotEnvFile)
	assert.Nil(t, err)
	assert.Equal(t, testEnvVal, os.Getenv(testParam))
}

func setupTest() {
	os.Setenv(testParam, testEnvVal)
}

func teardownTest() {
	os.Unsetenv(testParam)
}
