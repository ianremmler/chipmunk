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

func Test_SpaceNewFree(t *testing.T) {
  s := SpaceNew()

  if s == nil {
    t.Fatal("nil space")
  }

  s.Free()
}

func Test_SpaceUserData(t *testing.T) {
  s := SpaceNew()
  x := "w00t"
  s.SetUserData(x)

  testEq(t, (s.UserData()).(string), "w00t")

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

  testEq(t, s.Contains(b), false)
  testEq(t, s.Contains(b2), false)
  testEq(t, s.Contains(c), false)
  testEq(t, s.Contains(p), false)
  s.AddBody(b)
  s.AddBody(b2)
  s.AddShape(c)
  s.AddConstraint(p)
  testEq(t, s.Contains(b), true)
  testEq(t, s.Contains(b2), true)
  testEq(t, s.Contains(c), true)
  testEq(t, s.Contains(p), true)
  s.RemoveBody(b)
  s.RemoveBody(b2)
  s.RemoveShape(c)
  s.RemoveConstraint(p)
  testEq(t, s.Contains(b), false)
  testEq(t, s.Contains(b2), false)
  testEq(t, s.Contains(c), false)
  testEq(t, s.Contains(p), false)

  b.Free()
  b2.Free()
  c.Free()
  p.Free()

  s.Free()
}

// Local Variables:
// indent-tabs-mode: nil
// tab-width: 2
// End:
// ex: set tabstop=2 shiftwidth=2 expandtab:
