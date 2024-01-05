package logging

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnv(t *testing.T) {
	assert := assert.New(t)
	items := []struct {
		input  string
		output Env
		err    error
	}{
		{"prod", ProdEnv, nil},
		{"PROD", ProdEnv, nil},
		{"ProD", ProdEnv, nil},
		{"", ProdEnv, nil},
		{"dev", DevEnv, nil},
		{"DEV", DevEnv, nil},
		{"Dev", DevEnv, nil},
		{"xxx", ProdEnv, errors.New("unrecognized env: xxx")},
	}

	for _, item := range items {
		var e Env
		err := e.Set(item.input)
		assert.Equal(item.err, err)
		// fmt.Printf("input=%s, env: %d\n", item.input, e)
		assert.Equal(item.output, e)
	}
}
