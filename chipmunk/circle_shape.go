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

// CircleShape is a circle shape type.
type CircleShape struct {
  shapeBase
}

// CircleShapeNew creates a new circle shape.
func CircleShapeNew(body Body, radius float64, offset Vect) Shape {
  s := C.cpCircleShapeNew(body.b, C.cpFloat(radius), offset.c())
  return CircleShape{shapeBase{s}}
}

// Offset returns the offset from the center of gravity.
func (s CircleShape) Offset() Vect {
  return cpVect(C.cpCircleShapeGetOffset(s.s))
}

// Radius returns the radius of the circle.
func (s CircleShape) Radius() float64 {
  return float64(C.cpCircleShapeGetRadius(s.s))
}

// String converts a circle shape to a human-readable string.
func (s CircleShape) String() string {
  return "circle shape"
}

// Local Variables:
// indent-tabs-mode: nil
// tab-width: 2
// End:
// ex: set tabstop=2 shiftwidth=2 expandtab:
