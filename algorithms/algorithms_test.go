package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	result = 662
	val    = []int{1, 2, 43, 45, 34, 3, 534}
)

func TestSimpleArraySum(t *testing.T) {
	assert.Equal(t, SimpleArraySum(val), result, "Sum Array Failed")
}
