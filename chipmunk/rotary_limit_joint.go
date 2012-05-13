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

// #include <chipmunk.h>
import "C"

type RotaryLimitJoint struct {
  constraintBase
}

func RotaryLimitJointNew(a, b Body, min, max float64) RotaryLimitJoint {
  return RotaryLimitJoint{ constraintBase{ C.cpRotaryLimitJointNew(a.c(), b.c(), C.cpFloat(min), C.cpFloat(max)) } }
}

/////////////////////////////////////////////////////////////////////////////

func (j RotaryLimitJoint) Min() float64 {
  return float64(C.cpRotaryLimitJointGetMin(j.ct))
}

func (j RotaryLimitJoint) Max() float64 {
  return float64(C.cpRotaryLimitJointGetMax(j.ct))
}

/////////////////////////////////////////////////////////////////////////////

func (j RotaryLimitJoint) SetMin(m float64) {
  C.cpRotaryLimitJointSetMin(j.ct, C.cpFloat(m))
}

func (j RotaryLimitJoint) SetMax(m float64) {
  C.cpRotaryLimitJointSetMax(j.ct, C.cpFloat(m))
}
