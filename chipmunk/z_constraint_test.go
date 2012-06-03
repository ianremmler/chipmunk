package chipmunk

/*
Copyright (c) 2012 Serge Zirukin

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

func Test_ConstraintNew(t *testing.T) {
  s := SpaceNew()

  b1 := s.AddBody(BodyNew(1.0, 1.0))
  s1 := s.AddShape(SegmentShapeNew(b1, VectNew(-10, 10), VectNew(10, 10), 0.0))

  b2 := s.AddBody(BodyNew(1.0, 1.0))
  s2 := s.AddShape(SegmentShapeNew(b2, VectNew(-10, 10), VectNew(10, 10), 0.0))

  c := GrooveJointNew(b1, b2, VectNew(0.0, 0.0), VectNew(0.0, 0.0), Origin())
  s.AddConstraint(c)

  s.Step(0.1)

  b1.Free()
  b2.Free()
  s1.Free()
  s2.Free()
  s.Free()
}

func Test_ConstraintSetPostSolveFunc(t *testing.T) {
  s := SpaceNew()

  b1 := s.AddBody(BodyNew(1.0, 1.0))
  s1 := s.AddShape(SegmentShapeNew(b1, VectNew(-10, 10), VectNew(10, 10), 0.0))

  b2 := s.AddBody(BodyNew(1.0, 1.0))
  s2 := s.AddShape(SegmentShapeNew(b2, VectNew(-10, 10), VectNew(10, 10), 0.0))

  c := GrooveJointNew(b1, b2, VectNew(0.0, 0.0), VectNew(0.0, 0.0), Origin())
  s.AddConstraint(c)
  c.SetUserData(13)
  called := false

  c.SetPostSolveFunc(func(c Constraint, s Space) {
    assert.Equal(t, 13, c.UserData())
    called = true
  })

  s.Step(0.1)

  assert.Tf(t, called, "postsolve func wasn't called")

  b1.Free()
  b2.Free()
  s1.Free()
  s2.Free()
  s.Free()
}

func Test_ConstraintSetPreSolveFunc(t *testing.T) {
  s := SpaceNew()

  b1 := s.AddBody(BodyNew(1.0, 1.0))
  s1 := s.AddShape(SegmentShapeNew(b1, VectNew(-10, 10), VectNew(10, 10), 0.0))

  b2 := s.AddBody(BodyNew(1.0, 1.0))
  s2 := s.AddShape(SegmentShapeNew(b2, VectNew(-10, 10), VectNew(10, 10), 0.0))

  c := GrooveJointNew(b1, b2, VectNew(0.0, 0.0), VectNew(0.0, 0.0), Origin())
  s.AddConstraint(c)
  c.SetUserData(13)
  called := false

  c.SetPreSolveFunc(func(c Constraint, s Space) {
    assert.Equal(t, 13, c.UserData())
    called = true
  })

  s.Step(0.1)

  assert.Tf(t, called, "presolve func wasn't called")

  b1.Free()
  b2.Free()
  s1.Free()
  s2.Free()
  s.Free()
}

// Local Variables:
// indent-tabs-mode: nil
// tab-width: 2
// End:
// ex: set tabstop=2 shiftwidth=2 expandtab:
