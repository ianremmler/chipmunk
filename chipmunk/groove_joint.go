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

// GrooveJoint holds a pivot point on one body to line along a line segment on
// another like a pin in a groove.
type GrooveJoint struct {
  constraintBase
}

////////////////////////////////////////////////////////////////////////////////

// Anchr2 returns the anchor point on the second body that is held to the line segment on the first.
func (c GrooveJoint) Anchr2() Vect {
  return cpVect(C.cpGrooveJointGetAnchr2(c.c()))
}

// GrooveA returns the start of the line segment on the first body.
func (c GrooveJoint) GrooveA() Vect {
  return cpVect(C.cpGrooveJointGetGrooveA(c.c()))
}

// GrooveB returns the end of the line segment on the first body.
func (c GrooveJoint) GrooveB() Vect {
  return cpVect(C.cpGrooveJointGetGrooveB(c.c()))
}

// GrooveJointNew creates a new groove joint.
// Make sure you have the bodies in the right place as the joint will snap
// into shape as soon as you start simulating the space.
func GrooveJointNew(a, b Body, groove_a, groove_b, anchr2 Vect) GrooveJoint {
  c := C.cpGrooveJointNew(a.c(), b.c(), groove_a.c(), groove_b.c(), anchr2.c())
  return GrooveJoint{cpconstraint_new(c)}
}

// SetAnchr2 sets the anchor point on the second body that is held to the line segment on the first.
func (c GrooveJoint) SetAnchr2(v Vect) {
  C.cpGrooveJointSetAnchr2(c.c(), v.c())
}

// SetGrooveA sets the start of the line segment on the first body.
func (c GrooveJoint) SetGrooveA(v Vect) {
  C.cpGrooveJointSetGrooveA(c.c(), v.c())
}

// SetGrooveB sets the end of the line segment on the first body.
func (c GrooveJoint) SetGrooveB(v Vect) {
  C.cpGrooveJointSetGrooveB(c.c(), v.c())
}

// Local Variables:
// indent-tabs-mode: nil
// tab-width: 2
// End:
// ex: set tabstop=2 shiftwidth=2 expandtab:
