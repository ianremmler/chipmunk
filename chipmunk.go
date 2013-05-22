/*
Package chipmunk is an interface to Chipmunk Physics library.

Usage example: https://github.com/ftrvxmtrx/gochipmunk/blob/master/chipmunk-demo/main.go
*/
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

/*
#cgo LDFLAGS: -lchipmunk -lm

#include <chipmunk/chipmunk.h>

#if CP_USE_DOUBLES != 1
#error CP_USE_DOUBLES != 1 while Go bindings only has double precision support
#endif
*/
import "C"

import (
	"fmt"
	"unsafe"
)

////////////////////////////////////////////////////////////////////////////////

// Timestamp is type used for various timestamps.
type Timestamp uint

// CollisionType is a type used for Space.CollisionType.
type CollisionType C.cpCollisionType

// Group is a type used for Shape.Group.
type Group C.cpGroup

// Layers is a type used for Shape.Layers.
type Layers uint

////////////////////////////////////////////////////////////////////////////////

const (
	// NoGroup is a value for Shape.Group signifying that a shape is in no group.
	NoGroup = Group(0)
	// AllLayers is a value for Shape.Layers signifying that a shape is in every layer.
	AllLayers = Layers(0)
)

// Chipmunk version.
const (
	VersionMajor   = int(C.CP_VERSION_MAJOR)
	VersionMinor   = int(C.CP_VERSION_MINOR)
	VersionRelease = int(C.CP_VERSION_RELEASE)
)

////////////////////////////////////////////////////////////////////////////////

// AreaForCircle returns the area of a hollow circle.
func AreaForCircle(r1, r2 float64) float64 {
	return float64(C.cpAreaForCircle(C.cpFloat(r1), C.cpFloat(r2)))
}

// AreaForPoly returns the signed area of a polygon.
// A Clockwise winding gives positive area. This is probably backwards from what
// you expect, but matches Chipmunk's the winding for poly shapes.
func AreaForPoly(verts []Vect) float64 {
	v := (*C.cpVect)(unsafe.Pointer(&verts[0]))
	return float64(C.cpAreaForPoly(C.int(len(verts)), v))
}

// AreaForSegment returns the area of a fattened (capsule shaped) line segment.
func AreaForSegment(a, b Vect, r float64) float64 {
	return float64(C.cpAreaForSegment(a.c(), b.c(), C.cpFloat(r)))
}

// CentroidForPoly returns the natural centroid of a polygon.
func CentroidForPoly(verts []Vect) Vect {
	v := (*C.cpVect)(unsafe.Pointer(&verts[0]))
	return cpVect(C.cpCentroidForPoly(C.int(len(verts)), v))
}

// ConvexHull calculates the convex hull of a given set of points.
// Returns the points and index of the first vertex in the hull came from the input.
// Tolerance is the allowed amount to shrink the hull when simplifying it.
// A tolerance of 0.0 creates an exact hull.
func ConvexHull(verts []Vect, tolerance float64) ([]Vect, int) {
	result := make([]Vect, len(verts))
	var first C.int
	v := (*C.cpVect)(unsafe.Pointer(&verts[0]))
	r := (*C.cpVect)(unsafe.Pointer(&result[0]))
	num := int(C.cpConvexHull(C.int(len(verts)), v, r, &first, C.cpFloat(tolerance)))
	final := make([]Vect, num, num)
	copy(final, result[:num])
	return final, int(first)
}

// MomentForBox returns the moment of inertia for a solid box.
func MomentForBox(m, width, height float64) float64 {
	return float64(C.cpMomentForBox(C.cpFloat(m), C.cpFloat(width), C.cpFloat(height)))
}

// MomentForBox2 returns the moment of inertia for a solid box.
func MomentForBox2(m float64, box BB) float64 {
	return float64(C.cpMomentForBox2(C.cpFloat(m), box.c()))
}

// MomentForCircle returns the moment of inertia for a circle.
func MomentForCircle(m, r1, r2 float64, offset Vect) float64 {
	return float64(C.cpMomentForCircle(C.cpFloat(m), C.cpFloat(r1), C.cpFloat(r2), offset.c()))
}

// MomentForPoly returns the moment of inertia for a solid polygon shape assuming
// it's center of gravity is at it's centroid. The offset is added to each vertex.
func MomentForPoly(m float64, verts []Vect, offset Vect) float64 {
	v := (*C.cpVect)(unsafe.Pointer(&verts[0]))
	return float64(C.cpMomentForPoly(C.cpFloat(m), C.int(len(verts)), v, offset.c()))
}

// MomentForSegment returns the moment of inertia for a line segment.
func MomentForSegment(m float64, a, b Vect) float64 {
	return float64(C.cpMomentForSegment(C.cpFloat(m), a.c(), b.c()))
}

// RecenterPoly centers the polygon on the origin (subtracts the centroid
// of the polygon from each vertex).
func RecenterPoly(verts []Vect) {
	v := (*C.cpVect)(unsafe.Pointer(&verts[0]))
	C.cpRecenterPoly(C.int(len(verts)), v)
}

// String converts Group to a human-readable string.
func (g Group) String() string {
	if g == NoGroup {
		return "NoGroup"
	}

	return fmt.Sprintf("(Group){%d}", uint(g))
}

// String converts Layers to a human-readable string.
func (l Layers) String() string {
	if l == AllLayers {
		return "AllLayers"
	}

	return fmt.Sprintf("(Layers){0x%x}", uint(l))
}

// Version returns Chipmunk version string.
func Version() string {
	return C.GoString(C.cpVersionString)
}

// VertsEqual returns a boolean reporting whether a == b.
func VertsEqual(a, b []Vect) bool {
	if len(a) != len(b) {
		return false
	}

	for i, c := range a {
		if c != b[i] {
			return false
		}
	}

	return true
}

// boolToC converts bool to C.cpBool.
func boolToC(b bool) C.cpBool {
	v := 0

	if b {
		v = 1
	}

	return C.cpBool(v)
}

// c converts CollisionType to c.cpCollisionType.
func (c CollisionType) c() C.cpCollisionType {
	return C.cpCollisionType(c)
}

// c converts Group to c.cpGroup.
func (g Group) c() C.cpGroup {
	return C.cpGroup(g)
}

// c converts Layers to c.cpLayers.
func (l Layers) c() C.cpLayers {
	return C.cpLayers(l)
}

// cpBool converts C.cpBool to bool.
func cpBool(b C.cpBool) bool {
	return int(b) != 0
}

// cpData converts C.cpDataPointer to interface.
func cpData(p C.cpDataPointer) interface{} {
	data := *(*interface{})(p)
	return data
}

// dataToC converts interface to C.cpDataPointer.
func dataToC(data interface{}) C.cpDataPointer {
	return C.cpDataPointer(unsafe.Pointer(&data))
}
