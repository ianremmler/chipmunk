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

type Constraint interface {
  ContainedInSpace(Space) bool
  Destroy()
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

func (c constraintBase) Destroy() {
  C.cpConstraintDestroy(c.ct)
}

func (c constraintBase) c() *C.cpConstraint {
  return c.ct
}

/////////////////////////////////////////////////////////////////////////////

func (c constraintBase) A() Body {
  return cpBody(C.cpConstraintGetA(c.ct))
}

func (c constraintBase) B() Body {
  return cpBody(C.cpConstraintGetB(c.ct))
}

func (c constraintBase) Space() Space {
  return cpSpace(C.cpConstraintGetSpace(c.ct))
}

func (c constraintBase) MaxForce() float64 {
  return float64(C.cpConstraintGetMaxForce(c.ct))
}

func (c constraintBase) ErrorBias() float64 {
  return float64(C.cpConstraintGetErrorBias(c.ct))
}

func (c constraintBase) MaxBias() float64 {
  return float64(C.cpConstraintGetMaxBias(c.ct))
}

// UserData returns user defined data.
func (c constraintBase) UserData() interface{} {
  return cpData(C.cpConstraintGetUserData(c.ct))
}

/////////////////////////////////////////////////////////////////////////////

func (c constraintBase) SetMaxForce(f float64) {
  C.cpConstraintSetMaxForce(c.ct, C.cpFloat(f))
}

func (c constraintBase) SetErrorBias(b float64) {
  C.cpConstraintSetErrorBias(c.ct, C.cpFloat(b))
}

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

func (c constraintBase) ContainedInSpace(s Space) bool {
  return cpBool(C.cpSpaceContainsConstraint(s.s, c.ct))
}

func (c constraintBase) ActivateBodies() {
  c.A().Activate()
  c.B().Activate()
}

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
