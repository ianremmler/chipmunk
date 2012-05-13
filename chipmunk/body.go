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

/*
#include <chipmunk.h>

extern void bodyIterator(cpBody *body, void *d, void *p);

static void body_each_shape(cpBody *body, void *p) {
  cpBodyEachShape(body, (void *)bodyIterator, p);
}

static void body_each_constraint(cpBody *body, void *p) {
  cpBodyEachConstraint(body, (void *)bodyIterator, p);
}

static void body_each_arbiter(cpBody *body, void *p) {
  cpBodyEachArbiter(body, (void *)bodyIterator, p);
}

*/
import "C"

import (
  "unsafe"
)

// Space is a basic unit of simulation in Chipmunk.
type Body struct {
  b *C.cpBody
}

// NewBody creates a new body.
func NewBody(m, i float64) Body {
  return Body{ b : C.cpBodyNew(C.cpFloat(m), C.cpFloat(i)) }
}

func NewBodyStatic() Body {
  return Body{ b : C.cpBodyNewStatic() }
}

// Destroy removes a space.
func (b Body) Destroy() {
  C.cpBodyDestroy(b.b)
  b.b = nil
}

/////////////////////////////////////////////////////////////////////////////

func (b Body) ContainedInSpace(s Space) bool {
  return cpBool(C.cpSpaceContainsBody(s.s, b.b))
}

// Mass returns the mass of the body.
func (b Body) Mass() float64 {
  return float64(C.cpBodyGetMass(b.b))
}

// Moment returns a moment of intertia of the body.
func (b Body) Moment() float64 {
  return float64(C.cpBodyGetMoment(b.b))
}

// Position returns the position of the rigid body's center of gravity.
func (b Body) Position() Vect {
  return cpVect(C.cpBodyGetPos(b.b))
}

// Velocity returns the velocity of the rigid body's center of gravity.
func (b Body) Velocity() Vect {
  return cpVect(C.cpBodyGetVel(b.b))
}

// Force returns the force acting on the rigid body's center of gravity.
func (b Body) Force() Vect {
  return cpVect(C.cpBodyGetForce(b.b))
}

// Angle returns the rotation of the body around it's center of gravity in radians.
func (b Body) Angle() float64 {
  return float64(C.cpBodyGetAngle(b.b))
}

// AngularVelocity returns the angular velocity of the body around it's center of gravity in radians/second.
func (b Body) AngularVelocity() float64 {
  return float64(C.cpBodyGetAngVel(b.b))
}

// Torque returns the torque applied to the body around it's center of gravity.
func (b Body) Torque() float64 {
  return float64(C.cpBodyGetTorque(b.b))
}

// Rotation returns the cached unit length vector representing the angle of the body.
func (b Body) Rotation() Vect {
  return cpVect(C.cpBodyGetRot(b.b))
}

// VelocityLimit returns the maximum velocity allowed when updating the velocity.
func (b Body) VelocityLimit() float64 {
  return float64(C.cpBodyGetVelLimit(b.b))
}

// AngularVelocityLimit returns the maximum rotational rate (in radians/second) allowed when updating
// the angular velocity.
func (b Body) AngularVelocityLimit() float64 {
  return float64(C.cpBodyGetAngVelLimit(b.b))
}

func (b Body) UserData() interface{} {
  return cpData(C.cpBodyGetUserData(b.b))
}

func (b Body) SetUserData(data interface{}) {
  C.cpBodySetUserData(b.b, dataToC(data))
}

/////////////////////////////////////////////////////////////////////////////

func (b Body) UpdateVelocity(gravity Vect, damping float64, dt float64) {
  C.cpBodyUpdateVelocity(b.b, gravity.c(), C.cpFloat(damping), C.cpFloat(dt))
}

func (b Body) UpdatePosition(dt float64) {
  C.cpBodyUpdatePosition(b.b, C.cpFloat(dt))
}

func (b Body) LocalToWorld(v Vect) Vect {
  return cpVect(C.cpBodyWorld2Local(b.b, v.c()))
}

// ResetForces sets the forces and torque of a body to zero.
func (b Body) ResetForces() {
  C.cpBodyResetForces(b.b)
}

// ApplyForce applies a force (in world coordinates) to the body at a point relative
// to the center of gravity (also in world coordinates).
func (b Body) ApplyForce(f, r Vect) {
  C.cpBodyApplyForce(b.b, f.c(), r.c())
}

// ApplyImpulse applies an impulse (in world coordinates) to the body at a point relative
// to the center of gravity (also in world coordinates).
func (b Body) ApplyImpulse(j, r Vect) {
  C.cpBodyApplyImpulse(b.b, j.c(), r.c())
}

// VelocityAtWorldPoint returns the velocity on a body (in world units) at a point
// on the body in world coordinates.
func (b Body) VelocityAtWorldPoint(point Vect) Vect {
  return cpVect(C.cpBodyGetVelAtWorldPoint(b.b, point.c()))
}

// VelocityAtLocalPoint returns the velocity on a body (in world units) at a point
// on the body in local coordinates.
func (b Body) VelocityAtLocalPoint(point Vect) Vect {
  return cpVect(C.cpBodyGetVelAtLocalPoint(b.b, point.c()))
}

// KineticEnergy returns the kinetic energy of a body.
func (b Body) KineticEnergy() float64 {
  return float64(C.cpBodyKineticEnergy(b.b))
}

/////////////////////////////////////////////////////////////////////////////

// Activate wakes up a sleeping or idle body.
func (b Body) Activate() {
  if nil != b.b {
    C.cpBodyActivate(b.b)
  }
}

// ActivateStatic wakes up any sleeping or idle bodies touching a static body.
func (b Body) ActivateStatic(s Shape) {
  C.cpBodyActivateStatic(b.b, s.c())
}

// Sleep forces a body to fall asleep immediately.
func (b Body) Sleep() {
  C.cpBodySleep(b.b)
}

// SleepWithGroup forces a body to fall asleep immediately along with other bodies in a group.
func (b Body) SleepWithGroup(g Body) {
  C.cpBodySleepWithGroup(b.b, g.b)
}

// IsSleeping returns true if the body is sleeping.
func (b Body) IsSleeping() bool {
  return cpBool(C.cpBodyIsSleeping(b.b))
}

// IsStatic returns true if the body is static.
func (b Body) IsStatic() bool {
  return cpBool(C.cpBodyIsStatic(b.b))
}

// IsRogue returns true if the body has not been added to a space.
func (b Body) IsRogue() bool {
  return cpBool(C.cpBodyIsRogue(b.b))
}

/////////////////////////////////////////////////////////////////////////////

type ShapeIterator func(b Body, s Shape)

type ConstraintIterator func(b Body, c Constraint)

type ArbiterIterator func(b Body, a Arbiter)

//export bodyIterator
func bodyIterator(b *C.cpBody, d unsafe.Pointer, p unsafe.Pointer) {
  f := *(*interface{})(p)
  body := Body{ b }

  switch f.(type) {
  case ShapeIterator:
    f.(ShapeIterator)(body, cpShape((*C.cpShape)(d)))

  case ConstraintIterator:
    f.(ConstraintIterator)(body, cpConstraint((*C.cpConstraint)(d)))

  case ArbiterIterator:
    f.(ArbiterIterator)(body, Arbiter{ a : (*C.cpArbiter)(d) })

  default:
    panic("invalid type of iterator in body_iterator")
  }
}

// ForEach calls a callback function func once for each shape/constraint attached to a body and added to the space or each arbiter that is currently active on the body.
// What exactly should this function iterate over is decided by the callback type.
func (b Body) ForEach(f interface{}) {
  p := unsafe.Pointer(&f)

  switch f.(type) {
  case ShapeIterator:
    C.body_each_shape(b.b, p)

  case ConstraintIterator:
    C.body_each_constraint(b.b, p)

  case ArbiterIterator:
    C.body_each_arbiter(b.b, p)

  default:
    panic("invalid type of iterator in ForEach")
  }
}

func (b Body) c() *C.cpBody {
  return b.b
}

func cpBody(b *C.cpBody) Body {
  return Body{ b }
}
