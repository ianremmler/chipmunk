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
  "unsafe"
)

type Arbiter struct {
  a *C.cpArbiter
}

type CollisionBegin func(arb Arbiter, space Space) bool
type CollisionPreSolv func(arb Arbiter, space Space) bool
type CollisionPostSolve func(arb Arbiter, space Space)
type CollisionSeparate func(arb Arbiter, space Space)

func (a Arbiter) Elasticity() float64 {
  return float64(C.cpArbiterGetElasticity(a.a))
}

func (a Arbiter) Friction() float64 {
  return float64(C.cpArbiterGetFriction(a.a))
}

func (a Arbiter) SurfaceVelocity() Vect {
  return cpVect(C.cpArbiterGetSurfaceVelocity(a.a))
}

func (a Arbiter) SetElasticity(e float64) {
  C.cpArbiterSetElasticity(a.a, C.cpFloat(e))
}

func (a Arbiter) SetFriction(f float64) {
  C.cpArbiterSetFriction(a.a, C.cpFloat(f))
}

func (a Arbiter) SetSurfaceVelocity(v Vect) {
  C.cpArbiterSetSurfaceVelocity(a.a, v.c())
}

func (a Arbiter) TotalKE() float64 {
  return float64(C.cpArbiterTotalKE(a.a))
}

func (a Arbiter) Ignore() {
  C.cpArbiterIgnore(a.a)
}

func (arb Arbiter) Shapes() (Shape, Shape) {
  var a, b *C.cpShape
  C.cpArbiterGetShapes(arb.a, (**C.cpShape)(unsafe.Pointer(&a)), (**C.cpShape)(unsafe.Pointer(&b)))
  return cpShape(a), cpShape(b)
}

func (arb Arbiter) Bodies() (Body, Body) {
  var a, b *C.cpBody
  C.cpArbiterGetBodies(arb.a, (**C.cpBody)(unsafe.Pointer(&a)), (**C.cpBody)(unsafe.Pointer(&b)))
  return cpBody(a), cpBody(b)
}

func (a Arbiter) IsFirstContact() bool {
  return cpBool(C.cpArbiterIsFirstContact(a.a))
}

func (a Arbiter) Count() int {
  return int(C.cpArbiterGetCount(a.a))
}

func (a Arbiter) Normal(i int) Vect {
  return cpVect(C.cpArbiterGetNormal(a.a, C.int(i)))
}

func (a Arbiter) Point(i int) Vect {
  return cpVect(C.cpArbiterGetPoint(a.a, C.int(i)))
}

func (a Arbiter) Depth(i int) float64 {
  return float64(C.cpArbiterGetDepth(a.a, C.int(i)))
}
