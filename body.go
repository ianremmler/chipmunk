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
// #include "body.h"
import "C"

import (
	"fmt"
	"unsafe"
)

////////////////////////////////////////////////////////////////////////////////

// Body is a rigid body struct.
type Body uintptr

type bodyData struct {
	positionFunc func(Body, float64)
	userData     interface{}
	velocityFunc func(Body, Vect, float64, float64)
}

////////////////////////////////////////////////////////////////////////////////

var (
	bodyDataMap = make(map[Body]*bodyData)
	nullBody    = Body(uintptr(0))
)

////////////////////////////////////////////////////////////////////////////////

// Activate wakes up a sleeping or idle body.
func (b Body) Activate() {
	if b != nullBody {
		C.cpBodyActivate(b.c())
	}
}

// ActivateStatic wakes up any sleeping or idle bodies touching a static body.
func (b Body) ActivateStatic(s Shape) {
	C.cpBodyActivateStatic(b.c(), s.c())
}

// Angle returns the rotation of the body around it's center of gravity in radians.
func (b Body) Angle() float64 {
	return float64(C.cpBodyGetAngle(b.c()))
}

// SetAngle sets the rotation of the body around it's center of gravity in radians.
func (b Body) SetAngle(angle float64) {
	C.cpBodySetAngle(b.c(), C.cpFloat(angle))
}

// AngularVelocity returns the angular velocity of the body around it's center
// of gravity in radians/second.
func (b Body) AngularVelocity() float64 {
	return float64(C.cpBodyGetAngVel(b.c()))
}

// SetAngularVelocity sets the angular velocity of the body around it's center
// of gravity in radians/second.
func (b Body) SetAngleularVelocity(vel float64) {
	C.cpBodySetAngVel(b.c(), C.cpFloat(vel))
}

// AngularVelocityLimit returns the maximum rotational rate (in radians/second) allowed when updating
// the angular velocity.
func (b Body) AngularVelocityLimit() float64 {
	return float64(C.cpBodyGetAngVelLimit(b.c()))
}

// SetAngularVelocityLimit sets the maximum rotational rate (in radians/second) allowed when updating
// the angular velocity.
func (b Body) SetAngularVelocityLimit(limit float64) {
	C.cpBodySetAngVelLimit(b.c(), C.cpFloat(limit))
}

// ApplyForce applies a force (in world coordinates) to the body at a point relative
// to the center of gravity (also in world coordinates).
func (b Body) ApplyForce(f, r Vect) {
	C.cpBodyApplyForce(b.c(), f.c(), r.c())
}

// ApplyImpulse applies an impulse (in world coordinates) to the body at a point relative
// to the center of gravity (also in world coordinates).
func (b Body) ApplyImpulse(j, r Vect) {
	C.cpBodyApplyImpulse(b.c(), j.c(), r.c())
}

// BodyNew creates a new body.
func BodyNew(m, i float64) Body {
	b := cpBody(C.cpBodyNew(C.cpFloat(m), C.cpFloat(i)))
	bodyDataMap[b] = &bodyData{}
	return b
}

// BodyStaticNew creates a new static body.
func BodyStaticNew() Body {
	b := cpBody(C.cpBodyNewStatic())
	bodyDataMap[b] = &bodyData{}
	return b
}

// EachArbiter calls a callback function once for each arbiter which is currently
// active on the body.
func (b Body) EachArbiter(iter func(Body, Arbiter)) {
	p := unsafe.Pointer(&iter)
	C.body_each_arbiter(b.c(), p)
}

// EachConstraint calls a callback function once for each constraint attached
// to the body and added to the space.
func (b Body) EachConstraint(iter func(Body, Constraint)) {
	p := unsafe.Pointer(&iter)
	C.body_each_constraint(b.c(), p)
}

// EachShape calls a callback function once for each shape attached to the body
// and added to the space.
func (b Body) EachShape(iter func(Body, Shape)) {
	p := unsafe.Pointer(&iter)
	C.body_each_shape(b.c(), p)
}

// Free removes a body.
func (b Body) Free() {
	delete(bodyDataMap, b)
	C.cpBodyFree(b.c())
}

// Force returns the force acting on the rigid body's center of gravity.
func (b Body) Force() Vect {
	return cpVect(C.cpBodyGetForce(b.c()))
}

// SetForce sets the force acting on the rigid body's center of gravity.
func (b Body) SetForce(force Vect) {
	C.cpBodySetForce(b.c(), force.c())
}

// IsRogue returns true if the body has not been added to a space.
func (b Body) IsRogue() bool {
	return cpBool(C.cpBodyIsRogue(b.c()))
}

// IsSleeping returns true if the body is sleeping.
func (b Body) IsSleeping() bool {
	return cpBool(C.cpBodyIsSleeping(b.c()))
}

// IsStatic returns true if the body is static.
func (b Body) IsStatic() bool {
	return cpBool(C.cpBodyIsStatic(b.c()))
}

// KineticEnergy returns the kinetic energy of a body.
func (b Body) KineticEnergy() float64 {
	return float64(C.cpBodyKineticEnergy(b.c()))
}

// LocalToWorld converts body relative/local coordinates to absolute/world coordinates.
func (b Body) LocalToWorld(v Vect) Vect {
	return cpVect(C.cpBodyLocal2World(b.c(), v.c()))
}

// Mass returns the mass of the body.
func (b Body) Mass() float64 {
	return float64(C.cpBodyGetMass(b.c()))
}

// SetMass sets the mass of the body.
func (b Body) SetMass(mass float64) {
	C.cpBodySetMass(b.c(), C.cpFloat(mass))
}

// Moment returns a moment of intertia of the body.
func (b Body) Moment() float64 {
	return float64(C.cpBodyGetMoment(b.c()))
}

// SetMoment sets the moment of the body.
func (b Body) SetMoment(moment float64) {
	C.cpBodySetMoment(b.c(), C.cpFloat(moment))
}

// Position returns the position of the rigid body's center of gravity.
func (b Body) Position() Vect {
	return cpVect(C.cpBodyGetPos(b.c()))
}

// SetPosition sets the position of the body.
func (b Body) SetPosition(pos Vect) {
	C.cpBodySetPos(b.c(), pos.c())
}

// ResetForces sets the forces and torque of a body to zero.
func (b Body) ResetForces() {
	C.cpBodyResetForces(b.c())
}

// Rotation returns the cached unit length vector representing the angle of the body.
func (b Body) Rotation() Vect {
	return cpVect(C.cpBodyGetRot(b.c()))
}

// SetPositionFunc sets a function that is called to integrate the body's position.
func (b Body) SetPositionFunc(f func(b Body, dt float64)) {
	bodyDataMap[b].positionFunc = f
	C.body_set_position_func(b.c(), boolToC(f != nil))
}

// SetVelocityFunc sets a function that is called to integrate the body's velocity.
func (b Body) SetVelocityFunc(f func(b Body, gravity Vect, damping, dt float64)) {
	bodyDataMap[b].velocityFunc = f
	C.body_set_velocity_func(b.c(), boolToC(f != nil))
}

// Sleep forces a body to fall asleep immediately.
func (b Body) Sleep() {
	C.cpBodySleep(b.c())
}

// SleepWithGroup forces a body to fall asleep immediately along with other bodies in a group.
func (b Body) SleepWithGroup(g Body) {
	C.cpBodySleepWithGroup(b.c(), g.c())
}

// Space returns space the body was added to or 0 if the body doesn't belong to any space.
func (b Body) Space() Space {
	return cpSpace(C.cpBodyGetSpace(b.c()))
}

// String converts a body to a human-readable string.
func (b Body) String() string {
	return fmt.Sprintf("(Body)%+v", b.c())
}

// Torque returns the torque applied to the body around it's center of gravity.
func (b Body) Torque() float64 {
	return float64(C.cpBodyGetTorque(b.c()))
}

// SetTorque sets the torque applied to the body around it's center of gravity.
func (b Body) SetTorque(torq float64) {
	C.cpBodySetTorque(b.c(), C.cpFloat(torq))
}

// UpdatePosition is a default function that is called to integrate the body's position.
func (b Body) UpdatePosition(dt float64) {
	C.cpBodyUpdatePosition(b.c(), C.cpFloat(dt))
}

