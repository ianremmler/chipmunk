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

type GrooveJoint struct {
  constraintBase
}

func GrooveJointNew(a, b Body, groove_a, groove_b, anchr2 Vect) GrooveJoint {
  return GrooveJoint{ constraintBase{ C.cpGrooveJointNew(a.c(), b.c(), groove_a.c(), groove_b.c(), anchr2.c()) } }
}

/////////////////////////////////////////////////////////////////////////////

func (j GrooveJoint) GrooveA() Vect {
  return cpVect(C.cpGrooveJointGetGrooveA(j.ct))
}

func (j GrooveJoint) GrooveB() Vect {
  return cpVect(C.cpGrooveJointGetGrooveB(j.ct))
}

func (j GrooveJoint) Anchr2() Vect {
  return cpVect(C.cpGrooveJointGetAnchr2(j.ct))
}

/////////////////////////////////////////////////////////////////////////////

func (j GrooveJoint) SetGrooveA(v Vect) {
  C.cpGrooveJointSetGrooveA(j.ct, v.c())
}

func (j GrooveJoint) SetGrooveB(v Vect) {
  C.cpGrooveJointSetGrooveB(j.ct, v.c())
}

func (j GrooveJoint) SetAnchr2(v Vect) {
  C.cpGrooveJointSetAnchr2(j.ct, v.c())
}
