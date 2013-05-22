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

func Test_SpaceNewFree(t *testing.T) {
	s := SpaceNew()
	assert.NotEqual(t, nil, s)
	s.Free()
}

func Test_SpaceAddPostStepCallback(t *testing.T) {
	s := SpaceNew()
	b := BodyNew(1.0, 1.0)
	c := CircleShapeNew(b, 8.0, Vect{0.0, 0.0})

	s.AddBody(b)
	s.AddShape(c)
	result := false

	s.AddPostStepCallback(func(s2 *Space, key interface{}) {
		assert.Equal(t, b, key.(Body))
		assert.Equal(t, s, s2)
		s2.RemoveShape(c)
		s2.RemoveBody(b)
		c.Free()
		b.Free()
		result = true
	}, b)

	s.Step(0.1)
	assert.T(t, result)
	result = false
	s.Step(0.1)
	assert.T(t, !result)
	s.Step(0.1)
	assert.T(t, !result)

	s.Free()
}

func Test_SpaceUserData(t *testing.T) {
	s := SpaceNew()
	x := "w00t"
	s.SetUserData(x)

	assert.Equal(t, "w00t", (s.UserData()).(string))

	s.Free()
}

func Test_SpaceContains(t *testing.T) {
	// func Test_SpaceAddShape(t *testing.T) {
	// func Test_SpaceAddBody(t *testing.T) {
	// func Test_SpaceAddConstraint(t *testing.T) {
	// func Test_SpaceRemoveShape(t *testing.T) {
	// func Test_SpaceRemoveBody(t *testing.T) {
	// func Test_SpaceRemoveConstraint(t *testing.T) {
	s := SpaceNew()

	b := BodyNew(1.0, 1.0)
	b2 := BodyNew(1.0, 1.0)
	c := CircleShapeNew(b, 8.0, Vect{0.0, 0.0})
	p := PinJointNew(b, b2, Origin(), Origin())

	assert.T(t, !s.Contains(b))
	assert.T(t, !s.Contains(b2))
	assert.T(t, !s.Contains(c))
	assert.T(t, !s.Contains(p))
	s.AddBody(b)
	s.AddBody(b2)
	s.AddShape(c)
	s.AddConstraint(p)
	assert.T(t, s.Contains(b))
	assert.T(t, s.Contains(b2))
	assert.T(t, s.Contains(c))
	assert.T(t, s.Contains(p))
	s.RemoveBody(b)
	s.RemoveBody(b2)
	s.RemoveShape(c)
	s.RemoveConstraint(p)
	assert.T(t, !s.Contains(b))
	assert.T(t, !s.Contains(b2))
	assert.T(t, !s.Contains(c))
	assert.T(t, !s.Contains(p))

	b.Free()
	b2.Free()
	c.Free()
	p.Free()

	s.Free()
}
