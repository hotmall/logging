package logging

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRLogger(t *testing.T) {
	assert := assert.New(t)
	l1 := Logger("root")
	assert.NotNil(l1)

	l2 := Logger("ROOT")
	assert.NotNil(l2)
}

func TestLogger(t *testing.T) {
	assert := assert.New(t)
	l1 := Logger("mylog")
	assert.NotNil(l1)

	l2 := Logger("MYLOG")
	assert.NotNil(l2)

	l3 := Logger("mylog1")
	assert.NotNil(l3)
}
