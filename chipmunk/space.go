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

extern void pointQuery(cpShape *s, void *p);

static void space_point_query(cpSpace *s, cpVect point, cpLayers layers, cpGroup group, void *p) {
  cpSpacePointQuery(s, point, layers, group, pointQuery, p);
}

extern void nearestPointQuery(cpShape *s, cpFloat distance, cpVect point, void *p);

static void space_nearest_point_query(cpSpace *space, cpVect point, cpFloat maxDistance,
                                      cpLayers layers, cpGroup group, void *f) {
  cpSpaceNearestPointQuery(space, point, maxDistance, layers, group, nearestPointQuery, f);
}

extern void segmentQuery(cpShape *s, cpFloat t, cpVect n, void *p);

static void space_segment_query(cpSpace *space, cpVect start, cpVect end, cpLayers layers, cpGroup group, void *f) {
  cpSpaceSegmentQuery(space, start, end, layers, group, segmentQuery, f);
}

extern void bbQuery(cpShape *s, void *p);

static void space_bb_query(cpSpace *space, cpBB bb, cpLayers layers, cpGroup group, void *f) {
  cpSpaceBBQuery(space, bb, layers, group, bbQuery, f);
}
*/
import "C"

import (
  "unsafe"
)

// Space is a basic unit of simulation in Chipmunk.
type Space struct {
  s *C.cpSpace
}

type ContainedInSpace interface {
  ContainedInSpace(s Space) bool
}

// NewSpace creates a new space.
func NewSpace() Space {
  return Space{s: C.cpSpaceNew()}
}

// Destroy removes a space.
func (s Space) Destroy() {
  C.cpSpaceDestroy(s.s)
  s.s = nil
}

func cpSpace(s *C.cpSpace) Space {
  return Space{s}
}

/////////////////////////////////////////////////////////////////////////////

// Iterations returns the number of iterations to use in the impulse solver (to solve contacts).
func (s Space) Iterations() int {
  return int(C.cpSpaceGetIterations(s.s))
}

// Gravity returns current gravity used when integrating velocity for rigid bodies.
func (s Space) Gravity() Vect {
  return cpVect(C.cpSpaceGetGravity(s.s))
}

// Damping returns the damping rate expressed as the fraction of velocity bodies retain each second.
func (s Space) Damping() float64 {
  return float64(C.cpSpaceGetDamping(s.s))
}

// IdleSpeedThreshold returns speed threshold for a body to be considered idle.
func (s Space) IdleSpeedThreshold() float64 {
  return float64(C.cpSpaceGetIdleSpeedThreshold(s.s))
}

// SleepTimeThreshold returns the time a groups of bodies must remain idle in order to "fall asleep".
func (s Space) SleepTimeThreshold() float64 {
  return float64(C.cpSpaceGetSleepTimeThreshold(s.s))
}

// CollisionSlop returns amount of encouraged penetration between colliding shapes.
func (s Space) CollisionSlop() float64 {
  return float64(C.cpSpaceGetCollisionSlop(s.s))
}

// CollisionBias returns the speed of how fast overlapping shapes are pushed apart.
func (s Space) CollisionBias() float64 {
  return float64(C.cpSpaceGetCollisionBias(s.s))
}

// CollisionPersistence returns the number of frames that contact information should persist.
func (s Space) CollisionPersistence() uint {
  return uint(C.cpSpaceGetCollisionPersistence(s.s))
}

// EnableContactGraph returns true if rebuild of the contact graph during each step is enabled.
func (s Space) EnableContactGraph() bool {
  return 0 != int(C.cpSpaceGetEnableContactGraph(s.s))
}

// UserData returns user defined data.
func (s Space) UserData() interface{} {
  return cpData(C.cpSpaceGetUserData(s.s))
}

/////////////////////////////////////////////////////////////////////////////

// SetIterations sets the number of iterations to use in the impulse solver to solve contacts.
func (s Space) SetIterations(i int) {
  C.cpSpaceSetIterations(s.s, C.int(i))
}

// SetGravity sets the gravity to pass to rigid bodies when integrating velocity.
func (s Space) SetGravity(g Vect) {
  C.cpSpaceSetGravity(s.s, g.c())
}

// SetDamping sets the damping rate expressed as the fraction of velocity bodies retain each second.
// A value of 0.9 would mean that each body's velocity will drop 10% per second.
// The default value is 1.0, meaning no damping is applied.
// Note this damping value is different than those of DampedSpring and DampedRotarySpring.
func (s Space) SetDamping(d float64) {
  C.cpSpaceSetDamping(s.s, C.cpFloat(d))
}

// SetIdleSpeedThreshold sets the speed threshold for a body to be considered idle.
// The default value of 0.0 means to let the space guess a good threshold based on gravity.
func (s Space) SetIdleSpeedThreshold(t float64) {
  C.cpSpaceSetIdleSpeedThreshold(s.s, C.cpFloat(t))
}

// SetSleepTimeThreshold sets the time a group of bodies must remain idle in order to fall asleep.
// Enabling sleeping also implicitly enables the the contact graph.
// The default value of math.Inf(1) disables the sleeping algorithm.
func (s Space) SetSleepTimeThreshold(t float64) {
  C.cpSpaceSetSleepTimeThreshold(s.s, C.cpFloat(t))
}

// SetCollisionSlop sets amount of encouraged penetration between colliding shapes.
// Used to reduce oscillating contacts and keep the collision cache warm.
// Defaults to 0.1. If you have poor simulation quality,
// increase this number as much as possible without allowing visible amounts of overlap.
func (s Space) SetCollisionSlop(sl float64) {
  C.cpSpaceSetCollisionSlop(s.s, C.cpFloat(sl))
}

// SetCollisionBias sets the speed of how fast overlapping shapes are pushed apart.
// Expressed as a fraction of the error remaining after each second.
// Defaults to pow(1.0 - 0.1, 60.0) meaning that Chipmunk fixes 10% of overlap each frame at 60Hz.
func (s Space) SetCollisionBias(b float64) {
  C.cpSpaceSetCollisionBias(s.s, C.cpFloat(b))
}

// SetCollisionPersistence sets the number of frames that contact information should persist.
// Defaults to 3. There is probably never a reason to change this value.
func (s Space) SetCollisionPersistence(p uint) {
  C.cpSpaceSetCollisionPersistence(s.s, C.cpTimestamp(p))
}

// SetEnableContactGraph enables a rebuild of the contact graph during each step.
// Must be enabled to use the EachArbiter() method of Body.
// Disabled by default for a small performance boost. Enabled implicitly when the sleeping feature is enabled.
func (s Space) SetEnableContactGraph(cg bool) {
  C.cpSpaceSetEnableContactGraph(s.s, boolToC(cg))
}

// SetUserData sets user definable data pointer.
// Generally this points to your game's controller or game state
// so you can access it when given a Space reference in a callback.
func (s Space) SetUserData(data interface{}) {
  C.cpSpaceSetUserData(s.s, dataToC(data))
}

/////////////////////////////////////////////////////////////////////////////

// IsLocked returns true if objects cannot be added/removed inside a callback.
func (s Space) IsLocked() bool {
  return cpBool(C.cpSpaceIsLocked(s.s))
}

// RemoveCollisionHandler unsets a collision handler.
func (s Space) RemoveCollisionHandler(a CollisionType, b CollisionType) {
  C.cpSpaceRemoveCollisionHandler(s.s, C.cpCollisionType(a), C.cpCollisionType(b))
}

// Step makes the space step forward in time by dt seconds.
func (s Space) Step(dt float64) {
  C.cpSpaceStep(s.s, C.cpFloat(dt))
}

// UseSpatialHash switches the space to use a spatial has as it's spatial index.
func (s Space) UseSpatialHash(dim float64, count int) {
  C.cpSpaceUseSpatialHash(s.s, C.cpFloat(dim), C.int(count))
}

// ReindexStatic updates the collision detection info for the static shape in the space.
func (s Space) ReindexStatic() {
  C.cpSpaceReindexStatic(s.s)
}

// ReindexShape updates the collision detection data for a specific shape in the space.
func (s Space) ReindexShape(sh Shape) {
  C.cpSpaceReindexShape(s.s, sh.c())
}

// ReindexShapesForBody updates the collision detection data for all shapes attached to a body.
func (s Space) ReindexShapesForBody(b Body) {
  C.cpSpaceReindexShapesForBody(s.s, b.c())
}

/////////////////////////////////////////////////////////////////////////////

// AddShape adds a collision shape to the simulation.
// If the shape is attached to a static body, it will be added as a static shape.
func (s Space) AddShape(sh Shape) Shape {
  return cpShape(C.cpSpaceAddShape(s.s, sh.c()))
}

// AddStaticShape explicity adds a shape as a static shape to the simulation.
func (s Space) AddStaticShape(sh Shape) Shape {
  return cpShape(C.cpSpaceAddStaticShape(s.s, sh.c()))
}

// AddBody adds a rigid body to the simulation.
func (s Space) AddBody(b Body) Body {
  return Body{b: C.cpSpaceAddBody(s.s, b.b)}
}

// AddConstraint adds a constraint to the simulation.
func (s Space) AddConstraint(c Constraint) Constraint {
  return cpConstraint(C.cpSpaceAddConstraint(s.s, c.c()))
}

// RemoveShape removes a collision shape from the simulation.
func (s Space) RemoveShape(sh Shape) {
  C.cpSpaceRemoveShape(s.s, sh.c())
}

// RemoveStaticShape removes a collision shape added using AddStaticShape() from the simulation.
func (s Space) RemoveStaticShape(sh Shape) {
  C.cpSpaceRemoveStaticShape(s.s, sh.c())
}

// RemoveBody removes a rigid body from the simulation.
func (s Space) RemoveBody(b Body) {
  C.cpSpaceRemoveBody(s.s, b.b)
}

// RemoveConstraint removes a constraint from the simulation.
func (s Space) RemoveConstraint(c Constraint) {
  C.cpSpaceRemoveConstraint(s.s, c.c())
}

/////////////////////////////////////////////////////////////////////////////

// Contains tests if a collision shape, rigid body or a constraint has been added to the space.
func (s Space) Contains(o ContainedInSpace) bool {
  return o.ContainedInSpace(s)
}

// PointQuery is a callback function type for PointQuery function.
type PointQuery func(s Shape)

//export pointQuery
func pointQuery(s *C.cpShape, p unsafe.Pointer) {
  f := *(*PointQuery)(p)
  f(cpShape(s))
}

// PointQuery queries the space at a point and calls a callback function for each shape found.
func (s Space) PointQuery(point Vect, layers Layers, group Group, f PointQuery) {
  C.space_point_query(s.s, point.c(), layers.c(), group.c(), unsafe.Pointer(&f))
}

// PointQueryFirst queries the space at a point and returns
// the first shape found. Returns nil if no shapes were found.
func (s Space) PointQueryFirst(point Vect, layers Layers, group Group) Shape {
  return cpShape(C.cpSpacePointQueryFirst(s.s, point.c(), layers.c(), group.c()))
}

// NearestPointQuery is a callback function type for NearestPointQuery function.
type NearestPointQuery func(s Shape, distance float64, point Vect)

//export nearestPointQuery
func nearestPointQuery(s *C.cpShape, distance C.cpFloat, point C.cpVect, p unsafe.Pointer) {
  f := *(*NearestPointQuery)(p)
  f(cpShape(s), float64(distance), cpVect(point))
}

// NearestPointQuery queries the space at a point and calls a callback function for each shape found.
func (s Space) NearestPointQuery(point Vect, maxDistance float64, layers Layers, group Group, f NearestPointQuery) {
  C.space_nearest_point_query(s.s, point.c(), C.cpFloat(maxDistance), layers.c(), group.c(), unsafe.Pointer(&f))
}

// SegmentQuery is a query callback function type.
type SegmentQuery func(s Shape, t float64, n Vect)

//export segmentQuery
func segmentQuery(s *C.cpShape, t C.cpFloat, n C.cpVect, p unsafe.Pointer) {
  f := *(*SegmentQuery)(p)
  f(cpShape(s), float64(t), cpVect(n))
}

// SegmentQuery performs a directed line segment query (like a raycast)
// against the space calling a callback function for each shape intersected.
func (s Space) SegmentQuery(start, end Vect, layers Layers, group Group, f SegmentQuery) {
  C.space_segment_query(s.s, start.c(), end.c(), layers.c(), group.c(), unsafe.Pointer(&f))
}

// BBQuery is a rectangle query callback function type.
type BBQuery func(s Shape)

//export bbQuery
func bbQuery(s *C.cpShape, p unsafe.Pointer) {
  f := *(*BBQuery)(p)
  f(cpShape(s))
}

// BBQuery performs a fast rectangle query on the space calling a callback function for each shape found.
// Only the shape's bounding boxes are checked for overlap, not their full shape.
func (s Space) BBQuery(bb BB, layers Layers, group Group, f BBQuery) {
  C.space_bb_query(s.s, bb.c(), layers.c(), group.c(), unsafe.Pointer(&f))
}
