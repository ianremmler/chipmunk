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
// #include <chipmunk_unsafe.h>
import "C"

import (
  "fmt"
)

////////////////////////////////////////////////////////////////////////////////

// SegmentShape is a a beveled (rounded) segment shape.
type SegmentShape struct {
  shapeBase
}

////////////////////////////////////////////////////////////////////////////////

// A returns the start of the segment shape.
func (s SegmentShape) A() Vect {
  return cpVect(C.cpSegmentShapeGetA(s.c()))
}

// B returns the end of the segment shape.
func (s SegmentShape) B() Vect {
  return cpVect(C.cpSegmentShapeGetB(s.c()))
}

// Normal returns the normal of the segment shape.
func (s SegmentShape) Normal() Vect {
  return cpVect(C.cpSegmentShapeGetNormal(s.c()))
}

// Radius returns the beveling radius of the segment shape.
func (s SegmentShape) Radius() float64 {
  return float64(C.cpSegmentShapeGetRadius(s.c()))
}

// SegmentShapeNew creates a new segment shape.
func SegmentShapeNew(body Body, a, b Vect, radius float64) SegmentShape {
  s := C.cpSegmentShapeNew(body.c(), a.c(), b.c(), C.cpFloat(radius))
  return SegmentShape{cpshape(s)}
}

// SetEndpoints sets the endpoints of a segment shape.
// NOTE: this is unsafe function.
func (s SegmentShape) SetEndpoints(a, b Vect) {
  C.cpSegmentShapeSetEndpoints(s.c(), a.c(), b.c())
}

// SetNeighbors FIXME TODO OMG WTF
func (s SegmentShape) SetNeighbors(prev, next Vect) {
  C.cpSegmentShapeSetNeighbors(s.c(), prev.c(), next.c())
}

// SetRadius sets the radius of a segment shape.
// NOTE: this is unsafe function.
func (s SegmentShape) SetRadius(radius float64) {
  C.cpSegmentShapeSetRadius(s.c(), C.cpFloat(radius))
}

// String converts a segment shape to a human-readable string.
func (s SegmentShape) String() string {
  return fmt.Sprintf("(SegmentShape)%+v", s.c())
}

// Local Variables:
// indent-tabs-mode: nil
// tab-width: 2
// End:
// ex: set tabstop=2 shiftwidth=2 expandtab:
