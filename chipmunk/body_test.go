package chipmunk

/*
Copyright © 2012 Serge Zirukin

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

import (
	"github.com/bmizerany/assert"
	"testing"
)

func Test_BodyNewFree(t *testing.T) {
	b := BodyNew(1.0, 1.0)
	assert.NotEqual(t, nil, b.c())
	b.Free()
}

func Test_BodyInSpace(t *testing.T) {
	s := SpaceNew()

	b := BodyNew(1.0, 1.0)
	assert.Equal(t, (*Space)(nil), b.Space())
	assert.T(t, !s.Contains(b))

	s.AddBody(b)
	assert.NotEqual(t, nil, b.Space())
	assert.T(t, s.Contains(b))

	s.RemoveBody(b)
	assert.Equal(t, (*Space)(nil), b.Space())
	assert.T(t, !s.Contains(b))

	b.Free()
	s.Free()
}

func Test_BodySetPositionFunc(t *testing.T) {
	s := SpaceNew()
	b := BodyNew(1.0, 1.0)

	s.AddBody(b)

	type data struct {
		called bool
		s      string
	}

	b.SetUserData(&data{false, "userdata"})
	b.SetPositionFunc(func(b Body, dt float64) {
		d := b.UserData().(*data)
		d.called = true

		assert.Equal(t, "userdata", d.s)
		assert.Equal(t, 0.123, dt)
	})

	s.Step(0.123)

	assert.Tf(t, b.UserData().(*data).called, "position func wasn't called")

	b.Free()
	s.Free()
}

func Test_BodySetVelocityFunc(t *testing.T) {
	s := SpaceNew()
	b := BodyNew(1.0, 1.0)

	s.AddBody(b)

	type data struct {
		called bool
		s      string
	}

	b.SetUserData(&data{false, "userdata"})
	b.SetVelocityFunc(func(b Body, gravity Vect, damping, dt float64) {
		d := b.UserData().(*data)
		d.called = true

		assert.Equal(t, "userdata", d.s)
		assert.Equal(t, 0.123, dt)
	})

	s.Step(0.123)

	assert.Tf(t, b.UserData().(*data).called, "velocity func wasn't called")

	b.Free()
	s.Free()
}
