package p020

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	assert := assert.New(t)
	assert.False(isValid("]"))
	assert.True(isValid("[]"))
	assert.False(isValid("[(]"))
	assert.True(isValid(""))
	assert.True(isValid("[()]"))
	assert.True(isValid("[({})]"))
	assert.False(isValid("[(})]"))
}
