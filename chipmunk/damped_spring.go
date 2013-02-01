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

// DampedSprint is a spring with a damper.
// While a spring is not technically a constraint, the damper is.
// The spring forces are simply a convenience.
type DampedSpring struct {
  constraintBase
}

////////////////////////////////////////////////////////////////////////////////

// Anchr1 returns the anchor point on the first body.
func (c DampedSpring) Anchr1() Vect {
  return cpVect(C.cpDampedSpringGetAnchr1(c.c()))
}

// Anchr2 returns the anchor point on the second body.
func (c DampedSpring) Anchr2() Vect {
  return cpVect(C.cpDampedSpringGetAnchr2(c.c()))
}

// DampedSprintNew creates a new damped spring.
func DampedSpringNew(
  a, b Body,
  anchr1, anchr2 Vect,
  restLength, stiffness, damping float64) DampedSpring {

  c := C.cpDampedSpringNew(
    a.c(),
    b.c(),
    anchr1.c(),
    anchr2.c(),
    C.cpFloat(restLength),
    C.cpFloat(stiffness),
    C.cpFloat(damping))

  return DampedSpring{cpconstraint_new(c)}
}

// Damping returns the amount of viscous damping to apply.
func (c DampedSpring) Damping() float64 {
  return float64(C.cpDampedSpringGetDamping(c.c()))
}

// RestLength returns the length the spring wants to contract or expand to.
func (c DampedSpring) RestLength() float64 {
  return float64(C.cpDampedSpringGetRestLength(c.c()))
}

// SetAnchr1 sets the anchor point on the first body.
func (c DampedSpring) SetAnchr1(v Vect) {
  C.cpDampedSpringSetAnchr1(c.c(), v.c())
}

// SetAnchr2 sets the anchor point on the second body.
func (c DampedSpring) SetAnchr2(v Vect) {
  C.cpDampedSpringSetAnchr2(c.c(), v.c())
}

// SetDamping sets the amount of viscous damping to apply.
func (c DampedSpring) SetDamping(damping float64) {
  C.cpDampedSpringSetDamping(c.c(), C.cpFloat(damping))
}

// SetRestLength sets the length the spring wants to contract or expand to.
func (c DampedSpring) SetRestLength(restLength float64) {
  C.cpDampedSpringSetRestLength(c.c(), C.cpFloat(restLength))
}

// SetStiffness sets the young's modulus of the spring.
func (c DampedSpring) SetStiffness(stiffness float64) {
  C.cpDampedSpringSetStiffness(c.c(), C.cpFloat(stiffness))
}

// Stiffness returns the young's modulus of the spring.
func (c DampedSpring) Stiffness() float64 {
  return float64(C.cpDampedSpringGetStiffness(c.c()))
}

// Local Variables:
// indent-tabs-mode: nil
// tab-width: 2
// End:
// ex: set tabstop=2 shiftwidth=2 expandtab:
