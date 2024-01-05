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
