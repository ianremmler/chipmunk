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

extern void eachArbiter_body(cpBody *b, cpArbiter *a, void *p);
extern void eachConstraint_body(cpBody *b, cpConstraint *c, void *p);
extern void eachShape_body(cpBody *b, cpShape *s, void *p);

static void body_each_arbiter(cpBody *body, void *f) {
  cpBodyEachArbiter(body, eachArbiter_body, f);
}

static void body_each_constraint(cpBody *body, void *f) {
  cpBodyEachConstraint(body, eachConstraint_body, f);
}

static void body_each_shape(cpBody *body, void *f) {
  cpBodyEachShape(body, eachShape_body, f);
}
*/
import "C"

import (
  "fmt"
  "unsafe"
)

////////////////////////////////////////////////////////////////////////////////

// Body is a rigid body struct.
type Body struct {
  b *C.cpBody
}

////////////////////////////////////////////////////////////////////////////////

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

// Angle returns the rotation of the body around it's center of gravity in radians.
func (b Body) Angle() float64 {
  return float64(C.cpBodyGetAngle(b.b))
}

// AngularVelocity returns the angular velocity of the body around it's center
// of gravity in radians/second.
func (b Body) AngularVelocity() float64 {
  return float64(C.cpBodyGetAngVel(b.b))
}

// AngularVelocityLimit returns the maximum rotational rate (in radians/second) allowed when updating
// the angular velocity.
func (b Body) AngularVelocityLimit() float64 {
  return float64(C.cpBodyGetAngVelLimit(b.b))
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

// BodyNew creates a new body.
func BodyNew(m, i float64) Body {
  return Body{C.cpBodyNew(C.cpFloat(m), C.cpFloat(i))}
}

// BodyStaticNew creates a new static body.
func BodyStaticNew() Body {
  return Body{C.cpBodyNewStatic()}
}

// ContainedInSpace returns true if the body is in the space.
func (b Body) ContainedInSpace(s Space) bool {
  return cpBool(C.cpSpaceContainsBody(s.c(), b.b))
}

// EachArbiter calls a callback function once for each arbiter which is currently
// active on the body.
func (b Body) EachArbiter(iter func(Body, Shape)) {
  p := unsafe.Pointer(&iter)
  C.body_each_arbiter(b.b, p)
}

// EachConstraint calls a callback function once for each constraint attached
// to the body and added to the space.
func (b Body) EachConstraint(iter func(Body, Constraint)) {
  p := unsafe.Pointer(&iter)
  C.body_each_constraint(b.b, p)
}

// EachShape calls a callback function once for each shape attached to the body
// and added to the space.
func (b Body) EachShape(iter func(Body, Shape)) {
  p := unsafe.Pointer(&iter)
  C.body_each_shape(b.b, p)
}

// Free removes a body.
func (b Body) Free() {
  C.cpBodyFree(b.b)
}

// Force returns the force acting on the rigid body's center of gravity.
func (b Body) Force() Vect {
  return cpVect(C.cpBodyGetForce(b.b))
}

// IsRogue returns true if the body has not been added to a space.
func (b Body) IsRogue() bool {
  return cpBool(C.cpBodyIsRogue(b.b))
}

// IsSleeping returns true if the body is sleeping.
func (b Body) IsSleeping() bool {
  return cpBool(C.cpBodyIsSleeping(b.b))
}

// IsStatic returns true if the body is static.
func (b Body) IsStatic() bool {
  return cpBool(C.cpBodyIsStatic(b.b))
}

// KineticEnergy returns the kinetic energy of a body.
func (b Body) KineticEnergy() float64 {
  return float64(C.cpBodyKineticEnergy(b.b))
}

// LocalToWorld converts body relative/local coordinates to absolute/world coordinates.
func (b Body) LocalToWorld(v Vect) Vect {
  return cpVect(C.cpBodyWorld2Local(b.b, v.c()))
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

// ResetForces sets the forces and torque of a body to zero.
func (b Body) ResetForces() {
  C.cpBodyResetForces(b.b)
}

// Rotation returns the cached unit length vector representing the angle of the body.
func (b Body) Rotation() Vect {
  return cpVect(C.cpBodyGetRot(b.b))
}

// SetPosition sets the position of the body.
func (b Body) SetPosition(v Vect) {
  C.cpBodySetPos(b.b, v.c())
}

// SetUserData sets user definable data pointer.
// Generally this points to your the game object so you can access it
// when given a Body reference in a callback.
func (b Body) SetUserData(data interface{}) {
  C.cpBodySetUserData(b.b, dataToC(data))
}

// Sleep forces a body to fall asleep immediately.
func (b Body) Sleep() {
  C.cpBodySleep(b.b)
}

// SleepWithGroup forces a body to fall asleep immediately along with other bodies in a group.
func (b Body) SleepWithGroup(g Body) {
  C.cpBodySleepWithGroup(b.b, g.b)
}

// Space returns space the body was added to or nil if the body doesn't belong to any space.
func (b Body) Space() Space {
  return cpSpace(C.cpBodyGetSpace(b.b))
}

// String converts a body to a human-readable string.
func (b Body) String() string {
  return fmt.Sprintf("(Body)%+v", b.b)
}

// Torque returns the torque applied to the body around it's center of gravity.
func (b Body) Torque() float64 {
  return float64(C.cpBodyGetTorque(b.b))
}

// UpdatePosition is a default function that is called to integrate the body's position.
func (b Body) UpdatePosition(dt float64) {
  C.cpBodyUpdatePosition(b.b, C.cpFloat(dt))
}

// UpdateVelocity is a default function that is called to integrate the body's velocity.
func (b Body) UpdateVelocity(gravity Vect, damping float64, dt float64) {
  C.cpBodyUpdateVelocity(b.b, gravity.c(), C.cpFloat(damping), C.cpFloat(dt))
}

// UserData returns user defined data.
func (b Body) UserData() interface{} {
  return cpData(C.cpBodyGetUserData(b.b))
}

// Velocity returns the velocity of the rigid body's center of gravity.
func (b Body) Velocity() Vect {
  return cpVect(C.cpBodyGetVel(b.b))
}

// VelocityAtLocalPoint returns the velocity on a body (in world units) at a point
// on the body in local coordinates.
func (b Body) VelocityAtLocalPoint(point Vect) Vect {
  return cpVect(C.cpBodyGetVelAtLocalPoint(b.b, point.c()))
}

// VelocityAtWorldPoint returns the velocity on a body (in world units) at a point
// on the body in world coordinates.
func (b Body) VelocityAtWorldPoint(point Vect) Vect {
  return cpVect(C.cpBodyGetVelAtWorldPoint(b.b, point.c()))
}

// VelocityLimit returns the maximum velocity allowed when updating the velocity.
func (b Body) VelocityLimit() float64 {
  return float64(C.cpBodyGetVelLimit(b.b))
}

//export eachShape_body
func eachShape_body(b *C.cpBody, sh *C.cpShape, p unsafe.Pointer) {
  f := *(*func(Body, Shape))(p)
  f(cpBody(b), cpShape(sh))
}

//export eachConstraint_body
func eachConstraint_body(b *C.cpBody, c *C.cpConstraint, p unsafe.Pointer) {
  f := *(*func(Body, Constraint))(p)
  f(cpBody(b), cpConstraint(c))
}

//export eachArbiter_body
func eachArbiter_body(b *C.cpBody, a *C.cpArbiter, p unsafe.Pointer) {
  f := *(*func(Body, Arbiter))(p)
  f(cpBody(b), cpArbiter(a))
}

// c converts Body to c.cpBody pointer.
func (b Body) c() *C.cpBody {
  return b.b
}

// cpBody converts C.cpBody pointer to Body.
func cpBody(b *C.cpBody) Body {
  return Body{b}
}

// Local Variables:
// indent-tabs-mode: nil
// tab-width: 2
// End:
// ex: set tabstop=2 shiftwidth=2 expandtab:
