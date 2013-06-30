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

#include <chipmunk/chipmunk.h>

extern void bbQuery(cpShape *s, void *p);
extern void eachBodySpace(cpBody *b, void *p);
extern void eachConstraintSpace(cpConstraint *c, void *p);
extern void eachShapeSpace(cpShape *s, void *p);
extern void nearestPointQuery(cpShape *s, cpFloat distance, cpVect point, void *p);
extern void pointQuery(cpShape *s, void *p);
extern void postStep(cpSpace *space, cpDataPointer key, cpDataPointer data);
extern void segmentQuery(cpShape *s, cpFloat t, cpVect n, void *p);

extern cpBool begin(cpArbiter *arb, cpSpace *space, cpDataPointer data);
extern cpBool preSolve(cpArbiter *arb, cpSpace *space, cpDataPointer data);
extern void postSolve(cpArbiter *arb, cpSpace *space, cpDataPointer data);
extern void separate(cpArbiter *arb, cpSpace *space, cpDataPointer data);

extern cpBool beginDefault(cpArbiter *arb, cpSpace *space, cpDataPointer data);
extern cpBool preSolveDefault(cpArbiter *arb, cpSpace *space, cpDataPointer data);
extern void postSolveDefault(cpArbiter *arb, cpSpace *space, cpDataPointer data);
extern void separateDefault(cpArbiter *arb, cpSpace *space, cpDataPointer data);

inline cpBool space_add_poststep(cpSpace *space, cpDataPointer key, cpDataPointer data) {
	return cpSpaceAddPostStepCallback(space, (void *)postStep, key, data);
}

inline void space_add_collision_handler(cpSpace *space, cpCollisionType a, cpCollisionType b) {
	cpSpaceAddCollisionHandler(space, a, b, (void *)begin, (void *)preSolve, (void *)postSolve,
		(void *)separate, NULL);
}

inline void space_set_default_collision_handler(cpSpace *space) {
	cpSpaceSetDefaultCollisionHandler(space, (void *)beginDefault, (void *)preSolveDefault,
		(void *)postSolveDefault, (void *)separateDefault, NULL);
}

inline void space_bb_query(cpSpace *space, cpBB bb, cpLayers layers, cpGroup group, void *f) {
	cpSpaceBBQuery(space, bb, layers, group, bbQuery, f);
}

inline void space_each_body(cpSpace *space, void *f) {
	cpSpaceEachBody(space, eachBodySpace, f);
}

inline void space_each_constraint(cpSpace *space, void *f) {
	cpSpaceEachConstraint(space, eachConstraintSpace, f);
}

inline void space_each_shape(cpSpace *space, void *f) {
	cpSpaceEachShape(space, eachShapeSpace, f);
}

inline void space_nearest_point_query(cpSpace *space, cpVect point, cpFloat maxDistance,
	cpLayers layers, cpGroup group, void *f) {

	cpSpaceNearestPointQuery(space, point, maxDistance, layers, group, nearestPointQuery, f);
}

inline void space_point_query(cpSpace *s, cpVect point, cpLayers layers, cpGroup group, void *p) {
	cpSpacePointQuery(s, point, layers, group, pointQuery, p);
}

inline void space_segment_query(cpSpace *space, cpVect start, cpVect end, cpLayers layers,
	cpGroup group, void *f) {

	cpSpaceSegmentQuery(space, start, end, layers, group, segmentQuery, f);
}
