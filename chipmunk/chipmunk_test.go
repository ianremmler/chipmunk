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

func Test_UserData(t *testing.T) {
  s := NewSpace()
  x := "w00t"
  s.SetUserData(x)
  p := s.UserData()

  if p.(string) != "w00t" {
    t.Error("wrong user data")
  }

  s.Destroy()
}

func Test_Body(t *testing.T) {
  s := NewSpace()
  b := NewBody(1.0, 1.0)
  c := CircleShapeNew(b, 8.0, Vect{0.0, 0.0})

  s.AddBody(b)
  s.AddShape(c)

  if !s.Contains(c) {
    t.Error("space does not contain shape")
  }

  found := false

  b.ForEach(ShapeIterator(func(b Body, shape Shape) {
    found = (shape == c)
  }))

  if !found {
    t.Error("shape not found using body.ForEach")
  }

  found = false

  s.PointQuery(Vect{0.0, 0.0}, Layers(0xf), Group(0), PointQuery(func(shape Shape) {
    found = (shape == c)
  }))

  if !found {
    t.Error("shape not found using space.PointQuery")
  }

  found = true

  s.PointQuery(Vect{0.0, 0.0}, Layers(0), Group(0), PointQuery(func(shape Shape) {
    found = (shape == c)
  }))

  if !found {
    t.Error("shape found using space.PointQuery although Layers = 0")
  }

  s.RemoveShape(c)
  first := s.PointQueryFirst(Vect{0.0, 0.0}, Layers(0xf), Group(0))

  if nil != first {
    t.Error(first, "found using space.PointQueryFirst although it was removed")
  }

  x := PolyShapeNew(b, []Vect{Vect{1.0, 3.0}, Vect{6.0, -5.0}}, Vect{1.0, 2.0})
  t.Log(x)

  c.Destroy()
  b.Destroy()
  s.Destroy()
}
