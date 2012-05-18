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

// SegmentShape is a a beveled (rounded) segment shape.
type SegmentShape struct {
  shapeBase
}

// SegmentShapeNew creates a new segment shape.
func SegmentShapeNew(body Body, a, b Vect, radius float64) Shape {
  s := C.cpSegmentShapeNew(body.b, a.c(), b.c(), C.cpFloat(radius))
  return SegmentShape{shapeBase{s}}
}

// SetNeighbors FIXME TODO OMG WTF
func (s SegmentShape) SetNeighbors(prev, next Vect) {
  C.cpSegmentShapeSetNeighbors(s.s, prev.c(), next.c())
}

// A returns the start of the segment shape.
func (s SegmentShape) A() Vect {
  return cpVect(C.cpSegmentShapeGetA(s.s))
}

// B returns the end of the segment shape.
func (s SegmentShape) B() Vect {
  return cpVect(C.cpSegmentShapeGetB(s.s))
}

// Normal returns the normal of the segment shape.
func (s SegmentShape) Normal() Vect {
  return cpVect(C.cpSegmentShapeGetNormal(s.s))
}

// Radius returns the beveling radius of the segment shape.
func (s SegmentShape) Radius() float64 {
  return float64(C.cpSegmentShapeGetRadius(s.s))
}

// String converts a segment shape to a human-readable string.
func (s SegmentShape) String() string {
  return "segment shape"
}

// Local Variables:
// indent-tabs-mode: nil
// tab-width: 2
// End:
// ex: set tabstop=2 shiftwidth=2 expandtab:
