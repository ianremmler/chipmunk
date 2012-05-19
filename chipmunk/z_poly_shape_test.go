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
  "testing"
)

func Test_PolyShapeNew(t *testing.T) {
  b := NewBody(1.0, 1.0)
  points := []Vect{VectNew(0.0, 0.0), VectNew(0.0, 1.0), VectNew(1.0, 1.0), VectNew(1.0, 0.0)}
  s := PolyShapeNew(b, points, Origin())

  if s.c() == nil {
    t.Fatal("nil circle shape")
  }

  s.Free()
  b.Free()
}

func Test_PolyShapeUnsafe(t *testing.T) {
  b := NewBody(1.0, 1.0)
  points := []Vect{VectNew(0.0, 0.0), VectNew(0.0, 1.0), VectNew(1.0, 1.0), VectNew(1.0, 0.0)}
  s := PolyShapeNew(b, points, Origin())

  points = []Vect{VectNew(1.0, 0.0), VectNew(0.0, 0.0), VectNew(0.0, 1.0), VectNew(1.0, 1.0)}
  s.SetVerts(points, VectNew(1.0, -4.0))
  testEq(t, s.Vert(0), VectNew(2.0, -4.0))
  testEq(t, s.Vert(1), VectNew(1.0, -4.0))
  testEq(t, s.Vert(2), VectNew(1.0, -3.0))
  testEq(t, s.Vert(3), VectNew(2.0, -3.0))

  s.Free()
  b.Free()
}
