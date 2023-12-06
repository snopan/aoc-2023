package helpers

import (
	"testing"

	"gotest.tools/assert"
)

func Test_CombineIntArrToInt(t *testing.T) {
	input := []int{1, 23, 4}
	ouptut := CombineIntArrToInt(input)
	assert.Equal(t, ouptut, 1234)
}
