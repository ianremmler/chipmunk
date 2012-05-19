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

// RotaryLimitJoint constrains the angle between two bodies.
// This joint is often used in conjuction with a separate PivotJoint in order to limit
// the rotation around the pivot.
type RotaryLimitJoint struct {
  constraintBase
}

////////////////////////////////////////////////////////////////////////////////

// Max returns the maximum angular delta of the joint in radians.
func (c RotaryLimitJoint) Max() float64 {
  return float64(C.cpRotaryLimitJointGetMax(c.ct))
}

// Min returns the minimum angular delta of the joint in radians.
func (c RotaryLimitJoint) Min() float64 {
  return float64(C.cpRotaryLimitJointGetMin(c.ct))
}

// RotaryLimitJointNew creates a new rotary limit joint.
func RotaryLimitJointNew(a, b Body, min, max float64) RotaryLimitJoint {
  c := C.cpRotaryLimitJointNew(a.c(), b.c(), C.cpFloat(min), C.cpFloat(max))
  return RotaryLimitJoint{constraintBase{c}}
}

// SetMax sets the maximum angular delta of the joint in radians.
func (c RotaryLimitJoint) SetMax(m float64) {
  C.cpRotaryLimitJointSetMax(c.ct, C.cpFloat(m))
}

// SetMin sets the minimum angular delta of the joint in radians.
func (c RotaryLimitJoint) SetMin(m float64) {
  C.cpRotaryLimitJointSetMin(c.ct, C.cpFloat(m))
}

// Local Variables:
// indent-tabs-mode: nil
// tab-width: 2
// End:
// ex: set tabstop=2 shiftwidth=2 expandtab:
