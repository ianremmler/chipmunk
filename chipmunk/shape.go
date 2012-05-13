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

// shapeBase is a base for every shape.
type shapeBase struct {
  s *C.cpShape
}

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
  Destroy()
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

// Destroy removes a shape.
func (s shapeBase) Destroy() {
  C.cpShapeDestroy(s.s)
  s.s = nil
}

// c returns C shape pointer.
func (s shapeBase) c() *C.cpShape {
  return s.s
}

/////////////////////////////////////////////////////////////////////////////

// ContainedInSpace returns true if the space contains the shape.
func (s shapeBase) ContainedInSpace(space Space) bool {
  return cpBool(C.cpSpaceContainsShape(space.s, s.s))
}

// NearestPointQuery finds the closest point on the surface of shape to a specific point.
// The first returned value is the distance between the points. A negative distance means the point is inside the shape.
func (s shapeBase) NearestPointQuery(p Vect) (float64, NearestPointQueryInfo) {
  var out C.cpNearestPointQueryInfo
  d := float64(C.cpShapeNearestPointQuery(s.s, p.c(), &out))
  return d, NearestPointQueryInfo{Shape: cpShape(out.shape), P: cpVect(out.p), D: float64(out.d)}
}

// SegmentQuery performs a segment query against a shape.
func (s shapeBase) SegmentQuery(a, b Vect) (bool, SegmentQueryInfo) {
  var out C.cpSegmentQueryInfo
  d := cpBool(C.cpShapeSegmentQuery(s.s, a.c(), b.c(), &out))
  return d, SegmentQueryInfo{Shape: cpShape(out.shape), T: float64(out.t), N: cpVect(out.n)}
}

// PointQuery returns true if a point lies within a shape.
func (s shapeBase) PointQuery(p Vect) bool {
  return cpBool(C.cpShapePointQuery(s.s, p.c()))
}

// CacheBB updates, caches and returns the bounding box of a shape based on the body it's attached to.
func (s shapeBase) CacheBB() BB {
  return cpBB(C.cpShapeCacheBB(s.s))
}

// Update updates, caches and returns the bounding box of a shape with an explicit transformation.
func (s shapeBase) Update(pos, rot Vect) BB {
  return cpBB(C.cpShapeUpdate(s.s, pos.c(), rot.c()))
}

/////////////////////////////////////////////////////////////////////////////

// BB returns current bounding box of the shape.
func (s shapeBase) BB() BB {
  return cpBB(C.cpShapeGetBB(s.s))
}

// Sensor returns true when shape is "sensor" one, i.e. does not produce collisions.
func (s shapeBase) Sensor() bool {
  return cpBool(C.cpShapeGetSensor(s.s))
}

// Elasticity returns shape's coefficient of restitution.
func (s shapeBase) Elasticity() float64 {
  return float64(C.cpShapeGetElasticity(s.s))
}

// Friction returns shape's coefficient of friction.
func (s shapeBase) Friction() float64 {
  return float64(C.cpShapeGetFriction(s.s))
}

// SurfaceVelocity returns surface velocity used when solving friction.
func (s shapeBase) SurfaceVelocity() Vect {
  return cpVect(C.cpShapeGetSurfaceVelocity(s.s))
}

// CollisionType returns collision type of the shape used when picking collision handlers.
func (s shapeBase) CollisionType() CollisionType {
  return CollisionType(C.cpShapeGetCollisionType(s.s))
}

// Group returns a group of the shape. Shapes in the same group don't collide.
func (s shapeBase) Group() Group {
  return Group(C.cpShapeGetGroup(s.s))
}

// Layers returns layers bitmask of the shape. Shapes collide only if bitwise of their layers is non-zero.
func (s shapeBase) Layers() Layers {
  return Layers(C.cpShapeGetLayers(s.s))
}

// Body returns the rigid body this collision shape is attached to.
func (s shapeBase) Body() Body {
  return Body{b: C.cpShapeGetBody(s.s)}
}

func (s shapeBase) UserData() interface{} {
  return cpData(C.cpShapeGetUserData(s.s))
}

/////////////////////////////////////////////////////////////////////////////

// SetSensor sets if the shape is "sensor" one, i.e. does not produce collisions, but still calls collision callbacks.
func (s shapeBase) SetSensor(b bool) {
  C.cpShapeSetSensor(s.s, boolToC(b))
}

// SetElasticity sets coefficient of restitution.
func (s shapeBase) SetElasticity(e float64) {
  C.cpShapeSetElasticity(s.s, C.cpFloat(e))
}

// SetFriction sets coefficient of friction.
func (s shapeBase) SetFriction(f float64) {
  C.cpShapeSetFriction(s.s, C.cpFloat(f))
}

// SetSurfaceVelocity sets surface velocity used when solving friction.
func (s shapeBase) SetSurfaceVelocity(v Vect) {
  C.cpShapeSetSurfaceVelocity(s.s, v.c())
}

// SetCollisionType sets collision type of the shape used when picking collision handlers.
func (s shapeBase) SetCollisionType(t CollisionType) {
  C.cpShapeSetCollisionType(s.s, t.c())
}

// SetGroup sets a group of the shape. Shapes in the same group don't collide.
func (s shapeBase) SetGroup(g Group) {
  C.cpShapeSetGroup(s.s, g.c())
}

// SetLayers sets layers bitmask of the shape. Shapes collide only if bitwise of their layers is non-zero.
func (s shapeBase) SetLayers(l Layers) {
  C.cpShapeSetLayers(s.s, l.c())
}

// SetBody sets the rigid body this collision shape is attached to.
func (s shapeBase) SetBody(b Body) {
  C.cpShapeSetBody(s.s, b.b)
}

func (s shapeBase) SetUserData(data interface{}) {
  C.cpShapeSetUserData(s.s, dataToC(data))
}

/////////////////////////////////////////////////////////////////////////////

// ResetShapeIdCounter is used to reset the shape ID counter when recreating a space.
// When initializing a shape, it's hash value comes from a counter.
// Because the hash value may affect iteration order, you can reset the shape ID counter
// when recreating a space. This will make the simulation be deterministic.
func ResetShapeIdCounter() {
  C.cpResetShapeIdCounter()
}

/////////////////////////////////////////////////////////////////////////////

// CircleShape is a circle shape type.
type CircleShape struct {
  shapeBase
}

// CircleShapeNew creates a new circle shape.
func CircleShapeNew(body Body, radius float64, offset Vect) Shape {
  s := C.cpCircleShapeNew(body.b, C.cpFloat(radius), offset.c())
  return CircleShape{shapeBase{s}}
}

func (s CircleShape) Offset() Vect {
  return cpVect(C.cpCircleShapeGetOffset(s.s))
}

func (s CircleShape) Radius() float64 {
  return float64(C.cpCircleShapeGetRadius(s.s))
}

func (s CircleShape) String() string {
  return "circle shape"
}

/////////////////////////////////////////////////////////////////////////////

// SegmentShape is a segment shape type.
type SegmentShape struct {
  shapeBase
}

// SegmentShapeNew creates a new segment shape
func SegmentShapeNew(body Body, a, b Vect, radius float64) Shape {
  s := C.cpSegmentShapeNew(body.b, a.c(), b.c(), C.cpFloat(radius))
  return SegmentShape{shapeBase{s}}
}

func (s SegmentShape) SetNeighbors(prev, next Vect) {
  C.cpSegmentShapeSetNeighbors(s.s, prev.c(), next.c())
}

func (s SegmentShape) A() Vect {
  return cpVect(C.cpSegmentShapeGetB(s.s))
}

func (s SegmentShape) B() Vect {
  return cpVect(C.cpSegmentShapeGetA(s.s))
}

func (s SegmentShape) Normal() Vect {
  return cpVect(C.cpSegmentShapeGetNormal(s.s))
}

func (s SegmentShape) Radius() float64 {
  return float64(C.cpSegmentShapeGetRadius(s.s))
}

func (s SegmentShape) String() string {
  return "segment shape"
}

/////////////////////////////////////////////////////////////////////////////

type shapeType int

const (
  circleShapeType  = shapeType(C.CP_CIRCLE_SHAPE)
  segmentShapeType = shapeType(C.CP_SEGMENT_SHAPE)
  polyShapeType    = shapeType(C.CP_POLY_SHAPE)
)

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
