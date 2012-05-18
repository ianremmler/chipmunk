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

type constraintBase struct {
  ct *C.cpConstraint
}

// Constraint is a type of object which is used to connect two bodies together.
type Constraint interface {
  ContainedInSpace(Space) bool
  Free()
  c() *C.cpConstraint
  A() Body
  B() Body
  Space() Space
  MaxForce() float64
  ErrorBias() float64
  MaxBias() float64
  UserData() interface{}
  SetMaxForce(float64)
  SetErrorBias(float64)
  SetMaxBias(float64)
  SetUserData(interface{})
  ActivateBodies()
  Impulse() float64
}

// Free frees the constraint.
func (c constraintBase) Free() {
  C.cpConstraintFree(c.ct)
}

func (c constraintBase) c() *C.cpConstraint {
  return c.ct
}

/////////////////////////////////////////////////////////////////////////////

// A returns the first body the constraint controls.
func (c constraintBase) A() Body {
  return cpBody(C.cpConstraintGetA(c.ct))
}

// B returns the second body the constraint controls.
func (c constraintBase) B() Body {
  return cpBody(C.cpConstraintGetB(c.ct))
}

// Space returns space the constraint was added to or nil if the constraint
// doesn't belong to any space.
func (c constraintBase) Space() Space {
  return cpSpace(C.cpConstraintGetSpace(c.ct))
}

// MaxForce returns the maximum force this constraint is allowed to use.
func (c constraintBase) MaxForce() float64 {
  return float64(C.cpConstraintGetMaxForce(c.ct))
}

// ErrorBias returns the rate at which joint error is corrected.
func (c constraintBase) ErrorBias() float64 {
  return float64(C.cpConstraintGetErrorBias(c.ct))
}

// MaxBias returns the maximum rate (speed) that a joint can be corrected at.
func (c constraintBase) MaxBias() float64 {
  return float64(C.cpConstraintGetMaxBias(c.ct))
}

// UserData returns user defined data.
func (c constraintBase) UserData() interface{} {
  return cpData(C.cpConstraintGetUserData(c.ct))
}

/////////////////////////////////////////////////////////////////////////////

// SetMaxForce sets the maximum force this constraint is allowed to use (defalts to infinity).
// This allows joints to be pulled apart if too much force is applied to them.
// It also allows you to use constraints as force or friction generators for controlling bodies.
func (c constraintBase) SetMaxForce(f float64) {
  C.cpConstraintSetMaxForce(c.ct, C.cpFloat(f))
}

// SetErrorBias sets the rate at which joint error is corrected.
// Defaults to math.Pow(1.0 - 0.1, 60.0) meaning that it will correct 10% of the error
// every 1/60th of a second.
func (c constraintBase) SetErrorBias(b float64) {
  C.cpConstraintSetErrorBias(c.ct, C.cpFloat(b))
}

// SetMaxBias sets the maximum rate (speed) that a joint can be corrected at (defaults to infinity).
func (c constraintBase) SetMaxBias(b float64) {
  C.cpConstraintSetMaxBias(c.ct, C.cpFloat(b))
}

// SetUserData sets user definable data pointer.
// Generally this points to your the game object so you can access it
// when given a Constraint reference in a callback.
func (c constraintBase) SetUserData(data interface{}) {
  C.cpConstraintSetUserData(c.ct, dataToC(data))
}

/////////////////////////////////////////////////////////////////////////////

// ContainedInSpace returns true if the constraint is in the space.
func (c constraintBase) ContainedInSpace(s Space) bool {
  return cpBool(C.cpSpaceContainsConstraint(s.c(), c.ct))
}

// ActivateBodies calls Activate() on bodies the constraint controls.
func (c constraintBase) ActivateBodies() {
  c.A().Activate()
  c.B().Activate()
}

// Impulse returns the last impulse applied by this constraint.
func (c constraintBase) Impulse() float64 {
  return float64(C.cpConstraintGetImpulse(c.ct))
}

/////////////////////////////////////////////////////////////////////////////

var (
  pinJointClass           = C.cpPinJointGetClass()
  slideJointClass         = C.cpSlideJointGetClass()
  pivotJointClass         = C.cpPivotJointGetClass()
  grooveJointClass        = C.cpGrooveJointGetClass()
  dampedSpringClass       = C.cpDampedSpringGetClass()
  dampedRotarySpringClass = C.cpDampedRotarySpringGetClass()
  rotaryLimitJointClass   = C.cpRotaryLimitJointGetClass()
  ratchetJointClass       = C.cpRatchetJointGetClass()
  gearJointClass          = C.cpGearJointGetClass()
  simpleMotorClass        = C.cpSimpleMotorGetClass()
)

func cpConstraint(ct *C.cpConstraint) Constraint {
  if nil == ct {
    return nil
  }

  c := constraintBase{ct}

  switch c.ct.klass_private {
  case pinJointClass:
    return PinJoint{c}
  case slideJointClass:
    return SlideJoint{c}
  case pivotJointClass:
    return PivotJoint{c}
  case grooveJointClass:
    return GrooveJoint{c}
  case dampedSpringClass:
    return DampedSpring{c}
  case dampedRotarySpringClass:
    return DampedRotarySpring{c}
  case rotaryLimitJointClass:
    return RotaryLimitJoint{c}
  case ratchetJointClass:
    return RatchetJoint{c}
  case gearJointClass:
    return GearJoint{c}
  case simpleMotorClass:
    return SimpleMotor{c}
  }

  panic("unknown constraint class in cpConstraint")
}

// Local Variables:
// indent-tabs-mode: nil
// tab-width: 2
// End:
// ex: set tabstop=2 shiftwidth=2 expandtab:
