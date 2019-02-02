package helpers

import (
	"testing"
)

func TestAdd(t *testing.T) {
	var expectedResult = 12
	var result = Add(5, 7)
	if result != expectedResult {
		t.Error("Expected", expectedResult, "got: ", result)
	}
}
