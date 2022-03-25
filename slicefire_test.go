// Copyright 2022 helloshaohua <wu.shaohua@foxmail.com>;
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

package slicefire

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestSliceFireSuite(t *testing.T) {
	suite.Run(t, &SliceFireSuite{})
}

type SliceFireSuite struct {
	suite.Suite
}

func (suite *SliceFireSuite) Test_InStringSlice() {
	grids := []struct {
		haystack []string
		needle   string
		exists   bool
	}{
		{
			haystack: []string{"hello", "world"},
			needle:   "hello",
			exists:   true,
		},
		{
			haystack: []string{"hello", "world"},
			needle:   "hello-world",
			exists:   false,
		},
	}

	for _, grid := range grids {
		actual := InStringSlice(grid.haystack, grid.needle)
		assert.Equal(suite.T(), grid.exists, actual)
	}
}

func (suite *SliceFireSuite) Test_InSlice() {
	grids := []struct {
		haystack []interface{}
		needle   interface{}
		exists   bool
	}{
		{
			haystack: []interface{}{"hello", "world"},
			needle:   "hello",
			exists:   true,
		},
		{
			haystack: []interface{}{"hello", "world"},
			needle:   "hello-world",
			exists:   false,
		},
	}

	for _, grid := range grids {
		actual := InSlice(grid.haystack, grid.needle)
		assert.Equal(suite.T(), grid.exists, actual)
	}
}
