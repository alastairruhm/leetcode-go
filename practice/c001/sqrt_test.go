package c001

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSqrt(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(8.0, Sqrt(64.0, 0.000001))
	assert.Equal(1.414, Sqrt(2, 0.001))
}
