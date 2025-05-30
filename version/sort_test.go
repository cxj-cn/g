// Copyright (c) 2019 voidint <voidint@126.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package version

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {

	vs := []*Version{
		MustNew("1.21.0"),
		MustNew("1.20.10"),
		MustNew("1.21.4"),
		MustNew("1.20rc1"),
		MustNew("1.21rc2"),
		MustNew("1.19.12"),
		MustNew("1.21rc4"),
		MustNew("1.20.1"),
	}

	sort.Sort(Collection(vs))

	assert.Equal(t, vs[0].name, "1.19.12")
	assert.Equal(t, vs[1].name, "1.20rc1")
	assert.Equal(t, vs[2].name, "1.20.1")
	assert.Equal(t, vs[3].name, "1.20.10")
	assert.Equal(t, vs[4].name, "1.21rc2")
	assert.Equal(t, vs[5].name, "1.21rc4")
	assert.Equal(t, vs[6].name, "1.21.0")
	assert.Equal(t, vs[7].name, "1.21.4")
}
