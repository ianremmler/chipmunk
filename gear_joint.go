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
import "C"

////////////////////////////////////////////////////////////////////////////////

// GearJoint constrains the rotational speed of one body to another.
// A ratio of 1.0 will lock the rotation of two bodies together, and negative ratios
// will cause them to spin in opposite directions.
// You can also use gear joints as rotary servos by setting MaxForce and MaxBias to
// finite values and changing the Phase property.
type GearJoint struct {
	constraintBase
}

////////////////////////////////////////////////////////////////////////////////

// GearJointNew creates a new gear joint.
func GearJointNew(a, b Body, phase, ratio float64) GearJoint {
	c := C.cpGearJointNew(a.c(), b.c(), C.cpFloat(phase), C.cpFloat(ratio))
	return GearJoint{cpConstraintBaseNew(c)}
}

// Phase returns the angular offset in radians.
func (c GearJoint) Phase() float64 {
	return float64(C.cpGearJointGetPhase(c.c()))
}

// Ratio returns the ratio of the rotational speeds.
func (c GearJoint) Ratio() float64 {
	return float64(C.cpGearJointGetRatio(c.c()))
}

// SetPhase sets the angular offset in radians.
func (c GearJoint) SetPhase(m float64) {
	C.cpGearJointSetPhase(c.c(), C.cpFloat(m))
}

// SetRatio sets the ratio of the rotational speeds.
func (c GearJoint) SetRatio(m float64) {
	C.cpGearJointSetRatio(c.c(), C.cpFloat(m))
}
