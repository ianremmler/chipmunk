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
// #include "shape.h"
import "C"

import (
	"unsafe"
)

////////////////////////////////////////////////////////////////////////////////

// NearestPointQueryInfo is a nearest point query struct.
type NearestPointQueryInfo struct {
	// Shape is the nearest shape. nil if no shape was within range.
	Shape
	// P is the closest point on the shape's surface (world space coordinates).
	P Vect
	// D is the distance to the point. Negative if the point is inside the shape.
	D float64
	// G is the gradient of the signed distance field for the shape.
	G Vect
}

// SegmentQueryInfo holds the result of SegmentQuery.
type SegmentQueryInfo struct {
	// Shape is the one that was hit. nil if no collision occured.
	Shape
	// Normalized distance along the query segment in the range [0;1].
	T float64
	// Normal of the surface hit.
	N Vect
}

// Shape is an opaque collision shape struct.
type Shape interface {
	BB() BB
	Body() Body
	CacheBB() BB
	CollisionType() CollisionType
	Elasticity() float64
	Free()
	Friction() float64
	Group() Group
	Layers() Layers
	NearestPointQuery(Vect) (float64, NearestPointQueryInfo)
	PointQuery(Vect) bool
	SegmentQuery(Vect, Vect) (bool, SegmentQueryInfo)
	Sensor() bool
	SetBody(Body)
	SetCollisionType(CollisionType)
	SetElasticity(float64)
	SetFriction(float64)
	SetGroup(Group)
	SetLayers(Layers)
	SetSensor(bool)
	SetSurfaceVelocity(Vect)
	String() string
	SurfaceVelocity() Vect
	Update(Vect, Vect) BB
	c() *C.cpShape
}

// shapeBase is a base for every shape.
type shapeBase uintptr

type shapeType int

////////////////////////////////////////////////////////////////////////////////

const (
	circleShapeType  = shapeType(C.CP_CIRCLE_SHAPE)
	segmentShapeType = shapeType(C.CP_SEGMENT_SHAPE)
	polyShapeType    = shapeType(C.CP_POLY_SHAPE)
)

////////////////////////////////////////////////////////////////////////////////

// BB returns current bounding box of the shape.
func (s shapeBase) BB() BB {
	return cpBB(C.cpShapeGetBB(s.c()))
}

// Body returns the rigid body this collision shape is attached to.
func (s shapeBase) Body() Body {
	return cpBody(C.cpShapeGetBody(s.c()))
}

// CacheBB updates, caches and returns the bounding box of a shape based on the body it's attached to.
func (s shapeBase) CacheBB() BB {
	return cpBB(C.cpShapeCacheBB(s.c()))
}

// CollisionType returns collision type of the shape used when picking collision handlers.
func (s shapeBase) CollisionType() CollisionType {
	return CollisionType(C.cpShapeGetCollisionType(s.c()))
}

// Elasticity returns shape's coefficient of restitution.
func (s shapeBase) Elasticity() float64 {
	return float64(C.cpShapeGetElasticity(s.c()))
}

// Free removes a shape.
func (s shapeBase) Free() {
	C.cpShapeFree(s.c())
}

// Friction returns shape's coefficient of friction.
func (s shapeBase) Friction() float64 {
	return float64(C.cpShapeGetFriction(s.c()))
}

// Group returns a group of the shape. Shapes in the same group don't collide.
func (s shapeBase) Group() Group {
	return Group(C.cpShapeGetGroup(s.c()))
}

// HitDistance returns the hit distance for a segment query.
func (s SegmentQueryInfo) HitDistance(start, end Vect) float64 {
	return start.Dist(end) * s.T
}

// HitPoint returns the hit point for a segment query.
func (s SegmentQueryInfo) HitPoint(start, end Vect) Vect {
	return start.Lerp(end, s.T)
}

// Layers returns layers bitmask of the shape. Shapes collide only if bitwise
// of their layers is non-zero.
func (s shapeBase) Layers() Layers {
	return Layers(C.cpShapeGetLayers(s.c()))
}

// NearestPointQuery finds the closest point on the surface of shape to a specific point.
// The first returned value is the distance between the points.
// A negative distance means the point is inside the shape.
func (s shapeBase) NearestPointQuery(p Vect) (float64, NearestPointQueryInfo) {
	var out C.cpNearestPointQueryInfo
	d := float64(C.cpShapeNearestPointQuery(s.c(), p.c(), &out))
	return d, NearestPointQueryInfo{Shape: cpShape(out.shape), P: cpVect(out.p), D: float64(out.d), G: cpVect(out.g)}
}

// PointQuery returns true if a point lies within a shape.
func (s shapeBase) PointQuery(p Vect) bool {
	return cpBool(C.cpShapePointQuery(s.c(), p.c()))
}

// ResetShapeIDCounter is used to reset the shape ID counter when recreating a space.
// When initializing a shape, it's hash value comes from a counter.
// Because the hash value may affect iteration order, you can reset the shape ID counter
// when recreating a space. This will make the simulation be deterministic.
func ResetShapeIDCounter() {
	C.cpResetShapeIdCounter()
}

// SegmentQuery performs a segment query against a shape.
func (s shapeBase) SegmentQuery(a, b Vect) (bool, SegmentQueryInfo) {
	var out C.cpSegmentQueryInfo
	d := cpBool(C.cpShapeSegmentQuery(s.c(), a.c(), b.c(), &out))
	return d, SegmentQueryInfo{Shape: cpShape(out.shape), T: float64(out.t), N: cpVect(out.n)}
}

// Sensor returns true when shape is "sensor" one, i.e. does not produce collisions.
func (s shapeBase) Sensor() bool {
	return cpBool(C.cpShapeGetSensor(s.c()))
}

// SetBody sets the rigid body this collision shape is attached to.
func (s shapeBase) SetBody(b Body) {
	C.cpShapeSetBody(s.c(), b.c())
}

// SetCollisionType sets collision type of the shape used when picking collision handlers.
func (s shapeBase) SetCollisionType(t CollisionType) {
	C.cpShapeSetCollisionType(s.c(), t.c())
}

// SetElasticity sets coefficient of restitution.
func (s shapeBase) SetElasticity(e float64) {
	C.cpShapeSetElasticity(s.c(), C.cpFloat(e))
}

// SetFriction sets coefficient of friction.
func (s shapeBase) SetFriction(f float64) {
	C.cpShapeSetFriction(s.c(), C.cpFloat(f))
}

// SetGroup sets a group of the shape. Shapes in the same group don't collide.
func (s shapeBase) SetGroup(g Group) {
	C.cpShapeSetGroup(s.c(), g.c())
}

// SetLayers sets layers bitmask of the shape. Shapes collide only if bitwise
// of their layers is non-zero.
func (s shapeBase) SetLayers(l Layers) {
	C.cpShapeSetLayers(s.c(), l.c())
}

// SetSensor sets if the shape is "sensor" one, i.e. does not produce collisions,
// but still calls collision callbacks.
func (s shapeBase) SetSensor(b bool) {
	C.cpShapeSetSensor(s.c(), boolToC(b))
}

// SetSurfaceVelocity sets surface velocity used when solving friction.
func (s shapeBase) SetSurfaceVelocity(v Vect) {
	C.cpShapeSetSurfaceVelocity(s.c(), v.c())
}

// SetUserData sets user definable data pointer.
// Generally this points to your the game object so you can access it
// when given a Shape reference in a callback.
func (s shapeBase) SetUserData(data interface{}) {
	C.cpShapeSetUserData(s.c(), dataToC(data))
}

// Space returns space the body was added to or nil if the body doesn't belong to any space.
func (s shapeBase) Space() Space {
	return cpSpace(C.cpShapeGetSpace(s.c()))
}

// SurfaceVelocity returns surface velocity used when solving friction.
func (s shapeBase) SurfaceVelocity() Vect {
	return cpVect(C.cpShapeGetSurfaceVelocity(s.c()))
}

// Update updates, caches and returns the bounding box of a shape with an explicit transformation.
func (s shapeBase) Update(pos, rot Vect) BB {
	return cpBB(C.cpShapeUpdate(s.c(), pos.c(), rot.c()))
}

// UserData returns user defined data.
func (s shapeBase) UserData() interface{} {
	return cpData(C.cpShapeGetUserData(s.c()))
}

// addToSpace adds a shape to space.
func (s shapeBase) addToSpace(space Space) {
	space.AddShape(cpShape(s.c()))
}

// c converts Shape to C.cpShape pointer.
func (s shapeBase) c() *C.cpShape {
	return (*C.cpShape)(unsafe.Pointer(s))
}

// containedInSpace returns true if the space contains the shape.
func (s shapeBase) containedInSpace(space Space) bool {
	return cpBool(C.cpSpaceContainsShape(space.c(), s.c()))
}

// cpShape converts C.cpShape pointer to Shape.
func cpShape(s *C.cpShape) Shape {
	if nil == s {
		return nil
	}

	p := cpshape(s)

	switch shapeType(C.shape_type(s)) {
	case circleShapeType:
		return CircleShape{p}

	case segmentShapeType:
		return SegmentShape{p}

	case polyShapeType:
		return PolyShape{p}
	}

	panic("unknown type of shape in cpShape")
}

// cpshape converts C.cpShape pointer to shapeBase.
func cpshape(s *C.cpShape) shapeBase {
	return shapeBase(unsafe.Pointer(s))
}

// removeFromSpace removes a shape from space.
func (s shapeBase) removeFromSpace(space Space) {
	space.RemoveShape(cpShape(s.c()))
}
