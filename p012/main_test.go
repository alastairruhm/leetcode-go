package p012

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntToRoman(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("III", intToRoman(3))
	assert.Equal("IV", intToRoman(4))
	assert.Equal("IX", intToRoman(9))
	assert.Equal("LVIII", intToRoman(58))
	assert.Equal("MCMXCIV", intToRoman(1994))
	assert.Equal("LX", intToRoman(60))
}
