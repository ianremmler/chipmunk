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

func testEq(t *testing.T, r, m interface{}) {
  if r != m {
    t.Fatal("got", r, "instead of", m)
  }
}

func Test_VectNew(t *testing.T) {
  v := VectNew(1.2345, -2.34567)

  if v.X != 1.2345 || v.Y != -2.34567 {
    t.Fatal("wrong X/Y")
  }
}

// FIXME func Test_VectForAngle(t *testing.T) {

func Test_VectOrigin(t *testing.T) {
  v := Origin()

  if v.X != 0.0 || v.Y != 0.0 {
    t.Fatal("wrong X/Y for Origin")
  }
}

func Test_VectString(t *testing.T) {
  testEq(t, Origin().String(), "(0.000000, 0.000000)")
}

func Test_VectAdd(t *testing.T) {
  testEq(t, VectNew(1.5, -2.5).Add(VectNew(-1.5, 2.5)), Origin())
}

func Test_VectSub(t *testing.T) {
  testEq(t, VectNew(1.5, -2.5).Sub(VectNew(1.5, -2.5)), Origin())
}

func Test_VectMul(t *testing.T) {
  testEq(t, VectNew(1.5, -2.5).Mul(3.0), VectNew(4.5, -7.5))
}

func Test_VectDiv(t *testing.T) {
  testEq(t, VectNew(4.5, -7.5).Div(-3.0), VectNew(-1.5, 2.5))
}

func Test_VectDot(t *testing.T) {
  testEq(t, VectNew(3.0, 1.0).Dot(VectNew(2.0, 4.0)), 10.0)
}

func Test_VectLength(t *testing.T) {
  testEq(t, VectNew(0.0, -9.0).Length(), 9.0)
}

func Test_VectNeg(t *testing.T) {
  testEq(t, VectNew(-1.5, 999.9).Neg(), VectNew(1.5, -999.9))
}

// FIXME func Test_VectLerp(t *testing.T) {
// FIXME func Test_VectLerpConst(t *testing.T) {
// FIXME func Test_VectToAngle(t *testing.T) {
// FIXME func Test_Vect_c(t *testing.T) {
// FIXME func Test_Vect_cpVect(t *testing.T) {

// Local Variables:
// indent-tabs-mode: nil
// tab-width: 2
// End:
// ex: set tabstop=2 shiftwidth=2 expandtab:
