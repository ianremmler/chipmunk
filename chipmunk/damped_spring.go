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

type DampedSpring struct {
  constraintBase
}

func DampedSpringNew(
  a, b Body,
  anchr1, anchr2 Vect,
  restLength, stiffness, damping float64) DampedSpring {
  return DampedSpring{
    constraintBase{
      C.cpDampedSpringNew(
        a.c(),
        b.c(),
        anchr1.c(),
        anchr2.c(),
        C.cpFloat(restLength),
        C.cpFloat(stiffness),
        C.cpFloat(damping))}}
}

/////////////////////////////////////////////////////////////////////////////

func (c DampedSpring) Anchr1() Vect {
  return cpVect(C.cpDampedSpringGetAnchr1(c.ct))
}

func (c DampedSpring) Anchr2() Vect {
  return cpVect(C.cpDampedSpringGetAnchr2(c.ct))
}

func (c DampedSpring) RestLength() float64 {
  return float64(C.cpDampedSpringGetRestLength(c.ct))
}

func (c DampedSpring) Stiffness() float64 {
  return float64(C.cpDampedSpringGetStiffness(c.ct))
}

func (c DampedSpring) Damping() float64 {
  return float64(C.cpDampedSpringGetDamping(c.ct))
}

/////////////////////////////////////////////////////////////////////////////

func (c DampedSpring) SetAnchr1(v Vect) {
  C.cpDampedSpringSetAnchr1(c.ct, v.c())
}

func (c DampedSpring) SetAnchr2(v Vect) {
  C.cpDampedSpringSetAnchr2(c.ct, v.c())
}

func (c DampedSpring) SetRestLength(restLength float64) {
  C.cpDampedSpringSetRestLength(c.ct, C.cpFloat(restLength))
}

func (c DampedSpring) SetStiffness(stiffness float64) {
  C.cpDampedSpringSetStiffness(c.ct, C.cpFloat(stiffness))
}

func (c DampedSpring) SetDamping(damping float64) {
  C.cpDampedSpringSetDamping(c.ct, C.cpFloat(damping))
}
