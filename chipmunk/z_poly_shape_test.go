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

func Test_PolyShapeNew(t *testing.T) {
  b := BodyNew(1.0, 1.0)
  points := []Vect{VectNew(0.0, 0.0), VectNew(0.0, 1.0), VectNew(1.0, 1.0), VectNew(1.0, 0.0)}
  s := PolyShapeNew(b, points, Origin())

  assert.NotEqual(t, nil, s.c())

  s.Free()
  b.Free()
}

func Test_PolyShapeUnsafe(t *testing.T) {
  b := BodyNew(1.0, 1.0)
  points := []Vect{VectNew(0.0, 0.0), VectNew(0.0, 1.0), VectNew(1.0, 1.0), VectNew(1.0, 0.0)}
  s := PolyShapeNew(b, points, Origin())

  points = []Vect{VectNew(1.0, 0.0), VectNew(0.0, 0.0), VectNew(0.0, 1.0), VectNew(1.0, 1.0)}
  s.SetVerts(points, VectNew(1.0, -4.0))
  assert.Equal(t, VectNew(2.0, -4.0), s.VertLocal(0))
  assert.Equal(t, VectNew(1.0, -4.0), s.VertLocal(1))
  assert.Equal(t, VectNew(1.0, -3.0), s.VertLocal(2))
  assert.Equal(t, VectNew(2.0, -3.0), s.VertLocal(3))

  s.Free()
  b.Free()
}

// Local Variables:
// indent-tabs-mode: nil
// tab-width: 2
// End:
// ex: set tabstop=2 shiftwidth=2 expandtab:
