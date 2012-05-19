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

// PinJoint holds a set distance between points on two bodies.
// Think of them as connecting a solid pin or rod between the two anchor points.
type PinJoint struct {
  constraintBase
}

////////////////////////////////////////////////////////////////////////////////

// Anchr1 returns the anchor point on the first body.
func (c PinJoint) Anchr1() Vect {
  return cpVect(C.cpPinJointGetAnchr1(c.ct))
}

// Anchr2 returns the anchor point on the second body.
func (c PinJoint) Anchr2() Vect {
  return cpVect(C.cpPinJointGetAnchr2(c.ct))
}

// Dist returns the distance between the two anchor points that the joint keeps.
func (c PinJoint) Dist() float64 {
  return float64(C.cpPinJointGetDist(c.ct))
}

// PinJointNew creates a new pin joint.
func PinJointNew(a, b Body, anchr1, anchr2 Vect) PinJoint {
  return PinJoint{constraintBase{C.cpPinJointNew(a.c(), b.c(), anchr1.c(), anchr2.c())}}
}

// SetAnchr1 sets the anchor point on the first body.
func (c PinJoint) SetAnchr1(v Vect) {
  C.cpPinJointSetAnchr1(c.ct, v.c())
}

// SetAnchr2 sets the anchor point on the second body.
func (c PinJoint) SetAnchr2(v Vect) {
  C.cpPinJointSetAnchr2(c.ct, v.c())
}

// SetDist sets the distance between the two anchor points that the joint keeps.
func (c PinJoint) SetDist(d float64) {
  C.cpPinJointSetDist(c.ct, C.cpFloat(d))
}

// Local Variables:
// indent-tabs-mode: nil
// tab-width: 2
// End:
// ex: set tabstop=2 shiftwidth=2 expandtab:
