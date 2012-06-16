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

// #include <chipmunk.h>
// #include "constraint.h"
import "C"

import (
  . "unsafe"
)

////////////////////////////////////////////////////////////////////////////////

// Constraint is a type of object which is used to connect two bodies together.
type Constraint interface {
  A() Body
  ActivateBodies()
  B() Body
  ErrorBias() float64
  Free()
  Impulse() float64
  MaxBias() float64
  MaxForce() float64
  SetErrorBias(float64)
  SetMaxBias(float64)
  SetMaxForce(float64)
  SetUserData(interface{})
  Space() Space
  UserData() interface{}
  c() *C.cpConstraint
}

type constraintBase uintptr

type constraintData struct {
  postSolveFunc func(Constraint, Space)
  preSolveFunc  func(Constraint, Space)
  userData      interface{}
}

////////////////////////////////////////////////////////////////////////////////

var (
  constraintDataMap = make(map[constraintBase]*constraintData)
)

////////////////////////////////////////////////////////////////////////////////

var (
  dampedRotarySpringClass = C.cpDampedRotarySpringGetClass()
  dampedSpringClass       = C.cpDampedSpringGetClass()
  gearJointClass          = C.cpGearJointGetClass()
  grooveJointClass        = C.cpGrooveJointGetClass()
  pinJointClass           = C.cpPinJointGetClass()
  pivotJointClass         = C.cpPivotJointGetClass()
  ratchetJointClass       = C.cpRatchetJointGetClass()
  rotaryLimitJointClass   = C.cpRotaryLimitJointGetClass()
  simpleMotorClass        = C.cpSimpleMotorGetClass()
  slideJointClass         = C.cpSlideJointGetClass()
)

////////////////////////////////////////////////////////////////////////////////

// A returns the first body the constraint controls.
func (c constraintBase) A() Body {
  return cpBody(C.cpConstraintGetA(c.c()))
}

// ActivateBodies calls Activate() on bodies the constraint controls.
func (c constraintBase) ActivateBodies() {
  c.A().Activate()
  c.B().Activate()
}

// B returns the second body the constraint controls.
func (c constraintBase) B() Body {
  return cpBody(C.cpConstraintGetB(c.c()))
}

// ErrorBias returns the rate at which joint error is corrected.
func (c constraintBase) ErrorBias() float64 {
  return float64(C.cpConstraintGetErrorBias(c.c()))
}

// Free frees the constraint.
func (c constraintBase) Free() {
  constraintDataMap[c] = nil
  C.cpConstraintFree(c.c())
}

// Impulse returns the last impulse applied by this constraint.
func (c constraintBase) Impulse() float64 {
  return float64(C.cpConstraintGetImpulse(c.c()))
}

// MaxForce returns the maximum force this constraint is allowed to use.
func (c constraintBase) MaxForce() float64 {
  return float64(C.cpConstraintGetMaxForce(c.c()))
}

// MaxBias returns the maximum rate (speed) that a joint can be corrected at.
func (c constraintBase) MaxBias() float64 {
  return float64(C.cpConstraintGetMaxBias(c.c()))
}

// SetErrorBias sets the rate at which joint error is corrected.
// Defaults to math.Pow(1.0 - 0.1, 60.0) meaning that it will correct 10% of the error
// every 1/60th of a second.
func (c constraintBase) SetErrorBias(b float64) {
  C.cpConstraintSetErrorBias(c.c(), C.cpFloat(b))
}

// SetMaxBias sets the maximum rate (speed) that a joint can be corrected at (defaults to infinity).
func (c constraintBase) SetMaxBias(b float64) {
  C.cpConstraintSetMaxBias(c.c(), C.cpFloat(b))
}

// SetMaxForce sets the maximum force this constraint is allowed to use (defalts to infinity).
// This allows joints to be pulled apart if too much force is applied to them.
// It also allows you to use constraints as force or friction generators for controlling bodies.
func (c constraintBase) SetMaxForce(f float64) {
  C.cpConstraintSetMaxForce(c.c(), C.cpFloat(f))
}

// SetPostSolveFunc sets a callback function type that gets called after solving a joint.
// Use the applied impulse to perform effects like breakable joints.
func (c constraintBase) SetPostSolveFunc(f func(c Constraint, s Space)) {
  constraintDataMap[c].postSolveFunc = f
  C.constraint_set_postsolve_func(c.c(), boolToC(f != nil))
}

// SetPreSolveFunc sets a callback function type that gets called before solving a joint.
// Animate your joint anchors, update your motor torque, etc.
func (c constraintBase) SetPreSolveFunc(f func(c Constraint, s Space)) {
  constraintDataMap[c].preSolveFunc = f
  C.constraint_set_presolve_func(c.c(), boolToC(f != nil))
}

// SetUserData sets user definable data pointer.
// Generally this points to your the game object so you can access it
// when given a Constraint reference in a callback.
func (c constraintBase) SetUserData(data interface{}) {
  constraintDataMap[c].userData = data
}

// Space returns space the constraint was added to or nil if the constraint
// doesn't belong to any space.
func (c constraintBase) Space() Space {
  return cpSpace(C.cpConstraintGetSpace(c.c()))
}

// UserData returns user defined data.
func (c constraintBase) UserData() interface{} {
  return constraintDataMap[c].userData
}

// addToSpace adds a constraint to space.
func (c constraintBase) addToSpace(s Space) {
  s.AddConstraint(c)
}

// c converts Constraint to c.cpConstraint pointer.
func (c constraintBase) c() *C.cpConstraint {
  return (*C.cpConstraint)(Pointer(c))
}

//export constraint_postsolve
func constraint_postsolve(c *C.cpConstraint, s *C.cpSpace) {
  constraint := cpConstraint(c)
  data := constraintDataMap[cpconstraint(c)]
  data.postSolveFunc(constraint, cpSpace(s))
}

//export constraint_presolve
func constraint_presolve(c *C.cpConstraint, s *C.cpSpace) {
  constraint := cpConstraint(c)
  data := constraintDataMap[cpconstraint(c)]
  data.preSolveFunc(constraint, cpSpace(s))
}

// containedInSpace returns true if the constraint is in the space.
func (c constraintBase) containedInSpace(s Space) bool {
  return cpBool(C.cpSpaceContainsConstraint(s.c(), c.c()))
}

// cpConstraint converts C.cpConstraint pointer to Constraint.
func cpConstraint(ct *C.cpConstraint) Constraint {
  if nil == ct {
    return nil
  }

  c := cpconstraint(ct)

  switch c.c().klass_private {
  case gearJointClass:
    return GearJoint{c}
  case grooveJointClass:
    return GrooveJoint{c}
  case dampedRotarySpringClass:
    return DampedRotarySpring{c}
  case dampedSpringClass:
    return DampedSpring{c}
  case pinJointClass:
    return PinJoint{c}
  case pivotJointClass:
    return PivotJoint{c}
  case simpleMotorClass:
    return SimpleMotor{c}
  case slideJointClass:
    return SlideJoint{c}
  case ratchetJointClass:
    return RatchetJoint{c}
  case rotaryLimitJointClass:
    return RotaryLimitJoint{c}
  }

  panic("unknown constraint class in cpConstraint")
}

// cpconstraint converts C.cpConstraint pointer to constraintBase.
func cpconstraint(ct *C.cpConstraint) constraintBase {
  return constraintBase(Pointer(ct))
}

// cpconstraint_new creates a new constraintBase out of C.cpConstraint pointer.
func cpconstraint_new(ct *C.cpConstraint) constraintBase {
  c := cpconstraint(ct)
  constraintDataMap[c] = &constraintData{}
  return c
}

// removeFromSpace removes a constraint from space.
func (c constraintBase) removeFromSpace(s Space) {
  s.RemoveConstraint(c)
}

// Local Variables:
// indent-tabs-mode: nil
// tab-width: 2
// End:
// ex: set tabstop=2 shiftwidth=2 expandtab:
