package config

import (
	"testing"
)

func TestNewConfig(t *testing.T) {
	var expectedResult = ":9090"
	result, err := NewConfig()
	if err != nil {
		panic(err)
	}
	if result.Hostport != expectedResult {
		t.Error("Expected", expectedResult, "got: ", result)
	}
}
