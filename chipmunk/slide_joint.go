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

type SlideJoint struct {
  constraintBase
}

func SlideJointNew(a, b Body, anchr1, anchr2 Vect, min, max float64) SlideJoint {
  return SlideJoint{
    constraintBase{
      C.cpSlideJointNew(a.c(), b.c(), anchr1.c(), anchr2.c(), C.cpFloat(min), C.cpFloat(max))}}
}

/////////////////////////////////////////////////////////////////////////////

func (c SlideJoint) Anchr1() Vect {
  return cpVect(C.cpSlideJointGetAnchr1(c.ct))
}

func (c SlideJoint) Anchr2() Vect {
  return cpVect(C.cpSlideJointGetAnchr2(c.ct))
}

func (c SlideJoint) Min() float64 {
  return float64(C.cpSlideJointGetMin(c.ct))
}

func (c SlideJoint) Max() float64 {
  return float64(C.cpSlideJointGetMax(c.ct))
}

/////////////////////////////////////////////////////////////////////////////

func (c SlideJoint) SetAnchr1(v Vect) {
  C.cpSlideJointSetAnchr1(c.ct, v.c())
}

func (c SlideJoint) SetAnchr2(v Vect) {
  C.cpSlideJointSetAnchr2(c.ct, v.c())
}

func (c SlideJoint) SetMin(m float64) {
  C.cpSlideJointSetMin(c.ct, C.cpFloat(m))
}

func (c SlideJoint) SetMax(m float64) {
  C.cpSlideJointSetMax(c.ct, C.cpFloat(m))
}
