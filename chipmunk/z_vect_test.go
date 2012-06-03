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

func Test_VectNew(t *testing.T) {
  v := VectNew(1.2345, -2.34567)

  assert.Equal(t, 1.2345, v.X)
  assert.Equal(t, -2.34567, v.Y)
}

// FIXME func Test_VectForAngle(t *testing.T) {

func Test_VectOrigin(t *testing.T) {
  v := Origin()

  assert.Equal(t, 0.0, v.X)
  assert.Equal(t, 0.0, v.Y)
}

func Test_VectString(t *testing.T) {
  assert.Equal(t, "(Vect){0, 0}", Origin().String())
  assert.Equal(t, "(Vect){-1.3, 4.55}", VectNew(-1.3, 4.55).String())
}

func Test_VectAdd(t *testing.T) {
  assert.Equal(t, Origin(), VectNew(1.5, -2.5).Add(VectNew(-1.5, 2.5)))
}

func Test_VectSub(t *testing.T) {
  assert.Equal(t, Origin(), VectNew(1.5, -2.5).Sub(VectNew(1.5, -2.5)))
}

func Test_VectMul(t *testing.T) {
  assert.Equal(t, VectNew(4.5, -7.5), VectNew(1.5, -2.5).Mul(3.0))
}

func Test_VectDiv(t *testing.T) {
  assert.Equal(t, VectNew(-1.5, 2.5), VectNew(4.5, -7.5).Div(-3.0))
}

func Test_VectDot(t *testing.T) {
  assert.Equal(t, 10.0, VectNew(3.0, 1.0).Dot(VectNew(2.0, 4.0)))
}

func Test_VectLength(t *testing.T) {
  assert.Equal(t, 9.0, VectNew(0.0, -9.0).Length())
}

func Test_VectNeg(t *testing.T) {
  assert.Equal(t, VectNew(1.5, -999.9), VectNew(-1.5, 999.9).Neg())
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
