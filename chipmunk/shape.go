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

/*
#include <chipmunk.h>

static cpShapeType shape_type(cpShape *s) {
  return s->klass_private->type;
}
*/
import "C"

////////////////////////////////////////////////////////////////////////////////

// NearestPointQueryInfo is a nearest point query struct.
type NearestPointQueryInfo struct {
  // Shape is the nearest shape. nil if no shape was within range.
  Shape
  // P is the closest point on the shape's surface (world space coordinates).
  P Vect
  // D is the distance to the point. Negative if the point is inside the shape.
  D float64
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
  Free()
  String() string
  c() *C.cpShape
  ContainedInSpace(Space) bool
  NearestPointQuery(Vect) (float64, NearestPointQueryInfo)
  SegmentQuery(Vect, Vect) (bool, SegmentQueryInfo)
  PointQuery(Vect) bool
  CacheBB() BB
  Update(Vect, Vect) BB
  BB() BB
  Sensor() bool
  Elasticity() float64
  Friction() float64
  SurfaceVelocity() Vect
  CollisionType() CollisionType
  Group() Group
  Layers() Layers
  Body() Body
  SetSensor(bool)
  SetElasticity(float64)
  SetFriction(float64)
  SetSurfaceVelocity(Vect)
  SetCollisionType(CollisionType)
  SetGroup(Group)
  SetLayers(Layers)
  SetBody(Body)
}

// shapeBase is a base for every shape.
type shapeBase struct {
  s *C.cpShape
}
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
  return cpBB(C.cpShapeGetBB(s.s))
}

// Body returns the rigid body this collision shape is attached to.
func (s shapeBase) Body() Body {
  return Body{b: C.cpShapeGetBody(s.s)}
}

// CacheBB updates, caches and returns the bounding box of a shape based on the body it's attached to.
func (s shapeBase) CacheBB() BB {
  return cpBB(C.cpShapeCacheBB(s.s))
}

// CollisionType returns collision type of the shape used when picking collision handlers.
func (s shapeBase) CollisionType() CollisionType {
  return CollisionType(C.cpShapeGetCollisionType(s.s))
}

// ContainedInSpace returns true if the space contains the shape.
func (s shapeBase) ContainedInSpace(space Space) bool {
  return cpBool(C.cpSpaceContainsShape(space.c(), s.s))
}

// Elasticity returns shape's coefficient of restitution.
func (s shapeBase) Elasticity() float64 {
  return float64(C.cpShapeGetElasticity(s.s))
}

// Free removes a shape.
func (s shapeBase) Free() {
  C.cpShapeFree(s.s)
}

// Friction returns shape's coefficient of friction.
func (s shapeBase) Friction() float64 {
  return float64(C.cpShapeGetFriction(s.s))
}

// Group returns a group of the shape. Shapes in the same group don't collide.
func (s shapeBase) Group() Group {
  return Group(C.cpShapeGetGroup(s.s))
}

// Layers returns layers bitmask of the shape. Shapes collide only if bitwise
// of their layers is non-zero.
func (s shapeBase) Layers() Layers {
  return Layers(C.cpShapeGetLayers(s.s))
}

// NearestPointQuery finds the closest point on the surface of shape to a specific point.
// The first returned value is the distance between the points.
// A negative distance means the point is inside the shape.
func (s shapeBase) NearestPointQuery(p Vect) (float64, NearestPointQueryInfo) {
  var out C.cpNearestPointQueryInfo
  d := float64(C.cpShapeNearestPointQuery(s.s, p.c(), &out))
  return d, NearestPointQueryInfo{Shape: cpShape(out.shape), P: cpVect(out.p), D: float64(out.d)}
}

// PointQuery returns true if a point lies within a shape.
func (s shapeBase) PointQuery(p Vect) bool {
  return cpBool(C.cpShapePointQuery(s.s, p.c()))
}

// ResetShapeIdCounter is used to reset the shape ID counter when recreating a space.
// When initializing a shape, it's hash value comes from a counter.
// Because the hash value may affect iteration order, you can reset the shape ID counter
// when recreating a space. This will make the simulation be deterministic.
func ResetShapeIdCounter() {
  C.cpResetShapeIdCounter()
}

// SegmentQuery performs a segment query against a shape.
func (s shapeBase) SegmentQuery(a, b Vect) (bool, SegmentQueryInfo) {
  var out C.cpSegmentQueryInfo
  d := cpBool(C.cpShapeSegmentQuery(s.s, a.c(), b.c(), &out))
  return d, SegmentQueryInfo{Shape: cpShape(out.shape), T: float64(out.t), N: cpVect(out.n)}
}

// Sensor returns true when shape is "sensor" one, i.e. does not produce collisions.
func (s shapeBase) Sensor() bool {
  return cpBool(C.cpShapeGetSensor(s.s))
}

// SetBody sets the rigid body this collision shape is attached to.
func (s shapeBase) SetBody(b Body) {
  C.cpShapeSetBody(s.s, b.b)
}

// SetCollisionType sets collision type of the shape used when picking collision handlers.
func (s shapeBase) SetCollisionType(t CollisionType) {
  C.cpShapeSetCollisionType(s.s, t.c())
}

// SetElasticity sets coefficient of restitution.
func (s shapeBase) SetElasticity(e float64) {
  C.cpShapeSetElasticity(s.s, C.cpFloat(e))
}

// SetFriction sets coefficient of friction.
func (s shapeBase) SetFriction(f float64) {
  C.cpShapeSetFriction(s.s, C.cpFloat(f))
}

// SetGroup sets a group of the shape. Shapes in the same group don't collide.
func (s shapeBase) SetGroup(g Group) {
  C.cpShapeSetGroup(s.s, g.c())
}

// SetLayers sets layers bitmask of the shape. Shapes collide only if bitwise
// of their layers is non-zero.
func (s shapeBase) SetLayers(l Layers) {
  C.cpShapeSetLayers(s.s, l.c())
}

// SetSensor sets if the shape is "sensor" one, i.e. does not produce collisions,
// but still calls collision callbacks.
func (s shapeBase) SetSensor(b bool) {
  C.cpShapeSetSensor(s.s, boolToC(b))
}

// SetSurfaceVelocity sets surface velocity used when solving friction.
func (s shapeBase) SetSurfaceVelocity(v Vect) {
  C.cpShapeSetSurfaceVelocity(s.s, v.c())
}

// SetUserData sets user definable data pointer.
// Generally this points to your the game object so you can access it
// when given a Shape reference in a callback.
func (s shapeBase) SetUserData(data interface{}) {
  C.cpShapeSetUserData(s.s, dataToC(data))
}

// Space returns space the body was added to or nil if the body doesn't belong to any space.
func (s shapeBase) Space() Space {
  return cpSpace(C.cpShapeGetSpace(s.s))
}

// SurfaceVelocity returns surface velocity used when solving friction.
func (s shapeBase) SurfaceVelocity() Vect {
  return cpVect(C.cpShapeGetSurfaceVelocity(s.s))
}

// Update updates, caches and returns the bounding box of a shape with an explicit transformation.
func (s shapeBase) Update(pos, rot Vect) BB {
  return cpBB(C.cpShapeUpdate(s.s, pos.c(), rot.c()))
}

// UserData returns user defined data.
func (s shapeBase) UserData() interface{} {
  return cpData(C.cpShapeGetUserData(s.s))
}

// c converts Shape to C.cpShape pointer.
func (s shapeBase) c() *C.cpShape {
  return s.s
}

// cpShape converts C.cpShape pointer to Shape.
func cpShape(s *C.cpShape) Shape {
  if nil == s {
    return nil
  }

  switch shapeType(C.shape_type(s)) {
  case circleShapeType:
    return CircleShape{shapeBase{s}}

  case segmentShapeType:
    return SegmentShape{shapeBase{s}}

  case polyShapeType:
    return PolyShape{shapeBase{s}}
  }

  panic("unknown type of shape in cpShape")
}

// Local Variables:
// indent-tabs-mode: nil
// tab-width: 2
// End:
// ex: set tabstop=2 shiftwidth=2 expandtab:
