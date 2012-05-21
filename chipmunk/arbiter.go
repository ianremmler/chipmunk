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

import (
  "fmt"
  . "unsafe"
)

////////////////////////////////////////////////////////////////////////////////

// Arbiter is a type of colliding pair of shapes.
type Arbiter uintptr

// ContactPoint is a contact point type of collision.
type ContactPoint struct {
  // Point is position normal of the contact point.
  Point Vect
  // Normal is the normal of the contact point.
  Normat Vect
  // Dist is the depth of the contact point.
  Dist float64
}

////////////////////////////////////////////////////////////////////////////////

// Bodies returns the colliding bodies involved for this arbiter.
// The order of the CollisionType the bodies are associated with values will match
// the order set when the collision handler was registered.
func (arb Arbiter) Bodies() (Body, Body) {
  var a, b *C.cpBody
  C.cpArbiterGetBodies(arb.c(), (**C.cpBody)(Pointer(&a)), (**C.cpBody)(Pointer(&b)))
  return cpBody(a), cpBody(b)
}

// ContactPoints returns a contact set from an arbiter.
func (a Arbiter) ContactPoints() []ContactPoint {
  set := C.cpArbiterGetContactPointSet(a.c())
  c := make([]ContactPoint, int(set.count))

  for i := range c {
    c[i] = ContactPoint{
      cpVect(set.points[i].point),
      cpVect(set.points[i].normal),
      float64(set.points[i].dist)}
  }

  return c
}

// Count returns the number of contact points for this arbiter.
func (a Arbiter) Count() int {
  return int(C.cpArbiterGetCount(a.c()))
}

// Depth returns the depth of specific contact point.
func (a Arbiter) Depth(i int) float64 {
  return float64(C.cpArbiterGetDepth(a.c(), C.int(i)))
}

// Elasticity returns a calculated value to use for the elasticity coefficient.
// Override in a pre-solve collision handler for custom behavior.
func (a Arbiter) Elasticity() float64 {
  return float64(C.cpArbiterGetElasticity(a.c()))
}

// Friction returns a calculated value to use for the friction coefficient.
// Override in a pre-solve collision handler for custom behavior.
func (a Arbiter) Friction() float64 {
  return float64(C.cpArbiterGetFriction(a.c()))
}

// Ignore causes a collision pair to be ignored as if you returned false from a begin callback.
// If called from a pre-step callback, you will still need to return false
// if you want it to be ignored in the current step.
func (a Arbiter) Ignore() {
  C.cpArbiterIgnore(a.c())
}

// IsFirstContact returns true if this is the first step a pair of objects started colliding.
func (a Arbiter) IsFirstContact() bool {
  return cpBool(C.cpArbiterIsFirstContact(a.c()))
}

// Normal returns the normal of specific contact point.
func (a Arbiter) Normal(i int) Vect {
  return cpVect(C.cpArbiterGetNormal(a.c(), C.int(i)))
}

// Point returns the position of specific contact point.
func (a Arbiter) Point(i int) Vect {
  return cpVect(C.cpArbiterGetPoint(a.c(), C.int(i)))
}

// SetElasticity sets elasticity coefficient.
func (a Arbiter) SetElasticity(e float64) {
  C.cpArbiterSetElasticity(a.c(), C.cpFloat(e))
}

// SetFriction sets friction coefficient.
func (a Arbiter) SetFriction(f float64) {
  C.cpArbiterSetFriction(a.c(), C.cpFloat(f))
}

// SetSurfaceVelocity sets calculated value to use for applying surface velocities.
func (a Arbiter) SetSurfaceVelocity(v Vect) {
  C.cpArbiterSetSurfaceVelocity(a.c(), v.c())
}

// Shapes returns the colliding shapes involved for this arbiter.
// The order of their CollisionType values will match the order set when the collision
// handler was registered.
func (arb Arbiter) Shapes() (Shape, Shape) {
  var a, b *C.cpShape
  C.cpArbiterGetShapes(arb.c(), (**C.cpShape)(Pointer(&a)), (**C.cpShape)(Pointer(&b)))
  return cpShape(a), cpShape(b)
}

// String converts an arbiter to a human-readable string.
func (a Arbiter) String() string {
  return fmt.Sprintf("(Arbiter)%+v", a.c())
}

// SurfaceVelocity returns a calculated value to use for applying surface velocities.
// Override in a pre-solve collision handler for custom behavior.
func (a Arbiter) SurfaceVelocity() Vect {
  return cpVect(C.cpArbiterGetSurfaceVelocity(a.c()))
}

// TotalImpulse returns the total impulse that was applied by this arbiter.
// This function should only be called from a post-solve, post-step or
// body.EachArbiter callback.
func (a Arbiter) TotalImpulse() Vect {
  return cpVect(C.cpArbiterTotalImpulse(a.c()))
}

// TotalImpulseWithFriction returns the total impulse including the friction that
// was applied by this arbiter. This function should only be called from a post-solve,
// post-step or body.EachArbiter callback.
func (a Arbiter) TotalImpulseWithFriction() Vect {
  return cpVect(C.cpArbiterTotalImpulseWithFriction(a.c()))
}

// TotalKE returns the amount of energy lost in a collision including static,
// but not dynamic friction. This function should only be called from a post-solve,
// post-step or body.EachArbiter callback.
func (a Arbiter) TotalKE() float64 {
  return float64(C.cpArbiterTotalKE(a.c()))
}

// c converts Arbiter to c.cpArbiter pointer.
func (a Arbiter) c() *C.cpArbiter {
  return (*C.cpArbiter)(Pointer(a))
}

// cpArbiter converts C.cpArbiter pointer to Arbiter.
func cpArbiter(a *C.cpArbiter) Arbiter {
  return Arbiter(Pointer(a))
}

// Local Variables:
// indent-tabs-mode: nil
// tab-width: 2
// End:
// ex: set tabstop=2 shiftwidth=2 expandtab:
