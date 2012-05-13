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

// PolyShape is a polygon shape type.
type PolyShape struct {
  shapeBase
}

// PolyShapeNew creates a new polygon shape.
func PolyShapeNew(b Body, verts []Vect, offset Vect) Shape {
  v := (*C.cpVect)(unsafe.Pointer(&verts[0]))
  return PolyShape{shapeBase{s: C.cpPolyShapeNew(b.c(), C.int(len(verts)), v, offset.c())}}
}

// BoxShapeNew creates a new box shape.
func BoxShapeNew(b Body, width, height float64) Shape {
  return PolyShape{shapeBase{s: C.cpBoxShapeNew(b.c(), C.cpFloat(width), C.cpFloat(height))}}
}

// BoxShapeNew2 creates a new box shape.
func BoxShapeNew2(b Body, box BB) Shape {
  return PolyShape{shapeBase{s: C.cpBoxShapeNew2(b.c(), box.c())}}
}

func (s PolyShape) String() string {
  return "poly shape"
}

// PolyValidate returns true if a set of vertexes is convex and has a clockwise winding.
func PolyValidate(verts []Vect) bool {
  v := (*C.cpVect)(unsafe.Pointer(&verts[0]))
  return cpBool(C.cpPolyValidate(v, C.int(len(verts))))
}

// NumVerts returns the number of vertices in a polygon shape.
func (s PolyShape) NumVerts() int {
  return int(C.cpPolyShapeGetNumVerts(s.s))
}

// Vert returns a specific vertex of a polygon shape.
func (s PolyShape) Vert(idx int) Vect {
  return cpVect(C.cpPolyShapeGetVert(s.s, C.int(idx)))
}