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

#include <chipmunk.h>

extern void eachArbiter_body(cpBody *b, cpArbiter *a, void *p);
extern void eachConstraint_body(cpBody *b, cpConstraint *c, void *p);
extern void eachShape_body(cpBody *b, cpShape *s, void *p);
extern void updatePosition(cpBody *b, cpFloat dt);
extern void updateVelocity(cpBody *b, cpVect gravity, cpFloat damping, cpFloat dt);

inline void body_each_arbiter(cpBody *body, void *f) {
  cpBodyEachArbiter(body, eachArbiter_body, f);
}

inline void body_each_constraint(cpBody *body, void *f) {
  cpBodyEachConstraint(body, eachConstraint_body, f);
}

inline void body_each_shape(cpBody *body, void *f) {
  cpBodyEachShape(body, eachShape_body, f);
}

inline void body_set_position_func(cpBody *body, cpBool set) {
  body->position_func = set ? updatePosition : cpBodyUpdatePosition;
}

inline void body_set_velocity_func(cpBody *body, cpBool set) {
  body->velocity_func = set ? updateVelocity : cpBodyUpdateVelocity;
}

// Local Variables:
// indent-tabs-mode: nil
// c-basic-offset: 2
// End:
// ex: set tabstop=2 shiftwidth=2 expandtab:
