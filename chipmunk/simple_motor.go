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

// SimpleMotor makes two objects spin relative to each other.
// They are most often used with the MaxForce property set to a finite value.
type SimpleMotor struct {
  constraintBase
}

// SimpleMotorNew creates a new simple motor.
func SimpleMotorNew(a, b Body, rate float64) SimpleMotor {
  return SimpleMotor{
    constraintBase{
      C.cpSimpleMotorNew(
        a.c(),
        b.c(),
        C.cpFloat(rate))}}
}

/////////////////////////////////////////////////////////////////////////////

// Rate returns the relative rotation speed of the two bodies in radians per second.
func (c SimpleMotor) Rate() float64 {
  return float64(C.cpSimpleMotorGetRate(c.ct))
}

/////////////////////////////////////////////////////////////////////////////

// SetRate sets the relative rotation speed of the two bodies in radians per second.
func (c SimpleMotor) SetRate(m float64) {
  C.cpSimpleMotorSetRate(c.ct, C.cpFloat(m))
}
