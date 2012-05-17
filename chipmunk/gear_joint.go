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

type GearJoint struct {
  constraintBase
}

func GearJointNew(a, b Body, phase, ratio float64) GearJoint {
  return GearJoint{
    constraintBase{
      C.cpGearJointNew(
        a.c(),
        b.c(),
        C.cpFloat(phase),
        C.cpFloat(ratio))}}
}

/////////////////////////////////////////////////////////////////////////////

func (c GearJoint) Phase() float64 {
  return float64(C.cpGearJointGetPhase(c.ct))
}

func (c GearJoint) Ratio() float64 {
  return float64(C.cpGearJointGetRatio(c.ct))
}

/////////////////////////////////////////////////////////////////////////////

func (c GearJoint) SetPhase(m float64) {
  C.cpGearJointSetPhase(c.ct, C.cpFloat(m))
}

func (c GearJoint) SetRatio(m float64) {
  C.cpGearJointSetRatio(c.ct, C.cpFloat(m))
}