// UpdateVelocity is a default function that is called to integrate the body's velocity.
func (b Body) UpdateVelocity(gravity Vect, damping float64, dt float64) {
	C.cpBodyUpdateVelocity(b.c(), gravity.c(), C.cpFloat(damping), C.cpFloat(dt))
}

// UserData returns user defined data.
func (b Body) UserData() interface{} {
	return bodyDataMap[b].userData
}

// SetUserData sets user definable data pointer.
// Generally this points to your the game object so you can access it
// when given a Body reference in a callback.
func (b Body) SetUserData(data interface{}) {
	bodyDataMap[b].userData = data
}

// Velocity returns the velocity of the rigid body's center of gravity.
func (b Body) Velocity() Vect {
	return cpVect(C.cpBodyGetVel(b.c()))
}

// SetVelocity sets the velocity of the body.
func (b Body) SetVelocity(vel Vect) {
	C.cpBodySetVel(b.c(), vel.c())
}

// VelocityAtLocalPoint returns the velocity on a body (in world units) at a point
// on the body in local coordinates.
func (b Body) VelocityAtLocalPoint(point Vect) Vect {
	return cpVect(C.cpBodyGetVelAtLocalPoint(b.c(), point.c()))
}

// VelocityAtWorldPoint returns the velocity on a body (in world units) at a point
// on the body in world coordinates.
func (b Body) VelocityAtWorldPoint(point Vect) Vect {
	return cpVect(C.cpBodyGetVelAtWorldPoint(b.c(), point.c()))
}

// VelocityLimit returns the maximum velocity allowed when updating the velocity.
func (b Body) VelocityLimit() float64 {
	return float64(C.cpBodyGetVelLimit(b.c()))
}

// SetVelocityLimit sets the maximum velocity allowed when updating the velocity.
func (b Body) SetVelocityLimit(limit float64) {
	C.cpBodySetVelLimit(b.c(), C.cpFloat(limit))
}

// WorldToLocal converts body absolute/world coordinates to relative/local coordinates.
func (b Body) WorldToLocal(v Vect) Vect {
	return cpVect(C.cpBodyWorld2Local(b.c(), v.c()))
}

// addToSpace adds a body to space.
func (b Body) addToSpace(s Space) {
	s.AddBody(b)
}

// c converts Body to c.cpBody pointer.
func (b Body) c() *C.cpBody {
	return (*C.cpBody)(unsafe.Pointer(b))
}

// containedInSpace returns true if the body is in the space.
func (b Body) containedInSpace(s Space) bool {
	return cpBool(C.cpSpaceContainsBody(s.c(), b.c()))
}

// cpBody converts C.cpBody pointer to Body.
func cpBody(b *C.cpBody) Body {
	return Body(unsafe.Pointer(b))
}

//export eachArbiter_body
func eachArbiter_body(b *C.cpBody, a *C.cpArbiter, p unsafe.Pointer) {
	f := *(*func(Body, Arbiter))(p)
	f(cpBody(b), cpArbiter(a))
}

//export eachConstraint_body
func eachConstraint_body(b *C.cpBody, c *C.cpConstraint, p unsafe.Pointer) {
	f := *(*func(Body, Constraint))(p)
	f(cpBody(b), cpConstraint(c))
}

//export eachShape_body
func eachShape_body(b *C.cpBody, sh *C.cpShape, p unsafe.Pointer) {
	f := *(*func(Body, Shape))(p)
	f(cpBody(b), cpShape(sh))
}

// removeFromSpace removes a body from space.
func (b Body) removeFromSpace(s Space) {
	s.RemoveBody(b)
}

//export updatePosition
func updatePosition(b *C.cpBody, dt C.cpFloat) {
	d := bodyDataMap[cpBody(b)]
	d.positionFunc(cpBody(b), float64(dt))
}

//export updateVelocity
func updateVelocity(b *C.cpBody, gravity C.cpVect, damping, dt C.cpFloat) {
	d := bodyDataMap[cpBody(b)]
	d.velocityFunc(cpBody(b), cpVect(gravity), float64(damping), float64(dt))
}
