package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	valReverseArray    = []int{1, 2, 3, 4, 5}
	resultReverseArray = []int{5, 4, 3, 2, 1}
)

func TestReverseArray(t *testing.T) {
	assert.Equal(t, ReverseArray(valReverseArray), resultReverseArray, "Failed test Reverse Array")
}
