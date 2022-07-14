package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	resultTestSimpleArraySum = 662
	valTestSimpleArraySum    = []int{1, 2, 43, 45, 34, 3, 534}
)

func TestSimpleArraySum(t *testing.T) {
	assert.Equal(t, SimpleArraySum(valTestSimpleArraySum), resultTestSimpleArraySum, "Sum Array Failed")
}

var (
	resultTestaVeryBigSum int64 = 5000000015
	valueTestaVeryBigSum  = []int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}
)

func TestAVeryBigSum(t *testing.T) {
	assert.Equal(t, AVeryBigSum(valueTestaVeryBigSum), resultTestaVeryBigSum, "Very Big Sum Failed")

}
