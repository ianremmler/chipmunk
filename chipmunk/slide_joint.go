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

////////////////////////////////////////////////////////////////////////////////

// SlideJoint holds the distance between points on two bodies between a minimum and a maximum.
// Think of them as a telescoping PinJoint.
type SlideJoint struct {
  constraintBase
}

////////////////////////////////////////////////////////////////////////////////

// Anchr1 returns the anchor point on the first body.
func (c SlideJoint) Anchr1() Vect {
  return cpVect(C.cpSlideJointGetAnchr1(c.ct))
}

// Anchr2 returns the anchor point on the second body.
func (c SlideJoint) Anchr2() Vect {
  return cpVect(C.cpSlideJointGetAnchr2(c.ct))
}

// Max returns the maximum allowed distance between anchor points.
func (c SlideJoint) Max() float64 {
  return float64(C.cpSlideJointGetMax(c.ct))
}

// Min returns the minimum allowed distance between anchor points.
func (c SlideJoint) Min() float64 {
  return float64(C.cpSlideJointGetMin(c.ct))
}

// SetAnchr1 sets the anchor point on the first body.
func (c SlideJoint) SetAnchr1(v Vect) {
  C.cpSlideJointSetAnchr1(c.ct, v.c())
}

// SetAnchr2 sets the anchor point on the second body.
func (c SlideJoint) SetAnchr2(v Vect) {
  C.cpSlideJointSetAnchr2(c.ct, v.c())
}

// SetMax sets the maximum allowed distance between anchor points.
func (c SlideJoint) SetMax(m float64) {
  C.cpSlideJointSetMax(c.ct, C.cpFloat(m))
}

// SetMin sets the minimum allowed distance between anchor points.
func (c SlideJoint) SetMin(m float64) {
  C.cpSlideJointSetMin(c.ct, C.cpFloat(m))
}

// SlideJointNew creates a new slide joint.
func SlideJointNew(a, b Body, anchr1, anchr2 Vect, min, max float64) SlideJoint {
  c := C.cpSlideJointNew(a.c(), b.c(), anchr1.c(), anchr2.c(), C.cpFloat(min), C.cpFloat(max))
  return SlideJoint{constraintBase{c}}
}

// Local Variables:
// indent-tabs-mode: nil
// tab-width: 2
// End:
// ex: set tabstop=2 shiftwidth=2 expandtab:
