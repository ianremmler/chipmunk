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

// DampedRotarySpring is like a DampedSpring, but operates in a rotational fashion.
type DampedRotarySpring struct {
  constraintBase
}

////////////////////////////////////////////////////////////////////////////////

// DampedRotarySpringNew creates a new damped rotary spring.
func DampedRotarySpringNew(a, b Body, restAngle, stiffness, damping float64) DampedRotarySpring {
  c := C.cpDampedRotarySpringNew(
    a.c(),
    b.c(),
    C.cpFloat(restAngle),
    C.cpFloat(stiffness),
    C.cpFloat(damping))

  return DampedRotarySpring{constraintBase{c}}
}

// Damping returns the amount of viscous damping to apply.
func (c DampedRotarySpring) Damping() float64 {
  return float64(C.cpDampedRotarySpringGetDamping(c.ct))
}

// RestAngle returns the angular offset in radians the spring attempts to keep between the two bodies.
func (c DampedRotarySpring) RestAngle() float64 {
  return float64(C.cpDampedRotarySpringGetRestAngle(c.ct))
}

// SetDamping sets the amount of viscous damping to apply.
func (c DampedRotarySpring) SetDamping(damping float64) {
  C.cpDampedRotarySpringSetDamping(c.ct, C.cpFloat(damping))
}

// SetRestAngle sets the angular offset in radians the spring attempts to keep between the two bodies.
func (c DampedRotarySpring) SetRestAngle(restAngle float64) {
  C.cpDampedRotarySpringSetRestAngle(c.ct, C.cpFloat(restAngle))
}

// SetStiffness sets the young's modulus of the spring.
func (c DampedRotarySpring) SetStiffness(stiffness float64) {
  C.cpDampedRotarySpringSetStiffness(c.ct, C.cpFloat(stiffness))
}

// Stiffness returns the young's modulus of the spring.
func (c DampedRotarySpring) Stiffness() float64 {
  return float64(C.cpDampedRotarySpringGetStiffness(c.ct))
}

// Local Variables:
// indent-tabs-mode: nil
// tab-width: 2
// End:
// ex: set tabstop=2 shiftwidth=2 expandtab:
