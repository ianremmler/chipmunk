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

func Test_SegmentShapeNew(t *testing.T) {
  b := NewBody(1.0, 1.0)
  s := SegmentShapeNew(b, VectNew(1.0, 2.0), Origin(), 2.0)

  if s.c() == nil {
    t.Fatal("nil circle shape")
  }

  s.Free()
  b.Free()
}

func Test_SegmentShapeUnsafe(t *testing.T) {
  b := NewBody(1.0, 1.0)
  s := SegmentShapeNew(b, VectNew(1.0, 2.0), Origin(), 5.0)

  e1, e2 := VectNew(1.0, 1.0), VectNew(2.0, 2.0)
  s.SetEndpoints(e1, e2)
  testEq(t, s.A(), e1)
  testEq(t, s.B(), e2)

  radius := 2.0
  s.SetRadius(radius)
  testEq(t, s.Radius(), radius)

  s.Free()
  b.Free()
}
