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

type DampedRotarySpring struct {
  constraintBase
}

func DampedRotarySpringNew(a, b Body, restAngle, stiffness, damping float64) DampedRotarySpring {
  return DampedRotarySpring{
    constraintBase{
      C.cpDampedRotarySpringNew(
        a.c(),
        b.c(),
        C.cpFloat(restAngle),
        C.cpFloat(stiffness),
        C.cpFloat(damping))}}
}

/////////////////////////////////////////////////////////////////////////////

func (j DampedRotarySpring) RestAngle() float64 {
  return float64(C.cpDampedRotarySpringGetRestAngle(j.ct))
}

func (j DampedRotarySpring) Stiffness() float64 {
  return float64(C.cpDampedRotarySpringGetStiffness(j.ct))
}

func (j DampedRotarySpring) Damping() float64 {
  return float64(C.cpDampedRotarySpringGetDamping(j.ct))
}

/////////////////////////////////////////////////////////////////////////////

func (j DampedRotarySpring) SetRestAngle(restAngle float64) {
  C.cpDampedRotarySpringSetRestAngle(j.ct, C.cpFloat(restAngle))
}

func (j DampedRotarySpring) SetStiffness(stiffness float64) {
  C.cpDampedRotarySpringSetStiffness(j.ct, C.cpFloat(stiffness))
}

func (j DampedRotarySpring) SetDamping(damping float64) {
  C.cpDampedRotarySpringSetDamping(j.ct, C.cpFloat(damping))
}
