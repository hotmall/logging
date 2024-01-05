// Copyright © 2024 The Hot Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
