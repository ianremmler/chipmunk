package chipmunk

/*
Copyright © 2012 Serge Zirukin

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

// #include <chipmunk/chipmunk.h>
import "C"

////////////////////////////////////////////////////////////////////////////////

// RatchetJoint creates rotary ratches similar to a socket wrench.
type RatchetJoint struct {
  constraintBase
}

////////////////////////////////////////////////////////////////////////////////

// Angle returns the current ratchet position in radians.
func (c RatchetJoint) Angle() float64 {
  return float64(C.cpRatchetJointGetAngle(c.c()))
}

// Phase returns the angular offset of the ratchet positions in radians.
func (c RatchetJoint) Phase() float64 {
  return float64(C.cpRatchetJointGetPhase(c.c()))
}

// Ratchet returns the angle in radians of each ratchet position.
func (c RatchetJoint) Ratchet() float64 {
  return float64(C.cpRatchetJointGetRatchet(c.c()))
}

// RatchetJointNew creates a new ratchet joint.
func RatchetJointNew(a, b Body, phase, ratchet float64) RatchetJoint {
  c := C.cpRatchetJointNew(a.c(), b.c(), C.cpFloat(phase), C.cpFloat(ratchet))
  return RatchetJoint{cpconstraint_new(c)}
}

// SetAngle sets the ratchet position in radians.
func (c RatchetJoint) SetAngle(m float64) {
  C.cpRatchetJointSetAngle(c.c(), C.cpFloat(m))
}

// SetPhase sets the angular offset of the ratchet positions in radians.
func (c RatchetJoint) SetPhase(m float64) {
  C.cpRatchetJointSetPhase(c.c(), C.cpFloat(m))
}

// SetRatchet sets the angle in radians of each ratchet position.
// Negative values cause the ratchet to operate in the opposite direction.
func (c RatchetJoint) SetRatchet(m float64) {
  C.cpRatchetJointSetRatchet(c.c(), C.cpFloat(m))
}

// Local Variables:
// indent-tabs-mode: nil
// tab-width: 2
// End:
// ex: set tabstop=2 shiftwidth=2 expandtab:
