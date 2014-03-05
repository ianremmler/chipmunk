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
// #include <chipmunk/chipmunk_unsafe.h>
import "C"

import (
	"fmt"
	"reflect"
	"unsafe"
)

////////////////////////////////////////////////////////////////////////////////

// PolyShape is a polygon shape type.
type PolyShape struct {
	shapeBase
}

////////////////////////////////////////////////////////////////////////////////

// BoxShapeNew creates a new box shape.
func BoxShapeNew(b Body, width, height float64) Shape {
	s := C.cpBoxShapeNew(b.c(), C.cpFloat(width), C.cpFloat(height))
	return PolyShape{cpshape(s)}
}

// BoxShapeNew2 creates a new box shape.
func BoxShapeNew2(b Body, box BB) Shape {
	s := C.cpBoxShapeNew2(b.c(), box.c())
	return PolyShape{cpshape(s)}
}

// BoxShapeNew3 creates a new box shape.
func BoxShapeNew3(b Body, box BB, radius float64) Shape {
	s := C.cpBoxShapeNew3(b.c(), box.c(), C.cpFloat(radius))
	return PolyShape{cpshape(s)}
}

// NumVerts returns the number of vertices in a polygon shape.
func (s PolyShape) NumVerts() int {
	return int(C.cpPolyShapeGetNumVerts(s.c()))
}

// PolyShapeNew creates a new polygon shape.
func PolyShapeNew(b Body, verts []Vect, offset Vect) PolyShape {
	v := (*C.cpVect)(unsafe.Pointer(&verts[0]))
	s := C.cpPolyShapeNew(b.c(), C.int(len(verts)), v, offset.c())
	return PolyShape{cpshape(s)}
}

// PolyShapeNew2 creates a new polygon shape.
func PolyShapeNew2(b Body, verts []Vect, offset Vect, radius float64) PolyShape {
	v := (*C.cpVect)(unsafe.Pointer(&verts[0]))
	s := C.cpPolyShapeNew2(b.c(), C.int(len(verts)), v, offset.c(), C.cpFloat(radius))
	return PolyShape{cpshape(s)}
}

// PolyValidate returns true if a set of vertexes is convex and has a clockwise winding.
func PolyValidate(verts []Vect) bool {
	v := (*C.cpVect)(unsafe.Pointer(&verts[0]))
	return cpBool(C.cpPolyValidate(v, C.int(len(verts))))
}

// Radius returns the vertex radius.
func (s PolyShape) Radius() float64 {
	return float64(C.cpPolyShapeGetRadius(s.c()))
}

// SetRadius sets the vertex radius.
func (s PolyShape) SetRadius(radius float64) {
	C.cpPolyShapeSetRadius(s.c(), C.cpFloat(radius))
}

// SetVerts sets the vertexes of a poly shape.
func (s PolyShape) SetVerts(verts []Vect, offset Vect) {
	v := (*C.cpVect)(unsafe.Pointer(&verts[0]))
	C.cpPolyShapeSetVerts(s.c(), C.int(len(verts)), v, offset.c())
}

// String converts a polygon shape to a human-readable string.
func (s PolyShape) String() string {
	return fmt.Sprintf("(PolyShape)%+v", s.c())
}

// VertLocal returns a specific vertex of a polygon shape (local coordinates).
func (s PolyShape) VertLocal(idx int) Vect {
	return cpVect(C.cpPolyShapeGetVert(s.c(), C.int(idx)))
}

// VertsWorld returns vertex positions (world coordinates).
func (s PolyShape) VertsWorld() []Vect {
	num := s.NumVerts()
	var verts []Vect

	vertsH := (*reflect.SliceHeader)((unsafe.Pointer(&verts)))
	vertsH.Cap = num
	vertsH.Len = num
	vertsH.Data = uintptr(unsafe.Pointer(((*C.cpPolyShape)(unsafe.Pointer(s.c())).tVerts)))

	return verts
}

// VertsWorldFloat64 returns vertex positions (world coordinates) as array of float64 values.
func (s PolyShape) VertsWorldFloat64() []float64 {
	num := s.NumVerts() * 2
	var verts []float64

	vertsH := (*reflect.SliceHeader)((unsafe.Pointer(&verts)))
	vertsH.Cap = num
	vertsH.Len = num
	vertsH.Data = uintptr(unsafe.Pointer(((*C.cpPolyShape)(unsafe.Pointer(s.c())).tVerts)))

	return verts
}
