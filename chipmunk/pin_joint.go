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

type PinJoint struct {
  constraintBase
}

func PinJointNew(a, b Body, anchr1, anchr2 Vect) PinJoint {
  return PinJoint{ constraintBase{ C.cpPinJointNew(a.c(), b.c(), anchr1.c(), anchr2.c()) } }
}

/////////////////////////////////////////////////////////////////////////////

func (j PinJoint) Anchr1() Vect {
  return cpVect(C.cpPinJointGetAnchr1(j.ct))
}

func (j PinJoint) Anchr2() Vect {
  return cpVect(C.cpPinJointGetAnchr2(j.ct))
}

func (j PinJoint) Dist() float64 {
  return float64(C.cpPinJointGetDist(j.ct))
}

/////////////////////////////////////////////////////////////////////////////

func (j PinJoint) SetAnchr1(v Vect) {
  C.cpPinJointSetAnchr1(j.ct, v.c())
}

func (j PinJoint) SetAnchr2(v Vect) {
  C.cpPinJointSetAnchr2(j.ct, v.c())
}

func (j PinJoint) SetDist(d float64) {
  C.cpPinJointSetDist(j.ct, C.cpFloat(d))
}
